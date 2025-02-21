package payments

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListPaymentService(ctx context.Context, req requests.PaymentRequest) ([]response.PaymentResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.PaymentResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("payments AS p").
		Column("p.id", "p.price", "p.status", "p.updated_by", "p.bank_name", "p.account_name", "p.account_number", "p.created_at", "p.updated_at").
		ColumnExpr("sb.id AS systembank__id").
		ColumnExpr("sb.bank_name AS systembank__bank_name").
		ColumnExpr("sb.account_name AS systembank__account_name").
		ColumnExpr("sb.account_number AS systembank__account_number").
		ColumnExpr("sb.description AS systembank__description").
		ColumnExpr("json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description) AS image").
		ColumnExpr("i2.id AS imagesystembank__id").
		ColumnExpr("i2.ref_id AS imagesystembank__ref_id").
		ColumnExpr("i2.type AS imagesystembank__type").
		ColumnExpr("i2.description AS imagesystembank__description").
		Join("LEFT JOIN system_banks AS sb ON sb.id = p.system_bank_id").
		Join("LEFT JOIN Images AS i2 ON i2.ref_id = sb.id AND i2.type = 'systembank_image'").
		Join("LEFT JOIN images AS i ON i.ref_id = p.id AND i.type = 'payment_slip'").
		GroupExpr("p.id, sb.id,  i.id, i.ref_id, i.type, i.description, i2.id")

	if req.Search != "" {
		query.Where("p.status ILIKE ?", "%"+req.Search+"%")
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

func GetByIdPaymentService(ctx context.Context, id int64) (*response.PaymentResponses, error) {
	ex, err := db.NewSelect().TableExpr("payments").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("payment not found")
	}
	payment := &response.PaymentResponses{}

	err = db.NewSelect().TableExpr("payments AS p").
		Column("p.id", "p.price", "p.status", "p.updated_by", "p.bank_name", "p.account_name", "p.account_number", "p.created_at", "p.updated_at").
		ColumnExpr("sb.id AS systembank__id").
		ColumnExpr("sb.bank_name AS systembank__bank_name").
		ColumnExpr("sb.account_name AS systembank__account_name").
		ColumnExpr("sb.account_number AS systembank__account_number").
		ColumnExpr("sb.description AS systembank__description").
		ColumnExpr("json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description) AS image").
		ColumnExpr("i2.id AS imagesystembank__id").
		ColumnExpr("i2.ref_id AS imagesystembank__ref_id").
		ColumnExpr("i2.type AS imagesystembank__type").
		ColumnExpr("i2.description AS imagesystembank__description").
		Join("LEFT JOIN system_banks AS sb ON sb.id = p.system_bank_id").
		Join("LEFT JOIN Images AS i2 ON i2.ref_id = sb.id AND i2.type = 'systembank_image'").
		Join("LEFT JOIN images AS i ON i.ref_id = p.id AND i.type = 'payment_slip'").
		GroupExpr("p.id, sb.id,  i.id, i.ref_id, i.type, i.description, i2.id").
		Where("p.id = ?", id).Scan(ctx, payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func CreatePaymentService(ctx context.Context, req requests.PaymentCreateRequest) (*model.Payments, error) {

	payment := &model.Payments{
		SystemBankID: req.SystemBankID,
		Date:         req.Date,
	}
	payment.SetCreatedNow()
	payment.SetUpdateNow()

	_, err := db.NewInsert().Model(payment).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// img := requests.ImageCreateRequest{
	// 	RefID:       payment.ID,
	// 	Type:        "payment_slip",
	// 	Description: req.PaymentSlip,
	// }

	// _, err = image.CreateImageService(ctx, img)
	// if err != nil {
	// 	return nil, err
	// }

	return payment, nil

}

func UpdatePaymentService(ctx context.Context, id int64, req requests.PaymentUpdateRequest) (*model.Payments, error) {
	ex, err := db.NewSelect().Table("orders").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("order not found")
	}
	// ตรวจสอบว่ามี Payment อยู่ในระบบหรือไม่
	payment := &model.Payments{
		Date: req.Date,
	}
	payment.SetCreatedNow()
	payment.SetUpdateNow()

	_, err = db.NewInsert().Model(payment).Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.NewUpdate().Table("orders").
		Set("payment_id = ?", payment.ID).
		Set("status =  ?", "paid").
		Set("total_price_ship = ?", req.TotalPriceShip).
		Where("id = ?", id).Exec(ctx)
	return payment, err
}

func DeletePaymentService(ctx context.Context, id int64) error {
	// ตรวจสอบว่ามี Payment อยู่หรือไม่
	exists, err := db.NewSelect().TableExpr("payments").Where("id=?", id).Exists(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("payment not found")
	}

	// ลบรูปภาพที่เกี่ยวข้อง
	_, err = db.NewDelete().
		TableExpr("images").
		Where("ref_id = ? AND type = 'payment_slip'", id).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to delete")
	}

	// ลบ Payment
	_, err = db.NewDelete().TableExpr("payments").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return errors.New("failed to delete payment slip")
	}

	return nil
}
