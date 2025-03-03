package systembank

import (
	"context"
	"errors"

	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = config.Database()

func ListSystemBankService(ctx context.Context, req requests.SystemBankRequest) ([]response.SystemBankResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.SystemBankResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("system_banks AS sb").
		Column("sb.id", "sb.bank_name", "sb.account_name", "sb.account_number", "sb.description", "sb.image", "sb.is_active", "sb.created_at", "sb.updated_at")
		// ColumnExpr("json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description) AS image").
		// Join("LEFT JOIN images AS i ON i.ref_id = sb.id AND i.type = 'systembank_image'")

	if req.Search != "" {
		query.Where("sb.bank_name ILIKE ?", "%"+req.Search+"%")
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

func GetByIdSystemBankService(ctx context.Context, id int64) (*response.SystemBankResponses, error) {
	ex, err := db.NewSelect().TableExpr("system_banks").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("not found")
	}
	systembank := &response.SystemBankResponses{}

	err = db.NewSelect().TableExpr("system_banks AS sb").
		Column("sb.id", "sb.bank_name", "sb.account_name", "sb.account_number", "sb.description", "sb.image", "sb.is_active", "sb.created_at", "sb.updated_at").
		Where("sb.id = ?", id).Scan(ctx, systembank)
	if err != nil {
		return nil, err
	}
	return systembank, nil
}

// Add bank system
func CreateSystemBankService(ctx context.Context, req requests.SystemBankCreateRequest) (*model.SystemBanks, error) {

	// เพิ่มเลขบัญชีระบบ
	systembank := &model.SystemBanks{
		BankName:      req.BankName,
		AccountName:   req.AccountName,
		AccountNumber: req.AccountNumber,
		Description:   req.Description,
		Image:         req.Image,
		IsActive:      req.IsActive,
	}
	systembank.SetCreatedNow()
	systembank.SetUpdateNow()

	_, err := db.NewInsert().Model(systembank).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// img := requests.ImageCreateRequest{
	// 	RefID:       systembank.ID,
	// 	Type:        "systembank_image",
	// 	Description: req.ImageSystemBank,
	// }

	// _, err = image.CreateImageService(ctx, img)
	// if err != nil {
	// 	return nil, err
	// }

	return systembank, nil

}

func UpdateSystembankService(ctx context.Context, id int64, req requests.SystemBankUpdateRequest) (*model.SystemBanks, error) {
	// ตรวจสอบว่ามี systembank อยู่หรือไม่
	exists, err := db.NewSelect().TableExpr("system_banks").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("system bank not found")
	}

	// ดึงข้อมูล systembank ที่เกี่ยวข้อง
	systembank := &model.SystemBanks{}
	err = db.NewSelect().Model(systembank).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// อัปเดตข้อมูล systembank
	systembank.BankName = req.BankName
	systembank.AccountName = req.AccountName
	systembank.AccountNumber = req.AccountNumber
	systembank.Description = req.Description
	systembank.Image = req.Image
	systembank.IsActive = req.IsActive
	systembank.SetUpdateNow()

	_, err = db.NewUpdate().Model(systembank).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// // ลบรูปภาพเดิมที่เกี่ยวข้องกับ systembank
	// _, err = db.NewDelete().
	// 	TableExpr("images").
	// 	Where("ref_id = ? AND type = 'systembank_image'", id).
	// 	Exec(ctx)
	// if err != nil {
	// 	return nil, errors.New("failed to delete image")
	// }

	// // เพิ่มรูปภาพใหม่
	// img := requests.ImageCreateRequest{
	// 	RefID:       systembank.ID,
	// 	Type:        "systembank_image",
	// 	Description: req.ImageSystemBank,
	// }

	// _, err = image.CreateImageService(ctx, img)
	// if err != nil {
	// 	return nil, errors.New("failed to create image service")
	// }

	return systembank, nil
}

func DeleteSystemBankService(ctx context.Context, id int64) error {
	// ตรวจสอบว่ามี system_banks อยู่หรือไม่
	exists, err := db.NewSelect().TableExpr("system_banks").Where("id=?", id).Exists(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("SystemBank not found")
	}

	// // ลบรูปภาพที่เกี่ยวข้องกับ SystemBank
	// _, err = db.NewDelete().
	// 	TableExpr("images").
	// 	Where("ref_id = ? AND type = 'systembank_image'", id).
	// 	Exec(ctx)
	// if err != nil {
	// 	return errors.New("failed to delete image")
	// }

	// ลบ SystemBank
	_, err = db.NewDelete().TableExpr("system_banks").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return errors.New("failed to delete system bank")
	}

	return nil
}
