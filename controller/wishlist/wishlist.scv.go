package wishlist

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

func ListWishlistsService(ctx context.Context, req requests.WishlistsRequest) ([]response.WishlistResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.WishlistResponses{}
	// สร้าง query
	query := db.NewSelect().
		TableExpr("wishlists AS w").
		Column("w.id", "w.created_at", "w.updated_at").
		ColumnExpr("u.id AS username__id").
		ColumnExpr("u.username AS username__id").
		// ColumnExpr("p.id AS product__id").
		// ColumnExpr("p.name AS product__name").
		// ColumnExpr("p.price AS product__price").
		Join("LEFT JOIN products AS p ON p.id = w.product_id").
		Join("LEFT JOIN users AS u ON u.id = w.user_id")

	if req.Search != "" {
		query.Where("p.name ILIKE ?", "%"+req.Search+"%")
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


func GetByIdWishlistsService(ctx context.Context, id int64) (*response.WishlistResponses, error) {

	ex, err := db.NewSelect().TableExpr("wishlists").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("wish not found")
	}
	wish := &response.WishlistResponses{}

	err = db.NewSelect().TableExpr("wishlists AS w").
		Column("w.id", "w.created_at", "w.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("p.detail AS product__detail").
		ColumnExpr("p.price AS product__price").
		Join("LEFT JOIN products AS p ON p.id = w.product_id").Where("w.id = ?", id).Scan(ctx, wish)
	if err != nil {
		return nil, err
	}
	return wish, nil
}

func CreateWishlistsService(ctx context.Context, req requests.WishlistsAddRequest) error {
	wishlist := &model.Wishlists{
		UserID: req.UserID,
		ProductID: req.ProductID,
        PricePerProduct: req.PricePerProduct,
        AmountPerProduct: req.AmountPerProduct,
	}
	wishlist.SetCreatedNow()
	wishlist.SetUpdateNow()

	// บันทึกข้อมูลลงฐานข้อมูล
	if _, err := db.NewInsert().Model(wishlist).Exec(ctx); err != nil {
		return errors.New("failed to create wishlist: " + err.Error())
	}

	return nil
}

func DeleteWishlistsService(ctx context.Context, id int64) error {
	// ตรวจสอบว่า Wishlist มีอยู่หรือไม่
	ex, err := db.NewSelect().TableExpr("wishlists").Where("id = ?", id).Exists(ctx)

	if err != nil {
		// กรณีเกิดข้อผิดพลาดจากฐานข้อมูล
		return err
	}

	if !ex {
		// กรณี Wishlist ไม่พบในฐานข้อมูล
		return errors.New("Wishlist not found")
	}

	// ลบ Wishlist ที่พบในฐานข้อมูล
	_, err = db.NewDelete().TableExpr("wishlists").Where("id = ?", id).Exec(ctx)
	if err != nil {
		// กรณีลบไม่สำเร็จ
		return err
	}

	// สำเร็จ
	return nil
}


func UpdateWishlistsService(ctx context.Context, id int64, req requests.WishlistsUpdateRequest) (*model.Wishlists, error) {
	// ตรวจสอบว่า Wishlist มีอยู่ในระบบหรือไม่
	exists, err := db.NewSelect().
		TableExpr("wishlists").
		Where("id = ?", id).
		Exists(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to check wishlist existence: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("wishlist with id %d not found", id)
	}

	// ตรวจสอบว่า product_id มีอยู่ในระบบหรือไม่
	productExists, err := db.NewSelect().
		TableExpr("products").
		Where("id = ?", req.ProductID).
		Exists(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to check product existence: %v", err)
	}

	if !productExists {
		return nil, fmt.Errorf("product with id %d not found", req.ProductID)
	}

	// ดึงข้อมูล Wishlist
	wishlist := &model.Wishlists{}
	err = db.NewSelect().
		Model(wishlist).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve wishlist: %v", err)
	}

	// อัปเดต Wishlist
	wishlist.SetUpdateNow()

	_, err = db.NewUpdate().
		Model(wishlist).
		Where("id = ?", id).
		Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to update wishlist: %v", err)
	}

	return wishlist, nil
}
