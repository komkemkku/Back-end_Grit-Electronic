package cartitems

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

func ListCartItemService(ctx context.Context, req requests.CartItemRequest) ([]response.CartItemResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.CartItemResponses{}

	query := db.NewSelect().
		TableExpr("cart_items AS ci").
		Column("ci.id", "ci.cart_id", "ci.total_product_amount", "ci.status", "ci.created_at", "ci.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("p.price AS product__price").
		Join("LEFT JOIN products AS p ON p.id = ci.product_id")

	query.Order("ci.id ASC")

	// if req.Search != "" {
	// 	query.Where("ci.name ILIKE ?", "%"+req.Search+"%")
	// }

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

	err = db.NewSelect().TableExpr("cart_items AS ci").
		Column("ci.id", "ci.cart_id", "ci.total_product_amount", "ci.status", "ci.created_at", "ci.updated_at").
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
	// ตรวจสอบว่าสินค้ามีอยู่ในฐานข้อมูลหรือไม่
	product := &model.Products{}
	err := db.NewSelect().
		Model(product).
		Where("id = ?", req.ProductID).
		Where("is_active IS true").
		Scan(ctx)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if product.Stock < req.TotalProductAmount {
		return nil, errors.New("not enough stock")
	}

	if req.TotalProductAmount <= 0 {
		return nil, errors.New("the product amount must be greater than 1")
	}

	// ตรวจสอบว่าผู้ใช้มี cart อยู่หรือไม่
	cart := &model.Carts{}
	err = db.NewSelect().
		Model(cart).
		Where("user_id = ?", req.UserID).
		Scan(ctx)

	if err != nil {
		// ถ้าไม่มีตะกร้า ให้สร้างใหม่
		if errors.Is(err, sql.ErrNoRows) {
			cart = &model.Carts{
				UserID: req.UserID,
			}
			cart.SetCreatedNow()
			cart.SetUpdateNow()

			_, err = db.NewInsert().
				Model(cart).
				Returning("id").
				Exec(ctx, &cart.ID)
			if err != nil {
				return nil, errors.New("failed to insert cart")
			}
		} else {
			return nil, errors.New("failed to check cart")
		}
	}

	// ตรวจสอบว่าสินค้าอยู่ใน `cart_items` หรือไม่
	cartItem := &model.CartItem{}
	err = db.NewSelect().
		Model(cartItem).
		Where("cart_id = ?", cart.ID).
		Where("product_id = ?", req.ProductID).
		Scan(ctx)

	if err == nil {
		// ✅ ถ้าสินค้าอยู่ในตะกร้าแล้ว → อัปเดตจำนวนสินค้า
		newTotalAmount := cartItem.TotalProductAmount + req.TotalProductAmount

		// ตรวจสอบว่าสินค้าเกินสต็อกหรือไม่
		if newTotalAmount > product.Stock {
			return nil, errors.New("not enough stock")
		}

		cartItem.TotalProductAmount = newTotalAmount
		cartItem.SetUpdateNow()

		_, err = db.NewUpdate().
			Model(cartItem).
			Column("total_product_amount", "updated_at").
			Where("cart_id = ?", cart.ID).
			Where("product_id = ?", req.ProductID).
			Exec(ctx)
		if err != nil {
			return nil, errors.New("failed to update cart item")
		}

		return cartItem, nil
	}

	// ✅ ถ้าไม่มีสินค้า → เพิ่มสินค้าใหม่
	cartItem = &model.CartItem{
		CartID:             cart.ID,
		ProductID:          req.ProductID,
		TotalProductAmount: req.TotalProductAmount,
		Status:             "in_cart",
	}
	cartItem.SetCreatedNow()
	cartItem.SetUpdateNow()

	_, err = db.NewInsert().Model(cartItem).Exec(ctx)
	if err != nil {
		return nil, errors.New("failed to insert cart item")
	}

	return cartItem, nil
}



func UpdateCartItemService(ctx context.Context, UserID int, cartItemID int, req requests.CartItemUpdateRequest) (*model.CartItem, error) {
	// ตรวจสอบว่าผู้ใช้มีตะกร้าหรือไม่
	var cart model.Carts
	cartExists, err := db.NewSelect().
		Model(&cart).
		Where("user_id = ?", UserID).
		Exists(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to check cart: %v", err)
	}

	if !cartExists {
		return nil, errors.New("cart not found for this user")
	}

	// ดึง cart_id ที่ตรงกับ user
	err = db.NewSelect().
		Model(&cart).
		Where("user_id = ?", UserID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart ID: %v", err)
	}

	// ตรวจสอบว่าสินค้าอยู่ใน cart หรือไม่
	var cartItem model.CartItem
	cartItemExists, err := db.NewSelect().
		Model(&cartItem).
		Where("id = ?", cartItemID).
		Where("cart_id = ?", cart.ID).
		Exists(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to check cart item: %v", err)
	}

	if !cartItemExists {
		return nil, errors.New("cart item not found in this cart")
	}

	// ดึงข้อมูล cart_item
	err = db.NewSelect().
		Model(&cartItem).
		Where("id = ?", cartItemID).
		Where("cart_id = ?", cart.ID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart item: %v", err)
	}

	// ตรวจสอบว่า total_product_amount ต้องมากกว่า 0
	if req.TotalProductAmount <= 0 {
		return nil, errors.New("total product amount must be greater than 0")
	}

	cartItem.TotalProductAmount = req.TotalProductAmount
	cartItem.SetUpdateNow()

	_, err = db.NewUpdate().
		Model(&cartItem).
		Column("total_product_amount", "updated_at"). // อัปเดตเฉพาะฟิลด์ที่กำหนด
		Where("id = ?", cartItem.ID).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update cart item: %v", err)
	}

	return &cartItem, nil
}

func DeleteCartItemService(ctx context.Context, userID int, cartItemID int) error {
	// ดึง cart_id ของ user จากฐานข้อมูล
	var cart model.Carts
	err := db.NewSelect().
		Model(&cart).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return errors.New("cart not found for this user")
	}

	// ตรวจสอบว่า cart_item_id มีอยู่ใน cart_id นี้หรือไม่
	itemExists, err := db.NewSelect().
		TableExpr("cart_items").
		Where("id = ?", cartItemID).
		Where("cart_id = ?", cart.ID).
		Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check cart item: %v", err)
	}
	if !itemExists {
		return errors.New("cart item not found in this cart")
	}

	// ลบ cart_item_id ที่ระบุ
	_, err = db.NewDelete().
		TableExpr("cart_items").
		Where("id = ?", cartItemID).
		Where("cart_id = ?", cart.ID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %v", err)
	}

	query := db.NewSelect().
		TableExpr("cart_items").
		Where("cart_id = ?", cart.ID)

	total, err := query.Count(ctx)
	if err != nil {
		return fmt.Errorf("failed to count remaining cart items: %v", err)
	}

	if total == 0 {
		_, err = db.NewDelete().
			Model(&cart).
			Where("id = ?", cart.ID).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to delete empty cart: %v", err)
		}
	}

	return nil
}
