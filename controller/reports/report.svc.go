package reports

import (
	"context"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func GetDashboard(ctx context.Context) (*response.DashboardResponse, error) {

	// var totalSales float64
	// var totalOrders int
	var totalUsers int

	// คำนวณจำนวนผู้ใช้งานทั้งหมดจากตาราง users
	err := db.NewSelect().
		ColumnExpr("COUNT(u.id) AS totaluser").
		// Join("LEFT JOIN users AS u ON u.id = o.user_id").
		Scan(ctx, &totalUsers)
	if err != nil {
		return nil, err
	}

	// ส่งข้อมูลในรูปแบบที่ต้องการ
	result := &response.DashboardResponse{
		// TotalSales:  totalSales,
		// TotalOrders: totalOrders,
		TotalUsers: totalUsers,
	}

	return result, nil

}
