package adminlogs

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func AdminLogList(c *gin.Context) {
	req := requests.AdminLogRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListAdminLogsService(c.Request.Context(), req)
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

func CreateAdminLogs(c *gin.Context) {
    req := requests.AdminLogCreateRequest{}

    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, err.Error())
        return
    }

    err := CreateAdminLog(c.Request.Context(), req.AdminID, req.Action, req.Description)
    if err != nil {
        response.InternalError(c, err.Error())
        return
    }

    // ส่งข้อความตอบกลับเมื่อสำเร็จ
    response.Success(c, "Admin log created successfully")
}

