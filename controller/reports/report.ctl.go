package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func Dashboard(c *gin.Context) {

	data, err := GetDashboard(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}

func Report(c *gin.Context) {

	req := requests.ReportRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := GetReport(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	paginate := model.Paginate{
		Page:  req.Page,
		Size:  req.Size,
		Total: int64(total),
	}

	response.SuccessWithPaginate(c, data, paginate)
}

func DashboardTotalByCategory(c *gin.Context) {
    // สร้าง ReportRequest จากข้อมูลใน query parameters หรือ body
    var req requests.ReportRequest
    if err := c.BindQuery(&req); err != nil {
        response.BadRequest(c, err.Error())
        return
    }

    // เรียกฟังก์ชัน DashboardByCategory พร้อมทั้งส่ง parameters ที่จำเป็น
    data, _, err := DashboardByCategory(c.Request.Context(), req) // ใช้แค่ 2 ตัวแปร
    if err != nil {
        response.InternalError(c, err.Error())
        return
    }
	response.Success(c, data)
}





