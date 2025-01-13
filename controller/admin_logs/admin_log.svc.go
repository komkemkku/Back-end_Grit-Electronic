package adminlogs

import (
	"context"

	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
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

	// สร้าง query
	query := db.NewSelect().
		TableExpr("admin_logs AS al").
		Column("al.id", "al.action", "al.description", "al.created_at").
		ColumnExpr("a.id AS admin__id").
		ColumnExpr("a.name AS admin__name").
		Join("LEFT JOIN admins as a ON a.id = al.admin_id")

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
