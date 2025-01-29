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

func GetByIdOrderService(ctx context.Context, userID int64) (*response.OrderResponses, error) {
	// ตรวจสอบว่าผู้ใช้งานมีอยู่ในระบบหรือไม่
	exists, err := db.NewSelect().
		TableExpr("orders").
		Where("user_id = ?", userID).
		Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if !exists {
		return nil, errors.New("user not found")
	}

	// สร้าง response object
	order := &response.OrderResponses{}

	err = db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.status", "o.created_at", "o.updated_at").
		ColumnExpr("c.total_cart_amount", "c.total_cart_price").
		ColumnExpr("p.system_bank_id", "p.price", "p.bank_name", "p.account_name", "p.account_number", "p.status AS payment_status").
		ColumnExpr("s.firstname", "s.lastname", "s.address", "s.zip_code", "s.sub_district", "s.district", "s.province", "s.status AS shipment_status").
		Join("LEFT JOIN carts AS c ON o.cart_id = c.id").
		Join("LEFT JOIN payments AS p ON o.payment_id = p.id").
		Join("LEFT JOIN shipments AS s ON o.shipment_id = s.id").
		Where("o.user_id = ?", userID).
		Scan(ctx, order)

	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	return order, nil
}

func CreateOrderService(ctx context.Context, req requests.OrderCreateRequest) (*model.Orders, error) {
	// ตรวจสอบ user_id
	ex, err := db.NewSelect().TableExpr("users").Where("id = ?", req.UserID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("user not found")
	}

	// ดึงข้อมูล payments ที่เกี่ยวข้องกับ user_id
	ex, err = db.NewSelect().TableExpr("payments").Where("id = ?", req.PaymentID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("payment not found")
	}

	// ตรวจสอบ shipments
	ex, err = db.NewSelect().TableExpr("shipments").Where("id = ?", req.ShipmentID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("shipment not found")
	}

	// ตรวจสอบ carts
	ex, err = db.NewSelect().TableExpr("carts").Where("id = ?", req.CartID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("carts not found")
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

	return order, nil
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


