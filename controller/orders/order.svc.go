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
	"github.com/uptrace/bun"
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
	exists, err := db.NewSelect().TableExpr("orders").Where("user_id = ?", userID).Exists(ctx)
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
	exists, err := db.NewSelect().TableExpr("orders").Where("user_id = ?").Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if !exists {
		return nil, errors.New("user not found")
	}

	// Fetch cart ID for the user
	var cart struct {
		ID int64 `json:"id"`
	}
	err = db.NewSelect().
		TableExpr("carts").
		Column("id").
		Where("user_id = ?", req.UserID).
		Scan(ctx, &cart)
	if err != nil {
		return nil, fmt.Errorf("failed to find cart for user_id %d: %v", req.UserID, err)
	}

	// Fetch cart items and join with product details
	var products []struct {
		ProductID   int64   `json:"product_id"`
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Quantity    int64   `json:"quantity"`
		TotalAmount float64 `json:"total_amount"`
	}
	err = db.NewSelect().
		TableExpr("cart_items AS ci").
		Join("JOIN products AS p ON p.id = ci.product_id").
		ColumnExpr("p.id AS product_id, p.name, p.price, ci.quantity, (p.price * ci.quantity) AS total_amount").
		Where("ci.cart_id = ?", cart.ID).
		Scan(ctx, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}

	// Calculate total price
	totalPrice := 0.0
	for _, p := range products {
		totalPrice += p.TotalAmount
	}

	// Create order
	order := &model.Orders{
		UserID:     req.UserID,
		ShipmentID: req.ShipmentID,
		PaymentID:  req.PaymentID,
		Total_price: totalPrice,
		Status:     "pending",
	}
	_, err = db.NewInsert().
		Model(order).
		Returning("id").
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// Create order details
	orderDetails := make([]model.OrderDetail, len(products))
	for i, p := range products {
		orderDetails[i] = model.OrderDetail{
			OrderID:            order.ID,
			ProductName:        p.Name,
			TotalProductPrice:  p.Price,
			TotalProductAmount: int(p.Quantity),
		}
	}
	_, err = db.NewInsert().
		Model(&orderDetails).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to insert order details: %v", err)
	}

	// Delete cart items and cart in a transaction
	err = db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewDelete().
			TableExpr("cart_items").
			Where("cart_id = ?", cart.ID).
			Exec(ctx); err != nil {
			return fmt.Errorf("failed to delete cart items: %v", err)
		}
		if _, err := tx.NewDelete().
			TableExpr("carts").
			Where("id = ?", cart.ID).
			Exec(ctx); err != nil {
			return fmt.Errorf("failed to delete cart: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("transaction error: %v", err)
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
