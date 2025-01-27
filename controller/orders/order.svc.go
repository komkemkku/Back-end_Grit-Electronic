package orders

import (
	"context"
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
	ex, err := db.NewSelect().TableExpr("orders").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("order not found")
	}
	order := &response.OrderResponses{}

	err = db.NewSelect().TableExpr("orders AS o").
		TableExpr("orders AS o").
		Column("o.id", "o.total_price", "o.total_amount", "o.status", "o.created_at", "o.updated_at").
		Where("o.id = ?", id).Scan(ctx, order)
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
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// ตรวจสอบ payment_id
	payments := []model.Payments{}
	err = db.NewSelect().Model(&payments).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payments: %v", err)
	}

	// ตรวจสอบ shipment_id
	shipments := []model.Shipments{}
	err = db.NewSelect().Model(&shipments).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch shipments: %v", err)
	}

	// ตรวจสอบ cart_id
	carts := []model.Carts{}
	err = db.NewSelect().Model(&carts).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch carts: %v", err)
	}

	// สร้างคำสั่งซื้อใหม่
	order := &model.Orders{
		UserID:     req.UserID,
		PaymentID:  req.PaymentID,
		ShipmentID: req.ShipmentID,
		CartID:     req.CartID,
		Status:     req.Status,
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	// บันทึกคำสั่งซื้อในฐานข้อมูล
	_, err = db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %v", err)
	}

	// เตรียมข้อมูลสำหรับการตอบกลับ
	response := map[string]interface{}{
		"user": user, // นำ "user" ไว้ด้านบนสุดของ Response
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
		"carts":     carts,
	}

	return response, nil
}

func UpdateOrderService(ctx context.Context, id int64, req requests.OrderUpdateRequest) (*model.Orders, error) {
	// ตรวจสอบว่า order มีอยู่ในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().TableExpr("orders").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	// ดึงข้อมูล order
	order := &model.Orders{}
	err = db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// อัปเดตข้อมูล
	order.Status = req.Status
	order.SetUpdateNow() // ตั้งค่า UpdatedAt

	// บันทึกข้อมูลกลับไปยังฐานข้อมูล
	_, err = db.NewUpdate().Model(order).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
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
