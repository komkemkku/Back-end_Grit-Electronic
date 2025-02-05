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
	// คำนวณ offset สำหรับ pagination
	var offset int
	if req.Page > 0 {
		offset = int((req.Page - 1) * req.Size)
	}

	// สร้าง slice สำหรับ response
	resp := []response.OrderResponses{}

	// สร้าง query หลัก
	query := db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id", "u.username", "o.status", "o.created_at", "o.updated_at", "o.total_price", "o.total_amount").
		ColumnExpr("py.system_bank_id, py.price AS payment_price, py.bank_name, py.account_name, py.account_number, py.status AS payment_status").
		ColumnExpr("s.firstname, s.lastname, s.address, s.zip_code, s.sub_district, s.district, s.province, s.status AS shipment_status").
		Join("LEFT JOIN users AS u ON u.id = o.user_id"). 
		Join("LEFT JOIN payments AS py ON py.id = o.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id")

	// เงื่อนไขการค้นหา
	if req.Search != "" {
		query.Where("o.status ILIKE ?", "%"+req.Search+"%")
	}

	// สร้าง query สำหรับนับจำนวนทั้งหมด
	countQuery := db.NewSelect().
		TableExpr("orders AS o")
	if req.Search != "" {
		countQuery.Where("o.status ILIKE ?", "%"+req.Search+"%")
	}
	total, err := countQuery.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count orders: %v", err)
	}

	// ดึงข้อมูลพร้อม pagination
	err = query.Offset(offset).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch orders: %v", err)
	}

	// ส่ง response กลับ
	return resp, total, nil
}
func GetByIdOrderService(ctx context.Context, orderID int64) (*response.OrderResponses, error) {
	// ตรวจสอบว่าคำสั่งซื้อนั้นมีอยู่ในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().
		Table("orders").
		Where("id = ?", orderID). // ใช้ order_id แทน user_id
		Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	// สร้าง response object
	order := &response.OrderResponses{}

	// ดึงข้อมูลจากตาราง orders และข้อมูลที่เกี่ยวข้อง
	err = db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id","username","o.status", "o.created_at", "o.updated_at", "o.total_price", "o.total_amount").
		ColumnExpr("py.system_bank_id, py.price AS payment_price, py.bank_name, py.account_name, py.account_number, py.status AS payment_status").
		ColumnExpr("s.firstname, s.lastname, s.address, s.zip_code, s.sub_district, s.district, s.province, s.status AS shipment_status").
		Join("LEFT JOIN users AS u ON u.id = o.user_id").
		Join("LEFT JOIN payments AS py ON py.id = o.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id").
		Where("o.id = ?", orderID). // ใช้ order_id แทน
		Scan(ctx, order)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch order details: %v", err)
	}

	return order, nil
}

func CreateOrderService(ctx context.Context, req requests.OrderCreateRequest) (*model.Orders, error) {
	// เริ่ม Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	var cartID int64
	if err := tx.NewSelect().Table("carts").Column("id").Where("user_id = ?", req.UserID).Scan(ctx, &cartID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no cart found for user_id: %d", req.UserID)
		}
		return nil, fmt.Errorf("failed to find cart: %v", err)
	}

	var cartItems []struct {
		ProductID   int64   `json:"product_id"`
		ProductName string  `json:"product_name"`
		Amount      int64   `json:"amount"`
		Price       float64 `json:"price"`
		Stock       int64   `json:"stock"`
	}
	if err := tx.NewSelect().
		Table("cart_items").
		ColumnExpr("cart_items.product_id, products.name AS product_name, cart_items.total_product_amount AS amount, products.price, products.stock").
		Join("JOIN products ON products.id = cart_items.product_id").
		Where("cart_id = ?", cartID).
		Scan(ctx, &cartItems); err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}

	// ลดหรือเพิ่ม stock ของสินค้า
	// ลด stock ของสินค้าในกรณีที่ไม่ใช่การยกเลิกคำสั่งซื้อ
for _, item := range cartItems {
	if req.Status == "canceled" { // ตรวจสอบสถานะคำสั่งซื้อ
		// เพิ่ม stock เมื่อยกเลิก Order
		if _, err := tx.NewUpdate().Table("products").
			Set("stock = stock + ?", item.Amount). // เพิ่มจำนวน stock
			Where("id = ?", item.ProductID).
			Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to restore stock for product %s: %v", item.ProductName, err)
		}
	} else {
		// ลด stock ของสินค้าเมื่อทำการสร้าง Order ใหม่
		if _, err := tx.NewUpdate().Table("products").
			Set("stock = stock - ?", item.Amount). // ลดจำนวน stock
			Where("id = ?", item.ProductID).
			Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to update stock for product %s: %v", item.ProductName, err)
		}
	}
}


	// คำนวณราคาทั้งหมด
	totalPrice := 0.0
	totalAmount := 0
	for _, item := range cartItems {
		totalPrice += item.Price * float64(item.Amount)
		totalAmount += int(item.Amount)
	}

	// สร้างคำสั่งซื้อ
	order := &model.Orders{
		UserID:       req.UserID,
		ShipmentID:   req.ShipmentID,
		PaymentID:    req.PaymentID,
		Total_price:  totalPrice,
		Total_amount: totalAmount,
		Status:       req.Status, // ใช้สถานะจาก request
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	if _, err := tx.NewInsert().Model(order).Returning("id").Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// เพิ่ม Order Detail
	for _, item := range cartItems {
		orderDetail := &model.OrderDetail{
			OrderID:            order.ID,
			ProductName:        item.ProductName,
			TotalProductPrice:  item.Price * float64(item.Amount),
			TotalProductAmount: int(item.Amount),
		}
		if _, err := tx.NewInsert().Model(orderDetail).Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to create order detail: %v", err)
		}
	}

	// ลบรายการในตะกร้า
	if _, err := tx.NewDelete().Table("cart_items").Where("cart_id = ?", cartID).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to delete cart items: %v", err)
	}
	if _, err := tx.NewDelete().Table("carts").Where("id = ?", cartID).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to delete cart: %v", err)
	}

	// คอมมิตการทำธุรกรรม
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return order, nil
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
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	// อัปเดตแค่ status เท่านั้น
	if req.Status != "" {
		order.Status = req.Status
		order.SetUpdateNow() // ตั้งค่า UpdatedAt ถ้ามีการเปลี่ยนแปลง status
	} else {
		// หากไม่มีการเปลี่ยนแปลง status ก็ไม่ต้องอัปเดต updated_at
		return nil, errors.New("status is empty, no update performed")
	}

	// บันทึกข้อมูลกลับไปยังฐานข้อมูล โดยอัปเดตแค่ status และ updated_at
	_, err = db.NewUpdate().
		Model(order).
		Column("status", "updated_at").
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
