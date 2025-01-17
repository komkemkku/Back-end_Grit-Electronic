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

	// สร้าง query
	query := db.NewSelect().
		TableExpr("categories AS c").
		Column("c.id", "c.name", "c.is_active", "c.created_at", "c.updated_at")

	// query.Where("c.is_active = ?", true)

	if req.Search != "" {
		query.Where("c.name ILIKE ?", "%"+req.Search+"%")
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

func GetByIdCategoryService(ctx context.Context, id int64) (*response.CategoryResponses, error) {
	ex, err := db.NewSelect().TableExpr("categories").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("category not found")
	}
	category := &response.CategoryResponses{}

	err = db.NewSelect().TableExpr("categories AS c").
		Column("c.id", "c.name", "c.is_active", "c.created_at", "c.updated_at").
		Where("c.id = ?", id).Scan(ctx, category)
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
		IsActive: req.IsActive,
	}
	category.SetCreatedNow()
	category.SetUpdateNow()

	_, err = db.NewInsert().Model(category).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return category, nil

}

func UpdateCategoryService(ctx context.Context, id int64, req requests.CategoryUpdateRequest) (*model.Categories, error) {
	ex, err := db.NewSelect().TableExpr("categories").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("category not found")
	}

	category := &model.Categories{}

	err = db.NewSelect().Model(category).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	category.Name = req.Name
	category.IsActive = req.IsActive
	category.SetUpdateNow()

	_, err = db.NewUpdate().Model(category).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func DeleteCetegoryService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("categories").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("category not found")
	}

	_, err = db.NewDelete().TableExpr("categories").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
