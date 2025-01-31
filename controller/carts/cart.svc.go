package carts

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

// func ListCartService(ctx context.Context, req requests.CartRequest) ([]response.CartResponses, int, error) {
// 	var Offset int64
// 	if req.Page > 0 {
// 		Offset = (req.Page - 1) * req.Size
// 	}

// 	resp := []response.CartResponses{}

// 	// สร้าง query พร้อมเชื่อมโยงกับ cart_items
// 	query := db.NewSelect().
// 		TableExpr("carts AS c").
// 		Column("c.id", "c.user_id", "c.total_cart_amount", "c.total_cart_price", "c.status", "c.created_at", "c.updated_at").
// 		ColumnExpr("json_agg(json_build_object('id', ci.id, 'product_id', ci.product_id, 'product_name', ci.product_name, 'product_image_main', ci.product_image_main, 'total_product_price', ci.total_product_price, 'total_product_amount', ci.total_product_amount)) AS cart_items").
// 		Join("LEFT JOIN cart_items AS ci ON ci.cart_id = c.id").
// 		GroupExpr("c.id")

// 	total, err := query.Count(ctx)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	// Execute query
// 	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return resp, total, nil
// }

func GetByIdCartService(ctx context.Context, id int64) (*response.CartResponses, error) {
	// ตรวจสอบว่ามีตะกร้าอยู่หรือไม่
	ex, err := db.NewSelect().TableExpr("carts").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("cart not found")
	}

	cart := &response.CartResponses{}

	// Query ตะกร้าพร้อมรายการสินค้า
	err = db.NewSelect().
		TableExpr("carts AS c").
		Column("c.id", "c.created_at", "c.updated_at").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.username AS user__username").
		ColumnExpr(`
		COALESCE(
			json_agg(
				json_build_object(
					'id', ci.id,
					'product', json_build_object(
						'id', p.id,
						'name', p.name,
						'price', p.price
					),
					'total_product_amount', ci.total_product_amount
				)
				ORDER BY ci.id ASC
			) FILTER (WHERE ci.id IS NOT NULL), '[]'
		) AS cart_items
	`).
		Join("LEFT JOIN cart_items AS ci ON ci.cart_id = c.id").
		Join("LEFT JOIN users AS u ON u.id = c.user_id").
		Join("LEFT JOIN products AS p ON p.id = ci.product_id").
		GroupExpr("c.id, u.id").
		Where("c.id = ?", id).
		Scan(ctx, cart)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func CreateCartService(ctx context.Context, req requests.CartAddItemRequest) (*model.Carts, error) {
	// ตรวจสอบว่าตะกร้าของผู้ใช้งานมีอยู่แล้วหรือไม่
	exists, err := db.NewSelect().
		TableExpr("carts").
		Where("user_id = ?", req.UserID).
		Exists(ctx)
	if err != nil {
		return nil, err
	}

	// ถ้ายังไม่มีตะกร้าสำหรับผู้ใช้นี้ ให้สร้างใหม่
	var cart *model.Carts
	if !exists {
		cart = &model.Carts{
			UserID: req.UserID,
		}
		cart.SetCreatedNow()
		cart.SetUpdateNow()

		_, err := db.NewInsert().Model(cart).Exec(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		// ดึงข้อมูลตะกร้าที่มีอยู่แล้ว
		cart = &model.Carts{}
		err = db.NewSelect().
			Model(cart).
			Where("user_id = ?", req.UserID).
			Scan(ctx)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}

func UpdateCartService(ctx context.Context, userID int64, req requests.CartUpdateItemRequest) (*model.Carts, error) {
	// ตรวจสอบว่ามีตะกร้าของผู้ใช้นี้อยู่หรือไม่
	cart := &model.Carts{}
	err := db.NewSelect().
		Model(cart).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, errors.New("cart not found")
	}

	// อัปเดตรายละเอียดตะกร้า

	cart.SetUpdateNow()

	_, err = db.NewUpdate().
		Model(cart).
		Where("user_id = ?", userID).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func DeleteCartService(ctx context.Context, userID int64) error {
	// ตรวจสอบว่ามีตะกร้าของผู้ใช้นี้อยู่หรือไม่
	exists, err := db.NewSelect().
		TableExpr("carts").
		Where("user_id = ?", userID).
		Exists(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("cart not found")
	}

	// ลบตะกร้าเมื่อไม่มีสินค้าเหลือ
	_, err = db.NewDelete().
		TableExpr("carts").
		Where("user_id = ? AND total_cart_amount = 0", userID).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
