package orders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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

	// สร้าง slice สำหรับ response
	resp := []response.OrderResponses{}

	// สร้าง CASE WHEN เพื่อกำหนดลำดับของ status
	caseStatement := "CASE " +
		"WHEN o.status = 'pending' THEN 1 " +
		"WHEN o.status = 'prepare' THEN 2 " +
		"WHEN o.status = 'ship' THEN 3 " +
		"WHEN o.status = 'success' THEN 4 " +
		"WHEN o.status = 'failed' THEN 5 " +
		"WHEN o.status = 'cancelled' THEN 6 " +
		"ELSE 7 END"

	// สร้าง query
	query := db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.payment_id", "o.total_price", "o.total_amount", "o.status").
		ColumnExpr("to_timestamp(o.created_at) AS created_at").
		ColumnExpr("to_timestamp(o.updated_at) AS updated_at").
		ColumnExpr("u.username").
		ColumnExpr("u.firstname AS user_firstname").
		ColumnExpr("u.lastname AS user_lastname").
		ColumnExpr("u.phone AS user_phone").
		ColumnExpr("s.id AS shipment_id").
		ColumnExpr("s.firstname AS shipment_firstname").
		ColumnExpr("s.lastname AS shipment_lastname").
		ColumnExpr("s.address AS shipment_address").
		ColumnExpr("s.zip_code AS shipment_zip_code").
		ColumnExpr("s.sub_district AS shipment_sub_district").
		ColumnExpr("s.district AS shipment_district").
		ColumnExpr("s.province AS shipment_province").
		Join("LEFT JOIN users AS u ON u.id = o.user_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id")

	// กรองตามคำค้นหาหรือ search ที่ชื่อผู้ใช้ (firstname หรือ lastname)
	if req.Search != "" {
		query.Where("u.firstname ILIKE ? OR u.lastname ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	// กรองตามช่วงวันที่เริ่มต้นและสิ้นสุด
	if req.StartDate != "" && req.EndDate != "" {
		query.Where("o.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}	

	// เพิ่มเงื่อนไขการเรียงข้อมูลตามลำดับ status และวันที่สร้างล่าสุด
	query.OrderExpr(fmt.Sprintf("%s, o.status DESC", caseStatement))

	// นับจำนวนทั้งหมด
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//  
	if total == 0 {
		return nil, 0, nil
	}

	// ดึงข้อมูลพร้อม pagination โดยใช้ offset และ limit
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}
  
	return resp, total, nil
  }


func ListOrderUserPendingService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	return ListOrderUserServiceByStatus(ctx, req, "pending")
}


func ListOrderUserPrepareService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	return ListOrderUserServiceByStatus(ctx, req, "prepare")
}


func ListOrderUserShipService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	return ListOrderUserServiceByStatus(ctx, req, "ship")
}


func ListOrderUserSuccessService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	return ListOrderUserServiceByStatus(ctx, req, "success")
}


func ListOrderUserFailedService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	
	return ListOrderUserServiceByStatus(ctx, req, "failed")
}

func ListOrderUserCancelledService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	return ListOrderUserServiceByStatus(ctx, req, "cancelled ")
}


func ListOrderUserHistoryService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
    var allResponses []response.OrderResponses
    total := 0

    successOrders, successTotal, err := ListOrderUserServiceByStatus(ctx, req, "success")
    if err != nil {
        return nil, 0, err
    }
    allResponses = append(allResponses, successOrders...)
    total += successTotal

    failedOrders, failedTotal, err := ListOrderUserServiceByStatus(ctx, req, "failed")
    if err != nil {
        return nil, 0, err
    }
    allResponses = append(allResponses, failedOrders...)
    total += failedTotal

    cancelledOrders, cancelledTotal, err := ListOrderUserServiceByStatus(ctx, req, "cancelled")
    if err != nil {
        return nil, 0, err
    }
    allResponses = append(allResponses, cancelledOrders...)
    total += cancelledTotal

    return allResponses, total, nil
}


