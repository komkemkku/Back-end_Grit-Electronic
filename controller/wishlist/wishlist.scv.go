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
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		Join("LEFT JOIN products AS p ON p.id = w.product_id")

	// if req.Search != "" {
	//   query.Where("p.name ILIKE ?", "%"+req.Search+"%")
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
	Column("w.id", "p.created_at", "p.updated_at").
	ColumnExpr("p.id AS product__id").
	ColumnExpr("p.name AS product__name").
	Join("LEFT JOIN products AS p ON p.id = w.product_id").Where("w.id = ?", id).Scan(ctx, wish)
	if err != nil {
		return nil, err
	}
	return wish, nil
}

func CreateWishlistsService(ctx context.Context, req requests.WishlistsAddRequest) (*model.Wishlists, error) {

	// ตรวจสอบว่า category_id มีอยู่ในระบบหรือไม่
	// เพิ่มสินค้าใหม่
	Wishlists := &model.Wishlists{
		ProductID: req.ProductID,
	}
	Wishlists.SetCreatedNow()
	Wishlists.SetUpdateNow()

	_, err := db.NewInsert().Model(Wishlists).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return Wishlists, nil

}
