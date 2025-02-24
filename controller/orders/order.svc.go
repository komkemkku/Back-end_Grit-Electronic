package orders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

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

	// สร้าง slice สำหรับ response
	resp := []response.OrderResponses{}

	var startUnix, endUnix int64

	if req.StartDate > 0 {
		startUnix = req.StartDate
	}

	if req.EndDate > 0 {
		endUnix = req.EndDate
	}

	// สร้าง CASE WHEN เพื่อกำหนดลำดับของ status
	caseStatement := "CASE " +
		"WHEN o.status = 'pending' THEN 1 " +
		"WHEN o.status = 'paid' THEN 2 " +
		"WHEN o.status = 'prepare' THEN 3 " +
		"WHEN o.status = 'ship' THEN 4 " +
		"WHEN o.status = 'success' THEN 5 " +
		"WHEN o.status = 'failed' THEN 6 " +
		"WHEN o.status = 'cancelled' THEN 7 " +
		"ELSE 8 END"

	// สร้าง query
	query := db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.user_id", "o.payment_id", "o.total_price", "o.total_price_ship", "o.total_amount", "o.status").
		ColumnExpr("floor(EXTRACT(EPOCH FROM to_timestamp(o.created_at)))::bigint AS created_at").
		ColumnExpr("floor(EXTRACT(EPOCH FROM to_timestamp(o.updated_at)))::bigint AS updated_at").
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

	// กรองตามคำค้นหา (ชื่อ, นามสกุล, หรือ order ID)
	if req.Search != "" {
		query.Where("u.firstname ILIKE ? OR u.lastname ILIKE ? OR CAST(o.id AS TEXT) ILIKE ?",
			"%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if req.Status != "" {
		query.Where("o.status = ?", req.Status)
	}

	// กรองตามช่วงวันที่ (Unix Timestamp)
	if startUnix > 0 && endUnix > 0 {
		if startUnix == endUnix {
			query.Where("DATE(TO_TIMESTAMP(o.created_at)) = DATE(TO_TIMESTAMP(?))", startUnix)
		} else {
			query.Where("o.created_at >= ? AND o.created_at <= ?", startUnix, endUnix)
		}
	} else if startUnix > 0 {
		query.Where("DATE(TO_TIMESTAMP(o.created_at)) = DATE(TO_TIMESTAMP(?))", startUnix)
	}

	// เพิ่มเงื่อนไขการเรียงข้อมูลตามลำดับ status และวันที่สร้างล่าสุด
	query.OrderExpr(fmt.Sprintf("%s, o.status DESC", caseStatement))

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 🔹 ถ้าไม่มีข้อมูลให้ส่งค่าเป็นอาร์เรย์ว่าง
	if total == 0 {
		return []response.OrderResponses{}, 0, nil
	}

	// ดึงข้อมูลพร้อม pagination โดยใช้ offset และ limit
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func ListOrderUserPendingService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	orders, total, err := ListOrderUserServiceByStatus(ctx, req, "pending")
	if err != nil {
		return nil, 0, err
	}

	for i, order := range orders {
		var paymentDate string
		err := db.NewSelect().
			TableExpr("payments").
			Column("date").
			Where("id = ?", order.PaymentID).
			Scan(ctx, &paymentDate)

		if err == nil && paymentDate != "" {
			// อัปเดตสถานะเป็น "prepare"
			_, err := db.NewUpdate().
				TableExpr("orders").
				Set("status = ?", "paid").
				Where("id = ?", order.ID).
				Exec(ctx)

			if err == nil {
				orders[i].Status = "prepare"
			}
		}
	}

	return orders, total, nil
}

func ListOrderUserPaidService(ctx context.Context, req requests.OrderUserRequest) ([]response.OrderResponses, int, error) {
	return ListOrderUserServiceByStatus(ctx, req, "paid")
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
		Column("o.id", "o.user_id", "o.payment_id", "o.total_price", "o.total_price_ship", "o.total_amount", "o.status").
		ColumnExpr("floor(EXTRACT(EPOCH FROM to_timestamp(o.created_at)))::bigint AS created_at").
		ColumnExpr("floor(EXTRACT(EPOCH FROM to_timestamp(o.updated_at)))::bigint AS updated_at").
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
		Where("o.status = ?", status)

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

	err = query.OrderExpr("o.created_at desc").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdOrderService(ctx context.Context, orderID, UserID int64) (*response.OrderRespOrderDetail, error) {
	// 1) ตรวจสอบว่าคำสั่งซื้อนั้นมีอยู่หรือไม่
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

	// 3) ดึงข้อมูลหลักของ order
	err = db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.total_price", "o.total_price_ship", "o.total_amount", "o.status", "o.tracking_number", "o.created_at", "o.updated_at").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.firstname AS user__firstname").
		ColumnExpr("u.lastname AS user__lastname").
		ColumnExpr("u.phone AS user__phone").
		ColumnExpr("py.id AS payment__id").
		ColumnExpr("py.date AS payment__date").
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
		Where("o.id = ?", orderID).
		Scan(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order details: %v", err)
	}

	// 4) ดึงข้อมูลสินค้า (product_name, price, image, total_product_amount)
	var productItems []response.ProductInfo
	err = db.NewSelect().
		TableExpr("order_details AS od").
		ColumnExpr("p.id AS product_id, od.product_name, p.image, p.price, od.total_product_amount").
		ColumnExpr(`CASE
			WHEN EXISTS(SELECT * FROM reviews WHERE product_id = p.id AND user_id = ?) THEN true
			ELSE false END AS is_review`, UserID).
		Join("LEFT JOIN products AS p ON p.name = od.product_name").
		Where("od.order_id = ?", orderID).
		Scan(ctx, &productItems)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product details: %v", err)
	}

	// 5) เพิ่มสินค้าลงใน response
	order.Products = productItems

	return order, nil
}

