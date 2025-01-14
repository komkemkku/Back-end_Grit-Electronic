package orederdetail

import (
	"context"
	"errors"
	"fmt"

	// "strconv"

	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = config.Database()

func ListOrderDetailService(ctx context.Context, req requests.OrderDetailRequest) (response.OrderDetailResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := response.OrderDetailResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("order_details AS od").
		Column("od.id", "od.created_at", "od.updated_at").
		ColumnExpr("o.id AS order__id").
		ColumnExpr("o.total_price AS total_price").
		ColumnExpr("o.total_amount AS total_amount").
		ColumnExpr("o.status AS status").
		ColumnExpr("pr.id AS product__id").
		ColumnExpr("pr.name AS product__name").
		ColumnExpr("pr.price AS product__price").
		ColumnExpr("pay.id AS payment__id").
		ColumnExpr("pay.price AS payment__price").
		ColumnExpr("pay.slip AS payment__slip").
		ColumnExpr("pay.status AS payment__status").
		ColumnExpr("s.id AS shipment__id").
		ColumnExpr("s.firstname AS shipment__firstname").
		ColumnExpr("s.lastname AS shipment__lastname").
		ColumnExpr("s.address AS shipment__address").
		ColumnExpr("s.zip_code AS shipment__zip_code").
		ColumnExpr("s.sub_district AS shipment__sub_district").
		ColumnExpr("s.district AS shipment__district").
		ColumnExpr("s.province AS shipment__province").
		ColumnExpr("s.status AS shipment__status").
		Join("LEFT JOIN orders AS o ON o.id = od.order_id").
		Join("LEFT JOIN payments AS pay ON pay.id = od.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = od.shipment_id").
		Join("LEFT JOIN products AS pr ON pr.id = od.product_id")

	// if req.Search != "" {
	// 	query.Where("o.product_id::text ILIKE ?", "%"+req.Search+"%")
	// }

	total, err := query.Count(ctx)
	if err != nil {
		return response.OrderDetailResponses{}, 0, err
	}

	// Execute query
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return response.OrderDetailResponses{}, 0, err
	}

	return resp, total, nil

}

func CreateOrderDetailService(ctx context.Context, req requests.OrderDetailCreateRequest) (*model.Order_details, error) {

	order_detail := &model.Order_details{
		OrderID:    req.OrderID,
		ProductID:  req.ProductID,
		PaymentID:  req.PaymentID,
		ShipmentID: req.ShipmentID,
	}

	order_detail.SetCreatedNow()
	order_detail.SetUpdateNow()

	_, err := db.NewInsert().Model(order_detail).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return order_detail, nil

}

func DeleteOrderDetailService(ctx context.Context, id int64) error {
	// ตรวจสอบว่า Wishlist มีอยู่หรือไม่
	ex, err := db.NewSelect().TableExpr("order_details").Where("id = ?", id).Exists(ctx)

	if err != nil {
		// กรณีเกิดข้อผิดพลาดจากฐานข้อมูล
		return err
	}

	if !ex {
		// กรณี Wishlist ไม่พบในฐานข้อมูล
		return errors.New("orderdetail not found")
	}

	// ลบ Wishlist ที่พบในฐานข้อมูล
	_, err = db.NewDelete().TableExpr("order_details").Where("id = ?", id).Exec(ctx)
	if err != nil {
		// กรณีลบไม่สำเร็จ
		return err
	}

	// สำเร็จ
	return nil
}

func GetByOrderDetailService(ctx context.Context, id int64) (*response.OrderDetailResponses, error) {
	exists, err := db.NewSelect().
		TableExpr("order_details").
		Where("id = ?", id).
		Exists(ctx)

	if err != nil {
		return nil, errors.New("order details not found")
	}

	if !exists {
		return nil, errors.New("order details not found")
	}

	orderdetail := &response.OrderDetailResponses{}

	err = db.NewSelect().
		TableExpr("order_details AS od").
		Column("od.id", "od.created_at", "od.updated_at").
		ColumnExpr("o.id AS order__id").
		ColumnExpr("o.total_price AS total_price").
		ColumnExpr("o.total_amount AS total_amount").
		ColumnExpr("o.status AS status").
		ColumnExpr("pr.id AS product__id").
		ColumnExpr("pr.name AS product__name").
		ColumnExpr("pr.price AS product__price").
		ColumnExpr("pay.id AS payment__id").
		ColumnExpr("pay.price AS payment__price").
		ColumnExpr("pay.slip AS payment__slip").
		ColumnExpr("pay.status AS payment__status").
		ColumnExpr("s.id AS shipment__id").
		ColumnExpr("s.firstname AS shipment__firstname").
		ColumnExpr("s.lastname AS shipment__lastname").
		ColumnExpr("s.address AS shipment__address").
		ColumnExpr("s.zip_code AS shipment__zip_code").
		ColumnExpr("s.sub_district AS shipment__sub_district").
		ColumnExpr("s.district AS shipment__district").
		ColumnExpr("s.province AS shipment__province").
		ColumnExpr("s.status AS shipment__status").
		Join("LEFT JOIN orders AS o ON o.id = od.order_id").
		Join("LEFT JOIN payments AS pay ON pay.id = od.payment_id").
		Join("LEFT JOIN shipments AS s ON s.id = od.shipment_id").
		Join("LEFT JOIN products AS pr ON pr.id = od.product_id").
		Where("od.id = ?", id).
		Scan(ctx, orderdetail)

	if err != nil {
		return nil, err
	}

	// ตรวจสอบ TotalPrice กับ Payment.Amount
	// if orderdetail.TotalPrice != orderdetail.Payment.Price {
	// 	return nil, fmt.Errorf("mismatch: total price (%.2f) does not match payment amount (%.2f)", orderdetail.TotalPrice, orderdetail.Payment.Price)
	// }

	return orderdetail, nil
}

func UpdateOrderDetailService(ctx context.Context, id int64, req requests.OrderDetailUpdateRequest) (*model.Order_details, error) {
	// ตรวจสอบว่า Order Detail มีอยู่ในระบบหรือไม่
	var Order_details model.Order_details
	err := db.NewSelect().
		Model(&Order_details).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve order_detail: %v", err)
	}

	// อัปเดตข้อมูลใน OrderDetails
	Order_details.ProductID = req.ProductID
	Order_details.PaymentID = req.PaymentID
	Order_details.ShipmentID = req.ShipmentID
	Order_details.SetUpdateNow()

	// บันทึกข้อมูลลงในฐานข้อมูล
	_, err = db.NewUpdate().
		Model(&Order_details).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update order_detail: %v", err)
	}

	return &Order_details, nil
}
