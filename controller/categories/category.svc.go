package categories

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListCategoryService(ctx context.Context, req requests.CategoryRequest) ([]response.CategoryResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.CategoryResponses{}

	// สร้าง query พร้อมกับนับจำนวนสินค้าภายในประเภทนั้น ๆ
	query := db.NewSelect().
		TableExpr("categories AS c").
		Column("c.id", "c.name", "c.image", "c.is_active", "c.created_at", "c.updated_at").
		ColumnExpr("COALESCE(COUNT(p.id), 0) AS total_products").
		Join("LEFT JOIN products AS p ON p.category_id = c.id AND p.is_active = TRUE AND p.deleted_at IS NULL").
		GroupExpr("c.id")

	if req.Search != "" {
		query.Where("c.name ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// ถ้าไม่พบสินค้า ให้คืนค่าว่าง
	if total == 0 {
		return []response.CategoryResponses{}, 0, nil
	}

	// Execute query
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdCategoryService(ctx context.Context, id int64) (*response.CategoryResponses, error) {
	// ตรวจสอบว่าหมวดหมู่มีอยู่หรือไม่
	exists, err := db.NewSelect().
		Table("categories").
		Where("id = ?", id).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("category not found")
	}

	category := &response.CategoryResponses{}

	// ดึงข้อมูลหมวดหมู่พร้อมจำนวนสินค้า
	err = db.NewSelect().
		TableExpr("categories AS c").
		Column("c.id", "c.name", "c.image", "c.is_active", "c.created_at", "c.updated_at").
		ColumnExpr("COALESCE(COUNT(p.id), 0) AS total_products").
		Join("LEFT JOIN products AS p ON p.category_id = c.id AND p.is_active = TRUE AND p.deleted_at IS NULL").
		Where("c.id = ?", id).
		GroupExpr("c.id").
		Scan(ctx, category)

	if err != nil {
		return nil, err
	}
	return category, nil
}

func CreateCategoryService(ctx context.Context, req requests.CategoryCreateRequest) (*model.Categories, error) {

	// ตรวจสอบชื่อซ้ำ
	exists, err := db.NewSelect().
		TableExpr("categories").
		Where("name = ?", req.Name).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("category already exists")
	}

	// เพิ่ม
	category := &model.Categories{
		Name:     req.Name,
		Image:    req.Image,
		IsActive: true,
	}
	category.SetCreatedNow()
	category.SetUpdateNow()

	_, err = db.NewInsert().Model(category).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// img := requests.ImageCreateRequest{
	// 	RefID:       category.ID,
	// 	Type:        "categories",
	// 	Description: req.ImageCategories,
	// }

	// _, err = image.CreateImageService(ctx, img)
	// if err != nil {
	// 	return nil, err
	// }

	return category, nil

}

func UpdateCategoryService(ctx context.Context, id int64, req requests.CategoryUpdateRequest) (*model.Categories, error) {
	// ตรวจสอบว่าหมวดหมู่มีอยู่หรือไม่
	exists, err := db.NewSelect().TableExpr("categories").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("category not found")
	}

	// ดึงข้อมูลหมวดหมู่เดิม
	category := &model.Categories{}
	err = db.NewSelect().Model(category).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// อัปเดตข้อมูลหมวดหมู่
	category.Name = req.Name
	category.Image = req.Image
	category.IsActive = true
	category.SetUpdateNow()

	_, err = db.NewUpdate().Model(category).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// // ตรวจสอบว่ามีรูปภาพเดิมอยู่หรือไม่
	// imgExists, err := db.NewSelect().
	// 	TableExpr("images").
	// 	Where("ref_id = ? AND type = 'categories'", category.ID).
	// 	Exists(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// if imgExists {
	// 	// อัปเดตรูปภาพเดิม
	// 	_, err = db.NewUpdate().
	// 		TableExpr("images").
	// 		Set("description = ?", req.ImageCategories).
	// 		Where("ref_id = ? AND type = 'categories'", category.ID).
	// 		Exec(ctx)
	// 	if err != nil {
	// 		return nil, errors.New("failed to update category image")
	// 	}
	// } else {
	// 	// ถ้ายังไม่มีรูปภาพ ให้เพิ่มใหม่
	// 	img := requests.ImageCreateRequest{
	// 		RefID:       category.ID,
	// 		Type:        "categories",
	// 		Description: req.ImageCategories,
	// 	}

	// 	_, err = image.CreateImageService(ctx, img)
	// 	if err != nil {
	// 		return nil, errors.New("failed to create category image")
	// 	}
	// }

	return category, nil
}

func DeleteCategoryService(ctx context.Context, categoryID int64) error {
	exists, err := db.NewSelect().
		TableExpr("categories").
		Where("id = ?", categoryID).
		Exists(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("category not found")
	}

	// // ลบรูปภาพที่เกี่ยวข้องกับหมวดหมู่
	// _, err = db.NewDelete().
	// 	TableExpr("images").
	// 	Where("ref_id = ? AND type = 'categories'", categoryID).
	// 	Exec(ctx)
	// if err != nil {
	// 	return errors.New("failed to delete category images")
	// }

	// ลบหมวดหมู่จากฐานข้อมูล
	_, err = db.NewDelete().
		TableExpr("categories").
		Where("id = ?", categoryID).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to delete category")
	}

	return nil
}
