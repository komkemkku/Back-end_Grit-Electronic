package products

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

func ListProductService(ctx context.Context, req requests.ProductRequest) ([]response.ProductResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.ProductResponses{}

	// สร้าง query
	// query := db.NewSelect().
	// TableExpr("products AS p").
	// Column("p.id", "p.name", "p.price", "p.description", "p.created_at", "p.updated_at").
	// ColumnExpr("c.id AS category__id").
	// ColumnExpr("c.name AS category__name").
	// ColumnExpr("json_agg(json_build_object('id', r.id, 'text_review', r.text_review, 'rating', r.rating, 'username', u.username)) AS reviews").
	// Join("LEFT JOIN categories AS c ON c.id = p.category_id").
	// Join("LEFT JOIN reviews AS r ON r.product_id = p.id").
	// Join("LEFT JOIN users as u ON u.id = r.user_id").
	// GroupExpr("p.id, c.id")

	query := db.NewSelect().
		TableExpr("products AS p").
		Column("p.id", "p.name", "p.price", "p.description", "p.stock", "p.is_active", "p.created_at", "p.updated_at").
		ColumnExpr("c.id AS category__id").
		ColumnExpr("c.name AS category__name").
		ColumnExpr("i.description AS image").
		ColumnExpr("c.is_active AS is_active").
		Join("LEFT JOIN categories AS c ON c.id = p.category_id").
		Join("LEFT JOIN images AS i ON i.ref_id = p.id AND i.type = 'product_main'")

	// query.Where("p.is_active = ?", true)
	query.Order("p.id ASC")

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
		Column("p.id", "p.name", "p.price", "p.description", "p.stock", "p.is_active", "p.created_at", "p.updated_at").
		ColumnExpr("c.id AS category__id").
		ColumnExpr("c.name AS category__name").
		ColumnExpr("r.id AS review__id").
		ColumnExpr("r.description AS review__description").
		ColumnExpr("r.rating AS review__rating").
		// ColumnExpr("u.id AS user__id").
		// ColumnExpr("u.username AS user__username").
		ColumnExpr("i.description AS image").
		ColumnExpr("c.is_active AS is_active").
		Join("LEFT JOIN categories AS c ON c.id = p.category_id").
		Join("LEFT JOIN reviews AS r ON p.id = r.product_id").
		Join("LEFT JOIN users as u ON u.id = r.user_id").
		Join("LEFT JOIN images AS i ON i.ref_id = p.id AND i.type = 'product_main'").
		Where("p.id = ?", id).Scan(ctx, product)
		
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
		Name:        req.Name,
		Price:       float64(req.Price),
		Description: req.Description,
		Stock:       int(req.Stock),
		CategoryID:  int(req.CategoryID),
		IsActive:    req.IsActive,
	}
	product.SetCreatedNow()
	product.SetUpdateNow()

	_, err = db.NewInsert().Model(product).Exec(ctx)
	if err != nil {
		return nil, err
	}

	img := requests.ImageCreateRequest{
		RefID:       product.ID,
		Type:        "product_main",
		Description: req.ImageProduct,
	}

	_, err = image.CreateImageService(ctx, img)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func UpdateProductService(ctx context.Context, id int64, req requests.ProductUpdateRequest) (*model.Products, error) {
	// ตรวจสอบว่าสินค้ามีอยู่ในฐานข้อมูลหรือไม่
	exists, err := db.NewSelect().TableExpr("products").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("product not found")
	}

	product := &model.Products{}

	err = db.NewSelect().Model(product).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Price = float64(req.Price)
	product.Stock += int(req.Stock)
	product.Description = req.Description
	product.IsActive = req.IsActive
	product.CategoryID = int(req.CategoryID)
	product.SetUpdateNow()

	// อัปเดตข้อมูลสินค้าในฐานข้อมูล
	_, err = db.NewUpdate().Model(product).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	img := requests.ImageCreateRequest{
		RefID:       product.ID,
		Type:        "product_main",
		Description: req.ImageProduct,
	}

	_, err = image.CreateImageService(ctx, img)
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
