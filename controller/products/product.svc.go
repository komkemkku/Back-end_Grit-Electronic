package products

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListProductService(ctx context.Context, req requests.ProductRequest) ([]response.ProductResponses, int, error) {

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

func GetByIdProductService(ctx context.Context, id int64) (*response.ProductResponses, error) {
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
		ColumnExpr("c.id AS category__id").
		ColumnExpr("c.name AS category__name").
		Join("LEFT JOIN categories as c ON c.id = p.category_id").Where("p.id = ?", id).Scan(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// AddProduct ฟังก์ชันสำหรับเพิ่มสินค้าใหม่
func CreateProductService(ctx context.Context, req requests.ProductCreateRequest) (*model.Products, error) {

	// ตรวจสอบว่า category_id มีอยู่ในระบบหรือไม่
	ex, err := db.NewSelect().TableExpr("categories").Where("id =?", req.CategoryID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("categories not found")
	}

	// เพิ่มสินค้าใหม่
	product := &model.Products{
		Name:       req.Name,
		Price:      float64(req.Price),
		Detail:     req.Detail,
		Stock:      int(req.Stock),
		Image:      req.Image,
		Spec:       req.Spec,
		CategoryID: int(req.CategoryID),
	}
	product.SetCreatedNow()
	product.SetUpdateNow()

	_, err = db.NewInsert().Model(product).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func UpdateProductService(ctx context.Context, id int64, req requests.ProductUpdateRequest) (*model.Products, error) {
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

func DeleteProductService(ctx context.Context, id int64) error {
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
