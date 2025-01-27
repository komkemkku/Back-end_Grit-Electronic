package orders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListOrderService(ctx context.Context, req requests.OrderRequest) ([]response.OrderResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.OrderResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.payment_id", "o.shipment_id", "o.cart_id", "status", "o.created_at", "o.updated_at")

	if req.Search != "" {
		query.Where("o.status ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Execute query
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdOrderService(ctx context.Context, id int64) (*response.OrderResponses, error) {
    // ตรวจสอบว่าคำสั่งซื้อมีอยู่ในระบบหรือไม่
    ex, err := db.NewSelect().
        TableExpr("orders").
        Where("id = ?", id).
        Exists(ctx)
    if err != nil {
        return nil, err
    }
    if !ex {
        return nil, errors.New("order not found")
    }

    // สร้างตัวแปรสำหรับเก็บผลลัพธ์
    order := &response.OrderResponses{}

    // ดึงข้อมูลคำสั่งซื้อ
    err = db.NewSelect().
        TableExpr("orders AS o").
        Column("o.id", "o.user_id", "o.payment_id", "o.shipment_id", "o.cart_id", "o.status", "o.created_at", "o.updated_at").
        Where("o.id = ?", id).
        Scan(ctx, order)
    if err != nil {
        return nil, err
    }

    return order, nil
}



func CreateOrderService(ctx context.Context, req requests.OrderCreateRequest) (map[string]interface{}, error) {
	// ตรวจสอบ user_id
	user := struct {
		ID       int64  `json:"id"`
		Username string `json:"name"`
	}{}
	err := db.NewSelect().TableExpr("users").Column("id", "username").Where("id = ?", req.UserID).Scan(ctx, &user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: no user with id = %d", req.UserID)
		}
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	// ดึงข้อมูล payments ที่เกี่ยวข้องกับ user_id
	payments := []model.Payments{}
	err = db.NewSelect().
		Model(&payments).
		Where("id = ?", req.UserID). // ใช้คอลัมน์ admin_id แทน user_id
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payments: %v", err)
	}

	// ดึงข้อมูล shipments ที่เกี่ยวข้องกับ user_id
	shipments := []model.Shipments{}
	err = db.NewSelect().
		Model(&shipments).
		Where("id IN (SELECT shipment_id FROM orders WHERE user_id = ?)", req.UserID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch shipments: %v", err)
	}

	// ดึงข้อมูล carts ที่เกี่ยวข้องกับ user_id
	carts := []model.Carts{}
	err = db.NewSelect().
		Model(&carts).
		Where("user_id = ?", req.UserID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch carts: %v", err)
	}

	// สร้างคำสั่งซื้อ
	order := &model.Orders{
		UserID:     req.UserID,
		PaymentID:  req.PaymentID,
		ShipmentID: req.ShipmentID,
		CartID:     req.CartID,
		Status:     req.Status,
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err = db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %v", err)
	}

	// จัดเรียง response
	response := map[string]interface{}{
		"user":  user,
		"carts": carts,
		"order": map[string]interface{}{
			"id":          order.ID,
			"user_id":     order.UserID,
			"payment_id":  order.PaymentID,
			"shipment_id": order.ShipmentID,
			"cart_id":     order.CartID,
			"status":      order.Status,
		},
		"payments":  payments,
		"shipments": shipments,
	}

	return response, nil
}

func UpdateOrderService(ctx context.Context, id int64, req requests.OrderUpdateRequest) (*model.Orders, error) {
	// ตรวจสอบว่า order มีอยู่ในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().
		TableExpr("orders").
		Where("id = ?", id).
		Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if order exists: %v", err)
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	// ดึงข้อมูล order
	order := &model.Orders{}
	err = db.NewSelect().
		Model(order).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	// อัปเดตข้อมูล
	if req.Status != "" {
		order.Status = req.Status
	}
	if req.PaymentID != 0 {
		order.PaymentID = req.PaymentID
	}
	if req.ShipmentID != 0 {
		order.ShipmentID = req.ShipmentID
	}
	if req.CartID != 0 {
		order.CartID = req.CartID
	}
	order.SetUpdateNow() // ตั้งค่า UpdatedAt

	// บันทึกข้อมูลกลับไปยังฐานข้อมูล
	_, err = db.NewUpdate().
		Model(order).
		Column("status", "payment_id", "shipment_id", "cart_id", "updated_at").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %v", err)
	}

	return order, nil
}


func DeleteOrderService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("orders").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("order not found")
	}

	_, err = db.NewDelete().TableExpr("orders").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
