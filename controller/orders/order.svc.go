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
	var offset int64
	if req.Page > 0 {
		offset = (req.Page - 1) * req.Size
	}

	resp := []response.OrderResponses{}

	// สร้าง query
	query := db.NewSelect().TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.status", "o.created_at", "o.updated_at").
		ColumnExpr("SUM(ci.quantity) AS total_amount").
		ColumnExpr("SUM(p.price * ci.quantity) AS total_price").
		ColumnExpr("py.system_bank_id, py.price AS payment_price, py.bank_name, py.account_name, py.account_number, py.status AS payment_status").
		ColumnExpr("s.firstname, s.lastname, s.address, s.zip_code, s.sub_district, s.district, s.province, s.status AS shipment_status").
		Join("LEFT JOIN cart_items AS ci ON ci.order_id = o.id").
		// Join("LEFT JOIN products AS p ON p.id = ci.product_id").
		Join("LEFT JOIN payments AS py ON py.id = o.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id").
		Group("o.id, py.id, s.id")

	if req.Search != "" {
		query.Where("o.status ILIKE ?", "%"+req.Search+"%")
	}

	// นับจำนวน total records
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// ดึงข้อมูล orders พร้อม payment และ shipment
	err = query.Offset(int(offset)).Limit(int(req.Size)).Scan(ctx, &resp)
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

	err = db.NewSelect().TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.status", "o.created_at", "o.updated_at").
		ColumnExpr("SUM(ci.quantity) AS total_amount").
		ColumnExpr("SUM(p.price * ci.quantity) AS total_price").
		ColumnExpr("py.system_bank_id, py.price AS payment_price, py.bank_name, py.account_name, py.account_number, py.status AS payment_status").
		ColumnExpr("s.firstname, s.lastname, s.address, s.zip_code, s.sub_district, s.district, s.province, s.status AS shipment_status").
		Join("LEFT JOIN cart_items AS ci ON ci.order_id = o.id").
		Join("LEFT JOIN products AS p ON p.id = ci.product_id").
		Join("LEFT JOIN payments AS py ON py.id = o.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id").
		Group("o.id, py.id, s.id").
		Where("o.user_id = ?", userID).
		Scan(ctx, order)

	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	return order, nil
}
func CreateOrderService(ctx context.Context, req requests.OrderCreateRequest) (*model.Orders, error) {
	var cartID int64
	fmt.Printf("Finding cart with user_id: %d\n", req.UserID) 
	if err := db.NewSelect().Table("carts").Column("id").Where("user_id = ?", req.UserID).Scan(ctx, &cartID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cart found for user_id: %d", req.UserID)
		}
		return nil, fmt.Errorf("failed to find cart: %v", err)
	}
	fmt.Printf("Found cart ID: %d\n", cartID) 

	var cartItems []struct {
		ProductID   int64   `json:"product_id"`
		ProductName string  `json:"product_name"`
		Amount      int64   `json:"amount"`
		Price       float64 `json:"price"`
	}
	if err := db.NewSelect().
		Table("cart_items").
		ColumnExpr("cart_items.product_id, products.name AS product_name, cart_items.total_product_amount AS amount, products.price").
		Join("JOIN products ON products.id = cart_items.product_id").
		Where("cart_id = ?", cartID).
		Scan(ctx, &cartItems); err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}

	totalPrice := 0.0
	totalAmount := 0
	for _, item := range cartItems {
		totalPrice += item.Price * float64(item.Amount)
		totalAmount += int(item.Amount)
	}

	order := &model.Orders{
		UserID:       req.UserID,
		ShipmentID:   req.ShipmentID,
		PaymentID:    req.PaymentID,
		Total_price:  totalPrice,
		Total_amount: totalAmount,
		Status:       "pending",
	}
	if _, err := db.NewInsert().Model(order).Returning("id").Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	for _, item := range cartItems {
		orderDetail := &model.OrderDetail{
			OrderID:            order.ID,
			ProductName:        item.ProductName,
			TotalProductPrice:  item.Price * float64(item.Amount),
			TotalProductAmount: int(item.Amount),
		}
		if _, err := db.NewInsert().Model(orderDetail).Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to create order detail: %v", err)
		}
	}

	if _, err := db.NewDelete().Table("cart_items").Where("cart_id = ?", cartID).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to delete cart items: %v", err)
	}
	if _, err := db.NewDelete().Table("carts").Where("id = ?", cartID).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to delete cart: %v", err)
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
