package reports

import (
	"context"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func GetDashboard(ctx context.Context) (*response.DashboardResponse, error) {

	var totalSales float64
	var totalOrders int
	var totalUsers int
	var totalCancelled int

	// คำนวณจำนวนผู้ใช้งานทั้งหมดจากตาราง users
	totalUsers, err := db.NewSelect().
		Table("users").
		Count(ctx)
	if err != nil {
		return nil, err
	}

	// คำนวณจำนวนคำสั่งซื้อทั้งหมดจากตาราง orders
	totalOrders, err = db.NewSelect().
		Table("orders").
		Count(ctx)
	if err != nil {
		return nil, err
	}

	// คำนวณยอดขายทั้งหมดจากตาราง orders โดยกรองเฉพาะคำสั่งซื้อที่มี status เป็น "shipped"
	err = db.NewSelect().
		TableExpr("orders AS o").
		ColumnExpr("SUM(o.total_price) AS total_sales").
		Where("o.status = ?", "shipped").
		Scan(ctx, &totalSales)
	if err != nil {
		return nil, err
	}

	// คำนวณจำนวนสินค้าที่ถูกยกเลิกจากตาราง orders โดยกรองเฉพาะคำสั่งซื้อที่มี status เป็น "cancelled"
	totalCancelled, err = db.NewSelect().
		Table("orders").
		Where("status = ?", "cancelled").
		Count(ctx) // ใช้ Count() เพื่อทำการนับ
	if err != nil {
		return nil, err
	}

	// ส่งข้อมูลในรูปแบบที่ต้องการ
	result := &response.DashboardResponse{
		TotalSales:     totalSales,
		TotalOrders:    totalOrders,
		TotalUsers:     totalUsers,
		TotalCancelled: totalCancelled,
	}

	return result, nil

}

func GetReport(ctx context.Context, req requests.ReportRequest) ([]response.ReportReponses, int, error) {

	// var Offset int64
	// if req.Page > 0 {
	// 	Offset = (req.Page - 1) * req.Size
	// }

	resp := []response.ReportReponses{}

	// สร้าง query สำหรับดึงข้อมูลจาก order_details
	query := db.NewSelect().
		TableExpr("orders AS o").
		ColumnExpr("o.id AS order_id").
		ColumnExpr("o.created_at AS created_at").
		ColumnExpr("o.total_price AS total_price").
		ColumnExpr("u.username AS username").
		ColumnExpr("od.product_name AS product_name").
		ColumnExpr("od.total_product_amount AS amount").
		ColumnExpr("od.total_product_price AS price").
		Join("JOIN users AS u ON o.user_id = u.id").
		Join("JOIN order_details AS od ON o.id = od.order_id").
		Scan(ctx, &resp)

	// ตรวจสอบหากเกิดข้อผิดพลาดใน query
	if query != nil {
		return nil, 0, query
	}

	// สร้าง query สำหรับหาจำนวนแถวใน order_details
	var total int
	countQuery := db.NewSelect().
		TableExpr("order_details AS od").
		Join("JOIN orders AS o ON o.id = od.order_id").
		Join("JOIN users AS u ON o.user_id = u.id").
		ColumnExpr("COUNT(od.id)").
		Scan(ctx, &total)

	// ตรวจสอบหากเกิดข้อผิดพลาดใน query การนับจำนวน
	if countQuery != nil {
		return nil, 0, countQuery
	}

	// ดำเนินการ query หลักเพื่อดึงข้อมูลที่ต้องการ
	// err := query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	// if err != nil {
	// 	return nil, 0, err
	// }

	return resp, total, nil
}

// func DashboardByCategory(ctx context.Context) ([]response.DashboardCategoryResponses, error) {

//     // สร้างแผนที่สำหรับเก็บยอดขายตามประเภทสินค้า
//     categorySales := make([]response.DashboardCategoryResponses, 0)

//     // คำนวณยอดขายในแต่ละประเภทสินค้า
//     rows, err := db.NewSelect().
//         ColumnExpr("p.category, SUM(oi.quantity * oi.price) AS total_sales").
//         Table("order_items AS oi").
//         Join("products AS p ON oi.product_id = p.product_id").
//         Join("orders AS o ON oi.order_id = o.order_id").
//         Where("o.status = ?", "shipped").
//         GroupBy("p.category"). // ใช้ GroupBy() ในการจัดกลุ่มข้อมูล
//         Rows(ctx) // ใช้ Rows() เพื่อดึงข้อมูลหลายแถวที่มีการจัดกลุ่มตามประเภทสินค้า

//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     // แปลงข้อมูลจาก rows มาเป็น slice ของ DashboardCategoryResponses
//     for rows.Next() {
//         var category string
//         var totalSales float64
//         if err := rows.Scan(&category, &totalSales); err != nil {
//             return nil, err
//         }

//         // สร้าง response และเพิ่มเข้าใน slice
//         categorySales = append(categorySales, response.DashboardCategoryResponses{
//             Category:   category,
//             TotalSales: totalSales,
//         })
//     }

//     return categorySales, nil
// }
