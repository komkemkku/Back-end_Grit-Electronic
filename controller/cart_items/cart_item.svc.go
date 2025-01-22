package cartitems

import (
	"context"
	"errors"

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
	// ตรวจสอบว่ามี cart ที่เกี่ยวข้องหรือไม่
	cartExists, err := db.NewSelect().TableExpr("carts").Where("id = ?", req.CartID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !cartExists {
		return nil, errors.New("cart not found")
	}

	// ตรวจสอบว่าสินค้ามีอยู่ใน cart_items หรือไม่
	cartItem := &model.CartItem{}
	exists, err := db.NewSelect().Model(cartItem).
		Where("cart_id = ?", req.CartID).
		Where("product_id = ?", req.ProductID).
		Exists(ctx)
	if err != nil {
		return nil, err
	}

	if exists {
		// ถ้ามีสินค้าอยู่แล้ว เพิ่มจำนวนสินค้า
		cartItem.TotalProductAamount += req.TotalProductAmount
		cartItem.TotalProductPrice += req.TotalProductPrice
		cartItem.SetUpdateNow()

		_, err = db.NewUpdate().Model(cartItem).Where("id = ?", cartItem.ID).Exec(ctx)
		if err != nil {
			return nil, err
		}
		return cartItem, nil
	}

	// เพิ่มสินค้าใหม่ลงในตะกร้า
	cartItem = &model.CartItem{
		CartID:              req.CartID,
		ProductID:           req.ProductID,
		ProductName:         req.ProductName,
		ProductImageMain:    req.ProductImageMain,
		TotalProductPrice:   req.TotalProductPrice,
		TotalProductAamount: req.TotalProductAmount,
		Status:              req.Status,
	}
	cartItem.SetCreatedNow()
	cartItem.SetUpdateNow()

	_, err = db.NewInsert().Model(cartItem).Exec(ctx)
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


func UpdateCartItemService(ctx context.Context, id int64, req requests.CartItemUpdateRequest) (*model.CartItem, error) {
	cartItem := &model.CartItem{}
	err := db.NewSelect().Model(cartItem).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, errors.New("cart_item not found")
	}

	// อัปเดตข้อมูลสินค้าในตะกร้า
	cartItem.TotalProductAamount = req.TotalProductAmount
	cartItem.TotalProductPrice = req.TotalProductPrice
	cartItem.ProductName = req.ProductName
	cartItem.ProductImageMain = req.ProductImageMain
	cartItem.SetUpdateNow()

	_, err = db.NewUpdate().Model(cartItem).Where("id = ?", id).Exec(ctx)
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

func DeleteCartItemService(ctx context.Context, id int64) error {
	cartItem := &model.CartItem{}
	err := db.NewSelect().Model(cartItem).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return errors.New("cart_item not found")
	}

	// ลบสินค้าใน cart_items
	_, err = db.NewDelete().Model(cartItem).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	// ตรวจสอบว่าตะกร้ามีสินค้าเหลือหรือไม่
	count, err := db.NewSelect().TableExpr("cart_items").Where("cart_id = ?", cartItem.CartID).Count(ctx)
	if err != nil {
		return err
	}

	if count == 0 {
		// ลบ cart หากไม่มีสินค้าเหลือ
		_, err = db.NewDelete().TableExpr("carts").Where("id = ?", cartItem.CartID).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

