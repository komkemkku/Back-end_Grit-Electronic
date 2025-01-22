package orders

import (
	"context"
	"errors"

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
		Column("o.id", "o.user_id","o.payment_id","o.shipment_id","o.cart_id","status","o.created_at", "o.updated_at")

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

func CreateOrderService(ctx context.Context, req requests.OrderCreateRequest) (*model.Orders, error) {

	order := &model.Orders{

	}
	order.SetCreatedNow()
	order.SetUpdateNow()

	_, err := db.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return order, nil

}

func UpdateOrderService(ctx context.Context, id int64, req requests.OrderUpdateRequest) (*model.Orders, error) {
	ex, err := db.NewSelect().TableExpr("orders").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("order not found")
	}

	order := &model.Orders{}

	err = db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	order.SetUpdateNow()

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