// ฟังก์ชันที่ใช้ร่วมกันสำหรับการกรองตามสถานะ
func ListOrderUserServiceByStatus(ctx context.Context, req requests.OrderUserRequest, status string) ([]response.OrderResponses, int, error) {
	if req.UserID == 0 {
		return nil, 0, errors.New("invalid UserID")
	}

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.OrderResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.payment_id", "o.total_price", "o.total_amount", "o.status").
		ColumnExpr("to_timestamp(o.created_at) AS created_at").
		ColumnExpr("to_timestamp(o.updated_at) AS updated_at").
		ColumnExpr("u.username").
		ColumnExpr("u.firstname AS user_firstname").
		ColumnExpr("u.lastname AS user_lastname").
		ColumnExpr("u.phone AS user_phone").
		ColumnExpr("s.id AS shipment_id").
		ColumnExpr("s.firstname AS shipment_firstname").
		ColumnExpr("s.lastname AS shipment_lastname").
		ColumnExpr("s.address AS shipment_address").
		ColumnExpr("s.zip_code AS shipment_zip_code").
		ColumnExpr("s.sub_district AS shipment_sub_district").
		ColumnExpr("s.district AS shipment_district").
		ColumnExpr("s.province AS shipment_province").
		Join("LEFT JOIN users AS u ON u.id = o.user_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id").
		Where("o.status = ?", status) // กรองตามสถานะที่กำหนด


	if req.Search != "" {
		query.Where("o.status ILIKE ?", "%"+req.Search+"%")
	}

	if req.UserID != 0 {
		query.Where("o.user_id = ?", req.UserID)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}


  func GetByIdOrderService(ctx context.Context, orderID int64) (*response.OrderRespOrderDetail, error) {
	// 1) ตรวจสอบว่าคำสั่งซื้อนั้นมีอยู่ในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().
		Table("orders").
		Where("id = ?", orderID).
		Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	// 2) สร้าง response object
	order := &response.OrderRespOrderDetail{}

	// 3) ดึงข้อมูลจากตาราง orders และตารางที่เกี่ยวข้อง (User, Payment, Shipment, ... )
	err = db.NewSelect().
		TableExpr("orders AS o").
		// เพิ่ม "o.tracking_number" ใน Column หรือ ColumnExpr
		Column("o.id", "o.total_price", "o.total_amount", "o.status", "o.tracking_number", "o.created_at", "o.updated_at").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.firstname AS user__firstname").
		ColumnExpr("u.lastname AS user__lastname").
		ColumnExpr("u.phone AS user__phone").
		ColumnExpr("COALESCE(py.id, NULL) AS payment__id").
		ColumnExpr("COALESCE(py.price, NULL) AS payment__price").
		ColumnExpr("py.bank_name AS payment__bank_name").
		ColumnExpr("py.account_name AS payment__account_name").
		ColumnExpr("py.account_number AS payment__account_number").
		ColumnExpr("COALESCE(py.status, NULL) AS payment__status").
		ColumnExpr("sb.id AS system_bank__id").
		ColumnExpr("sb.bank_name AS system_bank__bank_name").
		ColumnExpr("sb.account_name AS system_bank__account_name").
		ColumnExpr("sb.account_number AS system_bank__account_number").
		ColumnExpr("sb.description AS system_bank__description").
		ColumnExpr("sb.image AS system_bank__image").
		ColumnExpr("s.id AS shipment__id").
		ColumnExpr("s.firstname AS shipment__firstname").
		ColumnExpr("s.lastname AS shipment__lastname").
		ColumnExpr("s.address AS shipment__address").
		ColumnExpr("s.zip_code AS shipment__zip_code").
		ColumnExpr("s.sub_district AS shipment__sub_district").
		ColumnExpr("s.district AS shipment__district").
		ColumnExpr("s.province AS shipment__province").
		Join("LEFT JOIN users AS u ON u.id = o.user_id").
		Join("LEFT JOIN payments AS py ON py.id = o.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = o.shipment_id").
		Join("LEFT JOIN system_banks AS sb ON sb.id = py.system_bank_id").
		Where("o.id = ?", orderID).
		Scan(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order details: %v", err)
	}

	// 4) ดึงชื่อสินค้าจากตาราง order_details (กรณีออเดอร์มีแค่ 1 สินค้า)
	//    ถ้ามีหลายสินค้าและต้องการหลายชื่อ อาจต้องปรับเป็น slice แทน
	var productItems []struct {
		ProductName string `bun:"product_name"`
	}
	err = db.NewSelect().
		Table("order_details").
		Column("product_name").
		Where("order_id = ?", orderID).
		Scan(ctx, &productItems)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to fetch product names: %v", err)
	}

	// 5) ใส่ product_name หลายค่าใน order.Products
	for _, p := range productItems {
		order.Products = append(order.Products, p.ProductName)
	}

	return order, nil
}
  

