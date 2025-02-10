package reports

import (
	"context"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
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
