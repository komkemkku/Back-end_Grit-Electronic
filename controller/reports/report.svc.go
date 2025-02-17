package reports

import (
	"context"
	"time"

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

	// คำนวณยอดขายทั้งหมดจากตาราง orders โดยกรองเฉพาะคำสั่งซื้อที่มี status เป็น "success"
	err = db.NewSelect().
		TableExpr("orders AS o").
		ColumnExpr("SUM(o.total_price) AS total_sales").
		Where("o.status = ?", "success").
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

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

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
		Join("JOIN order_details AS od ON o.id = od.order_id")

	// เพิ่มตัวกรองเดือนและปี
	if req.Month != "" {
		query.Where("(TRIM(TO_CHAR(TO_TIMESTAMP(o.created_at), 'Month')) ILIKE ? OR "+
			"TO_CHAR(TO_TIMESTAMP(o.created_at), 'Mon') ILIKE ? OR "+
			"TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = ? OR "+
			"CASE "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '01' THEN 'มกราคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '02' THEN 'กุมภาพันธ์' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '03' THEN 'มีนาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '04' THEN 'เมษายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '05' THEN 'พฤษภาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '06' THEN 'มิถุนายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '07' THEN 'กรกฎาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '08' THEN 'สิงหาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '09' THEN 'กันยายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '10' THEN 'ตุลาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '11' THEN 'พฤศจิกายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '12' THEN 'ธันวาคม' "+
			"END ILIKE ?)", req.Month, req.Month, req.Month, req.Month)
	}

	if req.Year > 0 {
		query.Where("EXTRACT(YEAR FROM TO_TIMESTAMP(o.created_at)) = ?", req.Year)
	}

	// นับจำนวนแถวทั้งหมด
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

	// ดึงข้อมูลพร้อม pagination
	err := query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func DashboardlistCategorye(ctx context.Context, req requests.ReportRequest) ([]response.DashboardCategoryResponses, int, error) {
	categorySales := make([]response.DashboardCategoryResponses, 0)

	// เริ่มต้นการคิวรีข้อมูล
	query := db.NewSelect().
    TableExpr("order_details AS od"). // ใช้ตารางที่ถูกต้อง
    ColumnExpr("c.category_name, EXTRACT(YEAR FROM TO_TIMESTAMP(o.created_at)) AS year, EXTRACT(MONTH FROM TO_TIMESTAMP(o.created_at)) AS month, p.product_name, SUM(o.total_price) AS total_sales").
    Join("JOIN products p ON p.id = od.product_code"). // ใช้คอลัมน์ที่ถูกต้อง
    Join("JOIN categories c ON c.id = p.category_id").
    Join("JOIN orders o ON o.id = od.order_id").
    Where("o.status = ?", "success").
    Group("c.category_name, year, month, p.product_name")

	// กรองตามเดือน
	if req.Month != "" {
		query.Where("(TRIM(TO_CHAR(TO_TIMESTAMP(o.created_at), 'Month')) ILIKE ? OR "+
			"TO_CHAR(TO_TIMESTAMP(o.created_at), 'Mon') ILIKE ? OR "+
			"TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = ? OR "+"CASE "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '01' THEN 'มกราคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '02' THEN 'กุมภาพันธ์' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '03' THEN 'มีนาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '04' THEN 'เมษายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '05' THEN 'พฤษภาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '06' THEN 'มิถุนายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '07' THEN 'กรกฎาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '08' THEN 'สิงหาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '09' THEN 'กันยายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '10' THEN 'ตุลาคม' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '11' THEN 'พฤศจิกายน' "+
			"  WHEN TO_CHAR(TO_TIMESTAMP(o.created_at), 'MM') = '12' THEN 'ธันวาคม' "+
			"END ILIKE ?)", req.Month, req.Month, req.Month, req.Month)
	}

	// กรองตามปี
	if req.Year > 0 {
		query.Where("EXTRACT(YEAR FROM TO_TIMESTAMP(o.created_at)) = ?", req.Year)
	}

	// ดึงข้อมูลจากฐานข้อมูล
	rows, err := query.Rows(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// ทำการประมวลผลข้อมูลที่ดึงมา
	for rows.Next() {
		var category string
		var yearFloat, monthFloat float64
		var productName string
		var totalSales float64

		if err := rows.Scan(&category, &yearFloat, &monthFloat, &productName, &totalSales); err != nil {
			return nil, 0, err
		}

		year := int(yearFloat)
		month := int(monthFloat)
		monthName := time.Month(month).String()

		// ตรวจสอบว่า Category, Year และ Month นี้มีข้อมูลอยู่แล้วในผลลัพธ์หรือไม่
		var found bool
		for i, cat := range categorySales {
			if cat.Category == category && cat.Year == year && cat.Month == monthName {
				categorySales[i].Products = append(cat.Products, response.ProductSales{
					ProductName: productName,
					TotalSales:  totalSales,
				})
				found = true
				break
			}
		}

		// ถ้าไม่พบให้เพิ่มข้อมูลใหม่
		if !found {
			categorySales = append(categorySales, response.DashboardCategoryResponses{
				Category: category,
				Year:     year,
				Month:    monthName,
				Products: []response.ProductSales{
					{
						ProductName: productName,
						TotalSales:  totalSales,
					},
				},
			})
		}
	}

	return categorySales, len(categorySales), nil
}

