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

func ListWishlistsService(ctx context.Context, req requests.ProductRequest) ([]response.ProductResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.ProductResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("products AS p").
		Column("p.id", "p.name", "p.price", "p.detail", "p.stock", "p.image", "p.spec", "p.created_at", "p.updated_at").
		ColumnExpr("c.id AS category__id").
		ColumnExpr("c.name AS category__name").
		Join("LEFT JOIN categories as c ON c.id = p.category_id")

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

func GetByIdWishlistsService(ctx context.Context, id int64) (*response.ProductResponses, error) {
	ex, err := db.NewSelect().TableExpr("products").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("product not found")
	}
	product := &response.ProductResponses{}

	err = db.NewSelect().TableExpr("products AS p").
		Column("p.id", "p.name", "p.price", "p.detail", "p.stock", "p.image", "p.spec", "p.created_at", "p.updated_at").
		ColumnExpr("c.id AS Product__id").
		ColumnExpr("c.name AS Product__name").
		Join("LEFT JOIN Product as c ON c.id = w.Product_id").Where("p.id = ?", id).Scan(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
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

func UpdateWishlistsService(ctx context.Context, id int64, req requests.ProductUpdateRequest) (*model.Products, error) {
	ex, err := db.NewSelect().TableExpr("products").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("product not found")
	}

	product := &model.Products{}

	err = db.NewSelect().Model(product).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	product.Name = req.Name
	product.Price = float64(req.Price)
	product.Detail = req.Detail
	product.Stock = int(req.Stock)
	product.Image = req.Image
	product.Spec = req.Spec
	product.CategoryID = int(req.CategoryID)
	product.SetUpdateNow()

	_, err = db.NewUpdate().Model(product).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteWishlistsService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("products").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("product not found")
	}

	_, err = db.NewDelete().TableExpr("products").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
