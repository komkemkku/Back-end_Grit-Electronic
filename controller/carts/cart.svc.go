package carts

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

func ListCartService(ctx context.Context, req requests.CartRequest) ([]response.CartResponses, int, float64, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.CartResponses{}
	var grandTotal float64 // ประกาศตัวแปร grandTotal สำหรับเก็บยอดรวมทั้งหมด

	// สร้าง query
	query := db.NewSelect().
		TableExpr("carts AS c").
		Column("c.id", "c.quantity", "c.created_at", "c.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("p.price AS product__price").
		ColumnExpr("p.image AS product__image").
		Join("LEFT JOIN products as p ON p.id = c.product_id")

	// คำนวณจำนวนรายการทั้งหมด (Count)
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	// ดึงข้อมูลด้วย Offset และ Limit
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, 0, err
	}

	// คำนวณ TotalPrice สำหรับแต่ละรายการ และ GrandTotal
	for i := range resp {
		resp[i].TotalPrice = float64(resp[i].Quantity) * resp[i].Product.Price
		grandTotal += resp[i].TotalPrice
	}

	return resp, total, grandTotal, nil
}


func GetByIdCartService(ctx context.Context, id int64) (*response.CartResponses, error) {

	ex, err := db.NewSelect().TableExpr("carts").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("cart not found")
	}
	product := &response.CartResponses{}

	err = db.NewSelect().TableExpr("carts AS c").
		Column("c.id", "c.quantity", "c.created_at", "c.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("p.price AS product__price").
		Join("LEFT JOIN products as p ON p.id = c.product_id").Where("c.id = ?", id).Scan(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func CreateCartService(ctx context.Context, req requests.CartAddItemRequest) (*model.Carts, error) {

	// ตรวจสอบว่าสินค้ามีอยู่หรือไม่ และดึงจำนวน stock
	var stock int
	err := db.NewSelect().
		TableExpr("products").
		Column("stock").
		Where("id = ?", req.ProductID).
		Scan(ctx, &stock)
	if err != nil {
		return nil, err
	}

	// หากสินค้าไม่มีในฐานข้อมูล
	if stock == 0 {
		return nil, errors.New("product not found")
	}

	// ตรวจสอบว่าสต็อกเพียงพอหรือไม่
	if req.Quantity > stock {
		return nil, fmt.Errorf("not enough stock")
	}

	// เพิ่มสินค้าใหม่ลงในตะกร้า
	cart := &model.Carts{
		ProductID: int(req.ProductID),
		Quantity:  req.Quantity,
	}
	cart.SetCreatedNow()
	cart.SetUpdateNow()

	_, err = db.NewInsert().Model(cart).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return cart, nil
}


func UpdateCartService(ctx context.Context, id int64, req requests.CartUpdateItemRequest) (*model.Carts, error) {
	ex, err := db.NewSelect().TableExpr("carts").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("cart not found")
	}

	cart := &model.Carts{}

	err = db.NewSelect().Model(cart).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	cart.Quantity = req.Quantity
	cart.ProductID = int(req.ProductID)
	cart.SetUpdateNow()

	_, err = db.NewUpdate().Model(cart).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func DeleteCartService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("carts").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("cart not found")
	}

	_, err = db.NewDelete().TableExpr("carts").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
