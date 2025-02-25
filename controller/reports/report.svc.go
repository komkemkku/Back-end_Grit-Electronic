package reports

import (
	"context"
	"fmt"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func GetDashboard(ctx context.Context, req requests.ReportRequest) (*response.DashboardResponse, error) {
	var totalSales float64
	var totalOrders, totalUsers, totalCancelled int

	// คำนวณจำนวนผู้ใช้งานทั้งหมด
	totalUsers, err := db.NewSelect().
		Table("users").
		Count(ctx)
	if err != nil {
		return nil, err
	}

	// กรองเดือนในรูปแบบต่างๆ
	monthFilter := ""
	if req.Month != "" {
		monthFilter = `
			(EXTRACT(MONTH FROM TO_TIMESTAMP(created_at))::TEXT = ? OR
			TO_CHAR(TO_TIMESTAMP(created_at), 'Mon') ILIKE ? OR
			TO_CHAR(TO_TIMESTAMP(created_at), 'Month') ILIKE ? OR
			CASE
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 1 THEN 'มกราคม'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 2 THEN 'กุมภาพันธ์'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 3 THEN 'มีนาคม'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 4 THEN 'เมษายน'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 5 THEN 'พฤษภาคม'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 6 THEN 'มิถุนายน'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 7 THEN 'กรกฎาคม'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 8 THEN 'สิงหาคม'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 9 THEN 'กันยายน'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 10 THEN 'ตุลาคม'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 11 THEN 'พฤศจิกายน'
				WHEN EXTRACT(MONTH FROM TO_TIMESTAMP(created_at)) = 12 THEN 'ธันวาคม'
			END ILIKE ?)`
	}

	// คำนวณจำนวนออเดอร์ทั้งหมด
	queryOrders := db.NewSelect().
		Table("orders")

	if monthFilter != "" {
		queryOrders.Where(monthFilter, req.Month, req.Month, req.Month, req.Month)
	}
	if req.Year > 0 {
		queryOrders.Where("EXTRACT(YEAR FROM TO_TIMESTAMP(created_at)) = ?", req.Year)
	}

	totalOrders, err = queryOrders.Count(ctx)
	if err != nil {
		return nil, err
	}

	// คำนวณยอดขายรวม โดยดูเฉพาะออเดอร์ที่สำเร็จ (status = success)
	querySales := db.NewSelect().
		TableExpr("orders o").
		ColumnExpr("COALESCE(SUM(o.total_price), 0) AS total_sales").
		Where("o.status = ?", "success")

	if monthFilter != "" {
		querySales.Where(monthFilter, req.Month, req.Month, req.Month, req.Month)
	}
	if req.Year > 0 {
		querySales.Where("EXTRACT(YEAR FROM TO_TIMESTAMP(o.created_at)) = ?", req.Year)
	}

	err = querySales.Scan(ctx, &totalSales)
	if err != nil {
		return nil, err
	}

	// คำนวณจำนวนออเดอร์ที่ถูกยกเลิก
	queryCancelled := db.NewSelect().
		Table("orders").
		Where("status = ?", "cancelled")

	if monthFilter != "" {
		queryCancelled.Where(monthFilter, req.Month, req.Month, req.Month, req.Month)
	}
	if req.Year > 0 {
		queryCancelled.Where("EXTRACT(YEAR FROM TO_TIMESTAMP(created_at)) = ?", req.Year)
	}

	totalCancelled, err = queryCancelled.Count(ctx)
	if err != nil {
		return nil, err
	}

	// ส่งข้อมูลกลับ
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

	// สร้าง query สำหรับดึงข้อมูลจาก order_details และรวมสินค้าเป็น JSON array
	query := db.NewSelect().
		TableExpr("orders AS o").
		ColumnExpr("o.id AS order_id").
		ColumnExpr("o.created_at AS created_at").
		ColumnExpr("o.total_price AS total_price").
		ColumnExpr("o.total_amount AS total_amount"). // Add total_amount
		ColumnExpr("u.username AS username").
		ColumnExpr("u.firstname AS firstname").
		ColumnExpr("u.lastname AS lastname").
		ColumnExpr("json_agg(json_build_object('product_name', od.product_name, 'amount', od.total_product_amount, 'price', od.total_product_price, 'total_product_amount', od.total_product_amount)) AS products"). // Ensure total_product_amount is part of the products JSON
		Join("JOIN users AS u ON o.user_id = u.id").
		Join("JOIN order_details AS od ON o.id = od.order_id").
		GroupExpr("o.id, o.created_at, o.total_price, o.total_amount, u.username, u.firstname, u.lastname")

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
		TableExpr("orders AS o").
		Join("JOIN users AS u ON o.user_id = u.id").
		ColumnExpr("COUNT(o.id)").
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
	categorySales := []response.DashboardCategoryResponses{}
  
	query := db.NewSelect().
	  TableExpr("order_details AS od").
	  ColumnExpr("c.name AS category, SUM(od.total_product_amount * p.price) AS total_category_sales").
	  Join("JOIN products AS p ON od.product_name = p.name").
	  Join("JOIN categories AS c ON c.id = p.category_id").
	  Join("JOIN orders AS o ON o.id = od.order_id").
	  Where("o.status = ?", "success").
	  GroupExpr("c.name")
  
	// เพิ่มเงื่อนไขการกรองตามเดือน (Month)
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
  
	rows, err := query.Rows(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()
  
	for rows.Next() {
	  var category string
	  var totalCategorySales float64
  
	  if err := rows.Scan(&category, &totalCategorySales); err != nil {
		return nil, 0, fmt.Errorf("failed to scan row: %v", err)
	  }
  
	  categorySales = append(categorySales, response.DashboardCategoryResponses{
		Category:           category,
		TotalCategorySales: totalCategorySales,
	  })
	}

	return categorySales, len(categorySales), nil
}
