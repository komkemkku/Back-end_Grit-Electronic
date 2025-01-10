package orederdetail

import (
	"context"
	"errors"

	// "strconv"

	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = config.Database()

func ListOrderDetailService(ctx context.Context, req requests.OrderDetailRequest) ([]response.OrderDetailResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.OrderDetailResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("order_details AS od").
		Column("od.id", "od.order_id", "od.product_id", "od.payment_id", "od.shipment_id", "od.created_at", "od.updated_at", "od.quantity", "od.unit_price")

	if req.Search != "" {
		query.Where("o.product_id::text ILIKE ?", "%"+req.Search+"%")
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

func CreateOrderDetailService(ctx context.Context, req requests.OrderDetailCreateRequest) (*model.Order_details, error) {

	order_detail := &model.Order_details{
		OrderID:    req.OrderID,
		ProductID:  req.ProductID,
		PaymentID:  req.PaymentID,
		ShipmentID: req.ShipmentID,
		Quantity:   req.Quantity,
		UnitPrice:  req.UnitPrice,
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
