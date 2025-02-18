package wishlist

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListWishlistsService(ctx context.Context, req requests.WishlistsRequest) ([]response.WishListByUser, int, error) {
	var resp []response.WishListByUser

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	query := db.NewSelect().
		TableExpr("wishlists w").
		ColumnExpr("p.id AS product__id, p.name AS product__name, p.price AS product__price, p.image AS product__image").
		Join("LEFT JOIN users u ON u.id = w.user_id").
		Join("LEFT JOIN products p ON p.id = w.product_id").
		Where("w.user_id = ?", req.UserID) // แสดงเฉพาะ Wishlist ของผู้ใช้ที่ร้องขอ

	if req.Search != "" {
		query.Where("p.name ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

// func GetByIdWishlistsService(ctx context.Context, id int, userID int) ([]response.WishListByUser, error) {
// 	// ตรวจสอบว่า Wishlist มีอยู่หรือไม่ และเป็นของผู้ใช้ที่ร้องขอหรือไม่
// 	exists, err := db.NewSelect().
// 		TableExpr("wishlists AS w").
// 		Where("w.id = ? AND w.user_id = ?", id, userID).
// 		Exists(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !exists {
// 		return nil, errors.New("wishlist not found or unauthorized")
// 	}

// 	// ใช้ Slice ปกติ ไม่ต้องใช้ Pointer
// 	wish := []response.WishListByUser{}

// 	// ดึงข้อมูล Wishlist พร้อม Join ตารางที่เกี่ยวข้อง
// 	err = db.NewSelect().
// 		TableExpr("wishlists AS w").
// 		Column("w.id", "w.created_at", "w.updated_at").
// 		ColumnExpr("u.id AS user__id, u.username AS user__username").
// 		ColumnExpr("p.id AS product__id, p.name AS product__name").
// 		Join("LEFT JOIN users AS u ON u.id = w.user_id").
// 		Join("LEFT JOIN products AS p ON p.id = w.product_id").
// 		Where("w.id = ? AND w.user_id = ?", id, userID).
// 		Scan(ctx, &wish)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return wish, nil
// }

func CreateWishlistsService(ctx context.Context, req requests.WishlistsAddRequest) error {
	// ตรวจสอบว่ามีสินค้านั้นในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().
		Table("products").
		Where("id = ?", req.ProductID).
		Exists(ctx)

	if err != nil {
		return errors.New("product not found")
	}

	if !exists {
		return errors.New("product not found")
	}

	// ตรวจสอบว่าใน wishlist ของผู้ใช้งานมีสินค้านี้อยู่หรือไม่
	wishlistExists, err := db.NewSelect().
		Table("wishlists").
		Where("user_id = ? AND product_id = ?", req.UserID, req.ProductID).
		Exists(ctx)

	if err != nil {
		return errors.New("failed to check if product already in wishlist: " + err.Error())
	}

	// ถ้ามีสินค้าใน wishlist แล้ว
	if wishlistExists {
		return errors.New("this product is already in the wishlist")
	}

	wishlist := &model.Wishlists{
		UserID:    req.UserID,
		ProductID: req.ProductID,
	}
	wishlist.SetCreatedNow()
	wishlist.SetUpdateNow()

	// บันทึกข้อมูลลงฐานข้อมูล
	if _, err = db.NewInsert().Model(wishlist).Exec(ctx); err != nil {
		return errors.New("product not found")
	}

	return nil
}

func DeleteWishlistsService(ctx context.Context, id int64) error {
	// ตรวจสอบว่า Wishlist มีอยู่หรือไม่
	ex, err := db.NewSelect().TableExpr("wishlists").Where("id = ?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("Wishlist not found")
	}

	// ลบ Wishlist ที่พบในฐานข้อมูล
	_, err = db.NewDelete().TableExpr("wishlists").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func UpdateWishlistsService(ctx context.Context, userID int, productID int) (*model.Wishlists, string, bool, error) {
	var wishlistCount int
	err := db.NewSelect().
		Table("wishlists").
		Where("user_id = ? AND product_id = ?", userID, productID).
		ColumnExpr("COUNT(*)").
		Scan(ctx, &wishlistCount)

	if err != nil {
		return nil, "", false, errors.New("failed to check wishlist")
	}

	if wishlistCount > 0 {
		// ถ้ากดถูกใจอยู่แล้ว ให้ลบออกจาก Wishlist
		_, err = db.NewDelete().
			Table("wishlists").
			Where("user_id = ? AND product_id = ?", userID, productID).
			Exec(ctx)

		if err != nil {
			return nil, "", false, errors.New("failed to remove from wishlist")
		}

		return nil, "Product removed from wishlist", false, nil
	}

	// ถ้ายังไม่ได้กด ให้เพิ่มสินค้าเข้า Wishlist
	newWish := &model.Wishlists{
		UserID:    int(userID),
		ProductID: int(productID),
	}
	newWish.SetCreatedNow()

	_, err = db.NewInsert().Model(newWish).Returning("*").Exec(ctx)
	if err != nil {
		return nil, "", false, errors.New("failed to add to wishlist")
	}

	return newWish, "Product added to wishlist", true, nil
}

func GetWishlistStatusService(ctx context.Context, userID int64, productID int64) (bool, error) {
	var count int
	err := db.NewSelect().
		Table("wishlists").
		Where("user_id = ? AND product_id = ?", userID, productID).
		ColumnExpr("COUNT(*)").
		Scan(ctx, &count)

	if err != nil {
		return false, errors.New("failed to check wishlist status")
	}

	// ถ้าจำนวนมากกว่า 0 แสดงว่าผู้ใช้กดถูกใจสินค้าแล้ว
	return count > 0, nil
}

