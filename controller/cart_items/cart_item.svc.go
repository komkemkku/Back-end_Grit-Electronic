package cartitems

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/image"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListCartItemService(ctx context.Context, req requests.CartItemRequest) ([]response.CartItemResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.CartItemResponses{}

	query := db.NewSelect().
		TableExpr("cart_items AS ci").
		Column("ci.id", "ci.cart_id", "ci.product_name", "ci.product_image_main", "ci.total_product_price", "ci.total_product_amount", "ci.status", "ci.created_at", "ci.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("p.price AS product__price").
		Join("LEFT JOIN products AS p ON p.id = ci.product_id")

	query.Order("ci.id ASC")

	if req.Search != "" {
		query.Where("ci.name ILIKE ?", "%"+req.Search+"%")
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

func GetByIdCartItemService(ctx context.Context, id int64) (*response.CartItemResponses, error) {
	ex, err := db.NewSelect().TableExpr("cart_items").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("cart_item not found")
	}
	cart_item := &response.CartItemResponses{}

	// แบบที่ 1
	err = db.NewSelect().TableExpr("cart_items AS ci").
		Column("ci.id", "ci.cart_id", "ci.product_name", "ci.product_image_main", "ci.total_product_price", "ci.total_product_amount", "ci.status", "ci.created_at", "ci.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("p.price AS product__price").
		Join("LEFT JOIN products AS p ON p.id = ci.product_id").
		Where("ci.id = ?", id).
		Scan(ctx, cart_item)

	if err != nil {
		return nil, err
	}
	return cart_item, nil
}

func CreateCartItemService(ctx context.Context, req requests.CartItemCreateRequest) (*model.CartItem, error) {
	// ตรวจสอบว่ามีตะกร้าของผู้ใช้อยู่หรือไม่
	cart := &model.Carts{}
	err := db.NewSelect().
		Model(cart).
		Where("user_id = ?", req.UserID).
		Where("status = ?", "pending").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// ถ้าไม่มีกระร้า ให้สร้างใหม่
			cart = &model.Carts{
				UserID:          req.UserID,
			}
			cart.SetCreatedNow()
			cart.SetUpdateNow()

			_, err = db.NewInsert().Model(cart).Exec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create cart: %v", err)
			}
		} else {
			// กรณี error อื่นๆ
			return nil, fmt.Errorf("failed to check cart: %v", err)
		}
	}

	// ตรวจสอบว่าสินค้าอยู่ใน `cart_items` หรือไม่
	cartItem := &model.CartItem{}
	exists, err := db.NewSelect().Model(cartItem).
		Where("cart_id = ?", cart.ID).
		Where("product_id = ?", req.ProductID).
		Exists(ctx)
	if err != nil {
		return nil, err
	}

	if exists {
		// ถ้ามีสินค้าอยู่แล้ว เพิ่มจำนวน
		cartItem.TotalProductAamount += req.TotalProductAmount
		cartItem.SetUpdateNow()

		_, err = db.NewUpdate().Model(cartItem).Where("id = ?", cartItem.ID).Exec(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		// ถ้าไม่มีสินค้า ให้เพิ่มใหม่
		cartItem = &model.CartItem{
			CartID:              cart.ID,
			ProductID:           req.ProductID,
			TotalProductAamount: req.TotalProductAmount,
			Status:              req.Status,
		}
		cartItem.SetCreatedNow()
		cartItem.SetUpdateNow()

		_, err = db.NewInsert().Model(cartItem).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to insert cart item: %v", err)
		}
	}

	// อัปเดต `cart` ตามจำนวนและราคาของสินค้าที่เพิ่มเข้ามา
	cart.SetUpdateNow()

	_, err = db.NewUpdate().Model(cart).Where("id = ?", cart.ID).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update cart: %v", err)
	}

	return cartItem, nil
}

func UpdateCartItemService(ctx context.Context, id int64, req requests.CartItemUpdateRequest) (*model.CartItem, error) {
	cartItem := &model.CartItem{}

	// exist ในการค้นหาแทน
	ex, err := db.NewSelect().TableExpr("products").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("product not found")
	}

	// อัปเดตข้อมูลสินค้าในตะกร้า
	cartItem.TotalProductAamount = req.TotalProductAmount
	cartItem.SetUpdateNow()

	_, err = db.NewUpdate().Model(cartItem).Where("cart_id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	img := requests.ImageCreateRequest{
		RefID:       cartItem.ID,
		Type:        "product_cart_item",
		Description: req.ProductImageMain,
	}

	_, err = image.CreateImageService(ctx, img)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func DeleteCartItemService(ctx context.Context, cartID, userID, cartItemID int) error {
	cartItem := &model.CartItem{}

	// ตรวจสอบว่า cart_item มีอยู่จริง
	// exist ในการค้นหาแทน

	err := db.NewSelect().
		Model(cartItem).
		Where("id = ? AND cart_id = ? AND EXISTS (SELECT 1 FROM carts WHERE id = ? AND user_id = ?)", cartItemID, cartID, cartID, userID).
		Scan(ctx)
	if err != nil {
		return errors.New("cart_item not found")
	}

	// ลบ cart_item
	_, err = db.NewDelete().
		Model(cartItem).
		Where("id = ?", cartItemID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete cart_item: %v", err)
	}

	// ตรวจสอบจำนวนสินค้าที่เหลือใน cart_item
	// 	query.count ในการนับแทน

	var itemCount int
	err = db.NewSelect().
		TableExpr("cart_items").
		ColumnExpr("COUNT(*)").
		Where("cart_id = ?", cartID).
		Scan(ctx, &itemCount)
	if err != nil {
		return fmt.Errorf("failed to check remaining cart_items: %v", err)
	}

	// หากไม่มีสินค้าเหลือใน cart ให้ลบ cart อัตโนมัติ
	if itemCount == 0 {
		_, err = db.NewDelete().
			TableExpr("carts").
			Where("id = ?", cartID).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to delete empty cart: %v", err)
		}
	}

	return nil
}
