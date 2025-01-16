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
		TableExpr("system_bank AS sb").
		Column("sb.id", "sb.bank_name", "sb.account_name", "sb.account_number", "sb.image", "sb.created_at", "sb.updated_at")

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
	ex, err := db.NewSelect().TableExpr("system_bank").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("not found")
	}
	systembank := &response.SystemBankResponses{}

	err = db.NewSelect().TableExpr("system_bank AS sb").
	Column("sb.id", "sb.bank_name", "sb.account_name", "sb.account_number", "sb.image", "sb.created_at", "sb.updated_at").
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
	}
	systembank.SetCreatedNow()
	systembank.SetUpdateNow()

	_, err := db.NewInsert().Model(systembank).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return systembank, nil

}

func UpdateSystembankService(ctx context.Context, id int64, req requests.SystemBankUpdateRequest) (*model.SystemBanks, error) {
	ex, err := db.NewSelect().TableExpr("system_bank").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("not found")
	}

	systembank := &model.SystemBanks{}

	err = db.NewSelect().Model(systembank).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	systembank.BankName = req.BankName
	systembank.AccountName = req.AccountName
	systembank.AccountNumber = req.AccountNumber
	systembank.SetUpdateNow()

	_, err = db.NewUpdate().Model(systembank).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return systembank, nil
}

func DeleteSystemBankService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("system_bank").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("SysBank not found")
	}

	_, err = db.NewDelete().TableExpr("system_bank").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
