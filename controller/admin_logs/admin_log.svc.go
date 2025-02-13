package adminlogs

import (
	"context"
	"log"

	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = config.Database()

func ListAdminLogsService(ctx context.Context, req requests.AdminLogRequest) ([]response.AdminLogResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.AdminLogResponses{}

	// ใช้ค่า StartDate และ EndDate ตรง ๆ เพราะเป็น int64 อยู่แล้ว
	var startUnix, endUnix int64

	if req.StartDate > 0 {
		startUnix = req.StartDate
	}

	if req.EndDate > 0 {
		endUnix = req.EndDate
	}

	// ตรวจสอบค่าที่ได้รับ
	log.Printf("Filtering logs from %d to %d\n", startUnix, endUnix)

	// สร้าง query
	query := db.NewSelect().
		TableExpr("admin_logs AS al").
		Column("al.id", "al.action", "al.description", "al.created_at").
		ColumnExpr("a.id AS admin__id").
		ColumnExpr("a.name AS admin__name").
		Join("LEFT JOIN admins as a ON a.id = al.admin_id")

	if req.Search != "" {
		query.Where("a.name ILIKE ?", "%"+req.Search+"%")
	}

	// กรองตามช่วงวันที่ (Unix Timestamp)
	if startUnix > 0 && endUnix > 0 {
		if startUnix == endUnix {
			query.Where("DATE(TO_TIMESTAMP(al.created_at)) = DATE(TO_TIMESTAMP(?))", startUnix)
		} else {
			query.Where("al.created_at BETWEEN EXTRACT(EPOCH FROM TO_TIMESTAMP(?)) AND EXTRACT(EPOCH FROM TO_TIMESTAMP(?))", startUnix, endUnix)
		}
	} else if startUnix > 0 {
		query.Where("DATE(TO_TIMESTAMP(al.created_at)) = DATE(TO_TIMESTAMP(?))", startUnix)
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

func CreateAdminLog(ctx context.Context, adminID int, action, description string) error {
	adminLog := &model.AdminLogs{
		AdminID:     adminID,
		Action:      action,
		Description: description,
	}

	adminLog.SetCreatedNow()

	_, err := db.NewInsert().Model(adminLog).Exec(ctx)
	return err
}