func GetUserByIdOrderService(ctx context.Context, orderID int64) (*response.OrderRespOrderDetail, error) {
	// 1) ตรวจสอบว่าคำสั่งซื้อนั้นมีอยู่หรือไม่
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

	// 3) ดึงข้อมูลหลักของ order
	err = db.NewSelect().
		TableExpr("orders AS o").
		Column("o.id", "o.total_price", "o.total_price_ship", "o.total_amount", "o.status", "o.tracking_number", "o.created_at", "o.updated_at").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.firstname AS user__firstname").
		ColumnExpr("u.lastname AS user__lastname").
		ColumnExpr("u.phone AS user__phone").
		ColumnExpr("py.id AS payment__id").
		ColumnExpr("py.date AS payment__date").
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
		Where("o.id = ?", orderID).
		Scan(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order details: %v", err)
	}

	// 4) ดึงข้อมูลสินค้า (product_name, price, image, total_product_amount)
	var productItems []response.ProductInfo
	err = db.NewSelect().
		TableExpr("order_details AS od").
		ColumnExpr("p.id AS product_id, od.product_name, p.image, p.price, od.total_product_amount").
		Join("JOIN products AS p ON p.name = od.product_name").
		Where("od.order_id = ?", orderID).
		Scan(ctx, &productItems)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product details: %v", err)
	}

	// 5) เพิ่มสินค้าลงใน response
	order.Products = productItems

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

	// 2. ดึง payment_id จาก payments โดยใช้วันที่ที่ลูกค้ากรอก
	// var paymentID int64
	// err = tx.NewSelect().Table("payments").
	// 	Column("id").
	// 	Where("date = ?", req.PaymentDate). // ใช้ req.PaymentDate เพื่อดึง payment_id
	// 	Limit(1).                           // เฉพาะ 1 รายการ
	// 	Scan(ctx, &paymentID)

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to fetch payment data: %v", err)
	// }

	// 3. ดึงข้อมูลสินค้าในตะกร้า
	var cartItems []struct {
		ID          int64   `json:"id"`
		ProductID   int64   `json:"product_id"`
		ProductName string  `json:"product_name"`
		Amount      int64   `json:"amount"`
		Price       float64 `json:"price"`
		Stock       int64   `json:"stock"`
	}

	items := []int64{}

	for _, item := range req.CartItems {
		items = append(items, item.CartID)
	}

	if err := tx.NewSelect().Table("cart_items").
		ColumnExpr("cart_items.id,cart_items.product_id, products.name AS product_name,cart_items.total_product_amount AS amount, products.price, products.stock").
		Join("JOIN products ON products.id = cart_items.product_id").
		Where("cart_id = ?", cartID).
		Where("cart_items.id IN (?)", bun.In(items)).
		Scan(ctx, &cartItems); err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
	}

	log.Println(cartItems)
	// 5. คำนวณราคารวมและจำนวนรวม
	totalPrice := 0.0
	totalAmount := 0
	for _, item := range cartItems {
		totalPrice += item.Price * float64(item.Amount)
		totalAmount += int(item.Amount)
	}

	// // 6. บันทึก Order ลงในฐานข้อมูล พร้อมเชื่อมโยงกับ Payment ID ที่ได้จากการค้นหา
	order := &model.Orders{
		UserID:       req.UserID,
		ShipmentID:   req.ShipmentID,
		PaymentID:    0,
		Total_price:  totalPrice,
		Total_amount: totalAmount,
		Status:       "pending",
	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	if _, err := tx.NewInsert().Model(order).Returning("id").Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// // 7. หักสต็อกสินค้า
	for _, item := range cartItems {
		if _, err := tx.NewUpdate().
			Table("products").
			Set("stock = stock - ?", item.Amount).
			Where("id = ?", item.ProductID).
			Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to update stock for product %s: %v", item.ProductName, err)
		}
	}

	// // 8. บันทึกรายละเอียด Order
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

		if _, err := tx.NewDelete().Table("cart_items").Where("cart_id = ?", cartID).Where("id = ?", item.ID).Exec(ctx); err != nil {
			return nil, fmt.Errorf("failed to delete cart items: %v", err)
		}
	}

	ex, err := tx.NewSelect().Table("cart_items").Where("cart_id = ?", cartID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		tx.NewDelete().Table("carts").Where("id = ?", cartID).Exec(ctx)
	}

	// 10. คอมมิตธุรกรรม
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
	// order.ShipmentID = req.ShipmentID
	order.SetUpdateNow()

	// 4) ถ้าสถานะเป็น "ship" ให้บันทึก TrackingNumber ด้วย
	if req.Status == "" {
		// ตรวจสอบว่า TrackingNumber ถูกตั้งค่าหรือไม่
		if req.TrackingNumber == "" {
			return nil, errors.New("tracking number must be provided when the order is ship")
		}
		order.TrackingNumber = req.TrackingNumber
		_, err = db.NewUpdate().Model(order).Column("status", "tracking_number", "updated_at").Where("id = ?", id).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update order: %v", err)
		}
	} else if req.Status != "" && req.TrackingNumber != "" {
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

	// 5) ถ้าสถานะเป็น "success" ให้บันทึกข้อมูลลงในตาราง report
	if order.Status == "success" {
		// ดึงรายละเอียดคำสั่งซื้อที่เกี่ยวข้อง
		var orderDetails []struct {
			ProductID   int64  `bun:"product_id"`
			ProductName string `bun:"product_name"`
			Amount      int64  `bun:"total_product_amount"`
		}

		// JOIN order_details กับ products เพื่อให้ได้ product_id
		err := db.NewSelect().
			TableExpr("order_details AS od").
			Join("JOIN products AS p ON od.product_name = p.name").
			ColumnExpr("p.id AS product_id, od.product_name, od.total_product_amount").
			Where("od.order_id = ?", order.ID).
			Scan(ctx, &orderDetails)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch order details: %v", err)
		}

		// ดึงข้อมูลสินค้าและประเภทสินค้า
		for _, item := range orderDetails {
			var product model.Products
			if err := db.NewSelect().
				Model(&product).
				Where("id = ?", item.ProductID).
				Scan(ctx); err != nil {
				return nil, fmt.Errorf("failed to fetch product for product_id %d: %v", item.ProductID, err)
			}

			// ดึงข้อมูลประเภทสินค้า
			var category model.Categories
			if err := db.NewSelect().
				Model(&category).
				Where("id = ?", product.CategoryID).
				Scan(ctx); err != nil {
				return nil, fmt.Errorf("failed to fetch category for product_id %d: %v", item.ProductID, err)
			}

			// บันทึกข้อมูลลงในตาราง report
			report := &model.Report{
				ProductID:   product.ID,
				ProductName: item.ProductName,
				Date:        time.Now().Unix(),                    // แปลงวันที่เป็น Unix timestamp
				TotalPrice:  product.Price * float64(item.Amount), // แปลง Amount เป็น float64 ก่อนคูณ
				TotalAmount: int(item.Amount),                     // แปลงเป็น int ถ้าจำเป็น
				Status:      order.Status,
			}

			// บันทึกข้อมูลลงใน database
			_, err := db.NewInsert().Model(report).Exec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to insert report record: %v", err)
			}
		}
	}

	// 6) ถ้าสถานะเป็น "cancelled" ให้คืนสินค้าเข้าคลัง
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

func UpdateShipOrderService(ctx context.Context, id int, req requests.OrderUpdateRequest) (*model.Orders, error) {
	// ตรวจสอบว่าคำสั่งซื้อมีอยู่หรือไม่
	exists, err := db.NewSelect().TableExpr("orders").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	// ดึงข้อมูลออเดอร์ปัจจุบัน
	order := &model.Orders{}
	err = db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่ามี Shipment ID นี้อยู่จริง
	shipExists, err := db.NewSelect().TableExpr("shipments").Where("id = ?", req.ShipmentID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !shipExists {
		return nil, errors.New("shipment not found")
	}

	// อัปเดต ShipmentID
	_, err = db.NewUpdate().
		TableExpr("orders").
		Set("shipment_id = ?", req.ShipmentID).
		Set("updated_at = EXTRACT(EPOCH FROM NOW())").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	// ดึงข้อมูลใหม่ที่อัปเดตแล้วกลับไป
	err = db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return order, nil
}
