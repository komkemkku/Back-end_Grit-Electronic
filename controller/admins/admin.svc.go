package admins

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils"
)

var db = configs.Database()

func ListAdminService(ctx context.Context, req requests.AdminRequest) ([]response.AdminResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.AdminResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("admins AS a").
		Column("a.id", "a.name", "a.email", "a.is_active", "a.created_at", "a.updated_at").
		ColumnExpr("r.id AS role__id").
		ColumnExpr("r.name AS role__name").
		Join("LEFT JOIN roles as r ON r.id = a.role_id")

	if req.Search != "" {
		query.Where("a.name ILIKE ?", "%"+req.Search+"%")
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

func GetByIdAdminService(ctx context.Context, id int64) (*response.AdminResponses, error) {
	ex, err := db.NewSelect().TableExpr("admins").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("admin not found")
	}
	admin := &response.AdminResponses{}

	err = db.NewSelect().TableExpr("admins AS a").
	Column("a.id", "a.name", "a.email", "a.is_active", "a.created_at", "a.updated_at").
	ColumnExpr("r.id AS role__id").
	ColumnExpr("r.name AS role__name").
	Join("LEFT JOIN roles as r ON r.id = a.role_id").Where("a.id = ?", id).Scan(ctx, admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func CreateAdminService(ctx context.Context, req requests.AdminCreateRequest) (*model.Admins, error) {

	ex, err := db.NewSelect().TableExpr("roles").Where("id =?", req.RoleID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("role not found")
	}

	hashpassword, _ := utils.HashPassword(req.Password)

	// เพิ่มadmin
	admin := &model.Admins{
		RoleID:   int(req.RoleID),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashpassword,
		IsActive: req.IsActive,
	}
	admin.SetCreatedNow()
	admin.SetUpdateNow()

	_, err = db.NewInsert().Model(admin).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return admin, nil

}

func UpdateAdminService(ctx context.Context, id int64, req requests.AdminUpdateRequest) (*model.Admins, error) {
	ex, err := db.NewSelect().TableExpr("admins").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("admin not found")
	}

	admin := &model.Admins{}

	hashpassword, _ := utils.HashPassword(req.Password)

	err = db.NewSelect().Model(admin).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	admin.RoleID = req.RoleID
	admin.Name = req.Name
	admin.Email = req.Email
	admin.Password = hashpassword
	admin.IsActive = req.IsActive
	admin.SetUpdateNow()

	_, err = db.NewUpdate().Model(admin).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func DeleteAdminService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("admins").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("admin not found")
	}

	_, err = db.NewDelete().TableExpr("admins").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
