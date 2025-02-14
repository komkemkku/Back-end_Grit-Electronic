package products

import (
	"context"
	"errors"
	"fmt"
	"time"

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

	query := db.NewSelect().
		TableExpr("products AS p").
		Column("p.id", "p.name", "p.price", "p.description", "p.stock", "p.image", "p.is_active", "p.created_at", "p.updated_at", "p.deleted_at").
		ColumnExpr("c.id AS category__id").
		ColumnExpr("c.name AS category__name").
		ColumnExpr("COALESCE(json_agg(json_build_object('id', r.id, 'description', r.description, 'rating', r.rating, 'username', u.username)) FILTER (WHERE r.id IS NOT NULL), '[]') AS reviews").
		// ColumnExpr("COALESCE(json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description), '{}'::json) AS image").
		Join("LEFT JOIN categories AS c ON c.id = p.category_id").
		Join("LEFT JOIN reviews AS r ON r.product_id = p.id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id").
		// Join("LEFT JOIN images AS i ON i.ref_id = p.id AND i.type = 'product_main'").
		Where("p.deleted_at IS NULL").
		GroupExpr("p.id, c.id")

	// **เพิ่มเงื่อนไขกรองประเภทสินค้า**
	if req.CategoryID > 0 {
		query.Where("c.id = ?", req.CategoryID)
	}

	// **เพิ่มเงื่อนไขค้นหาด้วยชื่อสินค้า**
	if req.Search != "" {
		query.Where("p.name ILIKE ?", "%"+req.Search+"%")
	}

	query.Order("p.id ASC")

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

func GetByIdProductService(ctx context.Context, id int64) (*response.ProductDetailResponses, error) {
	// ตรวจสอบว่ามีสินค้าหรือไม่
	ex, err := db.NewSelect().
		TableExpr("products").
		Where("id = ?", id).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("product not found")
	}

	product := &response.ProductDetailResponses{}

	// ดึงข้อมูลสินค้า พร้อมรีวิวและรูปภาพ
	err = db.NewSelect().TableExpr("products AS p").
		Column("p.id", "p.name", "p.price", "p.description", "p.stock", "p.image", "p.is_active", "p.created_at", "p.updated_at", "p.deleted_at").
		ColumnExpr("c.id AS category__id").
		ColumnExpr("c.name AS category__name").
		ColumnExpr("COALESCE(json_agg(json_build_object('id', r.id, 'description', r.description, 'rating', r.rating, 'username', u.username)) FILTER (WHERE r.id IS NOT NULL), '[]') AS reviews").
		// ColumnExpr("COALESCE(json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description), '{}'::json) AS image").
		Join("LEFT JOIN categories AS c ON c.id = p.category_id").
		Join("LEFT JOIN reviews AS r ON r.product_id = p.id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id").
		// Join("LEFT JOIN images AS i ON i.ref_id = p.id AND i.type = 'product_main'").
		Where("p.deleted_at IS NULL").
		GroupExpr("p.id, c.id").
		Where("p.id = ?", id).
		Scan(ctx, product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// AddProduct ฟังก์ชันสำหรับเพิ่มสินค้าใหม่
func CreateProductService(ctx context.Context, req requests.ProductCreateRequest) (*model.Products, error) {

	// ตรวจสอบว่า category_id มีอยู่ในระบบหรือไม่
	ex, err := db.NewSelect().TableExpr("categories").Where("id = ?", req.CategoryID).Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	if !ex {
		return nil, fmt.Errorf("categories not found for ID: %d", req.CategoryID)
	}


	// ตรวจสอบว่าสินค้าชื่อนี้มีอยู่แล้วหรือไม่
	productExists, err := db.NewSelect().
		TableExpr("products").
		Where("name = ?", req.Name).
		Where("category_id = ?", req.CategoryID).
		Where("deleted_at IS NULL").
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if productExists {
		return nil, errors.New("product already exists in this category")
	}

	// เพิ่มสินค้าใหม่
	product := &model.Products{
		Name:        req.Name,
		Price:       float64(req.Price),
		Description: req.Description,
		Stock:       int(req.Stock),
		Image:       req.Image,
		CategoryID:  int(req.CategoryID),
		IsActive:    req.IsActive,
	}
	product.SetCreatedNow()
	product.SetUpdateNow()

	_, err = db.NewInsert().Model(product).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func UpdateProductService(ctx context.Context, id int, req requests.ProductUpdateRequest) (*model.Products, error) {
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

	// อัปเดตข้อมูลสินค้า
	product.Name = req.Name
	product.Price = float64(req.Price)
	product.Stock = int(req.Stock)
	product.Description = req.Description
	product.Image = req.Image
	product.IsActive = req.IsActive
	product.CategoryID = int(req.CategoryID)
	product.SetUpdateNow()

	// อัปเดตข้อมูลสินค้าในฐานข้อมูล
	_, err = db.NewUpdate().Model(product).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// // อัปเดตรูปภาพสินค้า
	// if req.ImageProduct != "" {
	// 	image := &model.Images{}
	// 	err = db.NewSelect().
	// 		Model(image).
	// 		Where("ref_id = ? AND type = 'product_main'", id).
	// 		Scan(ctx)

	// 	if err != nil && errors.Is(err, sql.ErrNoRows) {
	// 		// หากไม่มีรูปภาพ ให้เพิ่มใหม่
	// 		image = &model.Images{
	// 			RefID:       id,
	// 			Type:        "product_main",
	// 			Description: req.ImageProduct,
	// 		}
	// 		image.SetCreatedNow()
	// 		image.SetUpdateNow()

	// 		_, err = db.NewInsert().Model(image).Exec(ctx)
	// 		if err != nil {
	// 			return nil, errors.New("failed to create image")
	// 		}
	// 	} else if err == nil {
	// 		// หากมีรูปภาพอยู่แล้ว อัปเดตข้อมูล
	// 		image.Description = req.ImageProduct
	// 		image.SetUpdateNow()

	// 		_, err = db.NewUpdate().Model(image).Where("id = ?", image.ID).Exec(ctx)
	// 		if err != nil {
	// 			return nil, errors.New("failed to Update image")
	// 		}
	// 	} else {
	// 		return nil, errors.New("failed to check product image")
	// 	}
	// }

	return product, nil
}

func DeleteProductService(ctx context.Context, productID int64) error {
	// ตรวจสอบว่าสินค้ามีอยู่ในฐานข้อมูลหรือไม่
	product := &model.Products{}
	err := db.NewSelect().Model(product).Where("id = ?", productID).Scan(ctx)
	if err != nil {
		return errors.New("product not found") // หากไม่พบสินค้า
	}

	// ตั้งค่า is_active เป็น false และบันทึกเวลาใน deleted_at
	product.IsActive = false
	timestamp := time.Now().Unix()

	_, err = db.NewUpdate().
		Model(product).
		Set("is_active = ?", false).
		Set("deleted_at = ?", timestamp).
		Where("id = ?", productID).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to update product as deleted")
	}

	// ลบรูปภาพของสินค้า
	_, err = db.NewDelete().
		TableExpr("images").
		Where("ref_id = ? AND type = 'product_main'", productID).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to delete product images")
	}

	return nil
}