func CreateOrderService(ctx context.Context, req requests.OrderCreateRequest) (*model.Orders, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	// 1. ดึงข้อมูลตะกร้าสินค้า
	var cartID int64
	if err := tx.NewSelect().Table("carts").Column("id").Where("user_id = ?", req.UserID).Scan(ctx, &cartID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no cart found for user_id: %d", req.UserID)
		}
		return nil, fmt.Errorf("failed to find cart: %v", err)
	}

	// 2. ดึงข้อมูลสินค้าในตะกร้า
	var cartItems []struct {
		ProductID   int64   `json:"product_id"`
		ProductName string  `json:"product_name"`
		Amount      int64   `json:"amount"`
		Price       float64 `json:"price"`
		Stock       int64   `json:"stock"`
	}
	if err := tx.NewSelect().Table("cart_items").
		ColumnExpr("cart_items.product_id, products.name AS product_name, cart_items.total_product_amount AS amount, products.price, products.stock").
		Join("JOIN products ON products.id = cart_items.product_id").
		Where("cart_id = ?", cartID).
		Scan(ctx, &cartItems); err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}

	// 3. ตรวจสอบสต็อกสินค้า
	for _, item := range cartItems {
		if item.Amount > item.Stock {
			return nil, fmt.Errorf("not enough stock for product %s", item.ProductName)
		}
	}

	// 4. คำนวณราคาและจำนวนรวม
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
		Status:       "pending", // กำหนดสถานะคำสั่งซื้อ
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	// 5. สร้างคำสั่งซื้อในฐานข้อมูล
	if _, err := tx.NewInsert().Model(order).Returning("id").Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// 6. ถ้าสถานะไม่ใช่ "cancelled", หักสต็อกสินค้า
	for _, item := range cartItems {
		if _, err := tx.NewUpdate().
			Table("products").
			Set("stock = stock - ?", item.Amount).
			Where("id = ?", item.ProductID).
			Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to update stock for product %s: %v", item.ProductName, err)
		}
	}

	// 7. สร้างรายละเอียดคำสั่งซื้อ
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

	// 8. เคลียร์ตะกร้าสินค้า
	if _, err := tx.NewDelete().Table("cart_items").Where("cart_id = ?", cartID).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to delete cart items: %v", err)
	}
	if _, err := tx.NewDelete().Table("carts").Where("id = ?", cartID).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to delete cart: %v", err)
	}

	// 9. คอมมิตการทำธุรกรรม
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return order, nil
}

func UpdateOrderService(ctx context.Context, id int64, req requests.OrderUpdateRequest) (*model.Orders, error) {
	// 1) เช็กว่า Order นี้มีอยู่ในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().
		TableExpr("orders").
		Where("id = ?", id).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	// 2) ดึงข้อมูล Order จาก DB
	order := &model.Orders{}
	err = db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	// ตรวจสอบว่ากรณีเรากำลังจะเปลี่ยนไปเป็น cancelled แต่สถานะปัจจุบันเป็น "ship" หรือ "shipped"
	// ถ้าใช่ ให้ return error ไม่ให้ทำงานต่อ
	if (order.Status == "ship" || order.Status == "success") && req.Status == "cancelled" {
		return nil, errors.New("cannot cancel an order that is in ship or already shipped")
	}

	// 3) อัปเดต Status ตาม request
	order.Status = req.Status
	order.SetUpdateNow() // ฟังก์ชันกำหนด updated_at ใน struct

	// // 4) ถ้าสถานะเป็น "ship" ให้บันทึก TrackingNumber ด้วย
	if req.Status == "ship" {
		// ตรวจสอบว่า TrackingNumber ถูกตั้งค่าหรือไม่
		if req.TrackingNumber == "" {
			return nil, errors.New("tracking number must be provided when the order is ship")
		}
		order.TrackingNumber = req.TrackingNumber
		_, err = db.NewUpdate().Model(order).Column("status", "tracking_number", "updated_at").Where("id = ?", id).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update order: %v", err)
		}
	} else if req.Status != "ship" && req.TrackingNumber != "" {
		// ถ้าสถานะไม่ใช่ "ship" จะไม่สามารถอัปเดต TrackingNumber ได้
		return nil, errors.New("cannot set tracking number when order status is not ship")
	} else {
		// ถ้าเป็นสถานะอื่น ๆ ก็อัปเดตเฉพาะ status และ updated_at
		_, err = db.NewUpdate().
			Model(order).
			Column("status", "updated_at").
			Where("id = ?", id).
			Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update order: %v", err)
		}
	}

	// 5) ถ้าสถานะเป็น "cancelled" ให้คืนสินค้าเข้าคลัง
	if order.Status == "cancelled" {
		// ดึงรายละเอียดคำสั่งซื้อที่เกี่ยวข้อง
		var orderDetails []struct {
			ProductName string `bun:"product_name"`
			Amount      int64  `bun:"total_product_amount"`
		}
		if err := db.NewSelect().
			TableExpr("order_details AS od").
			ColumnExpr("od.product_name,od.total_product_amount").
			Where("od.order_id = ?", order.ID).
			Scan(ctx, &orderDetails); err != nil {
			return nil, fmt.Errorf("failed to fetch order details for cancellation: %v", err)
		}

		log.Printf("%v", orderDetails)

		// คืนสินค้าในกรณีที่ยกเลิกคำสั่งซื้อ
		for _, item := range orderDetails {
			// เพิ่มสินค้าในสต็อกกลับ
			if _, err := db.NewUpdate().
				Table("products").
				Set("stock = stock + ?", item.Amount).
				Where("name = ?", item.ProductName).
				Exec(ctx); err != nil {
				return nil, fmt.Errorf("failed to restore stock for product %v: %v", item.ProductName, err)
			}
		}

	}

	return order, nil
}

// func DeleteOrderService(ctx context.Context, id int64) error {
// 	ex, err := db.NewSelect().TableExpr("orders").Where("id=?", id).Exists(ctx)

// 	if err != nil {
// 		return err
// 	}

// 	if !ex {
// 		return errors.New("order not found")
// 	}

// 	_, err = db.NewDelete().TableExpr("orders").Where("id =?", id).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
