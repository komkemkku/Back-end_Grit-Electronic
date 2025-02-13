package systembank

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateSystembank(c *gin.Context) {
	req := requests.SystemBankCreateRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateSystemBankService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID : %d เพิ่มบัญชีธนาคารของระบบ", AdminID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "ADD_SYSTEMBANK", logMessage)

	response.Success(c, "system bank created successfully")
}

func DeleteSystemBank(c *gin.Context) {
	id := requests.SystemBankIdRequest{}
	AdminID := c.GetInt("admin_id")
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteSystemBankService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID : %d ลบบัญชีธนาคารของระบบ", AdminID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_SYSTEMBANK", logMessage)

	response.Success(c, "delete successfully")
}

func GetSystemBankByID(c *gin.Context) {
	id := requests.SystemBankIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdSystemBankService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func SystemBankList(c *gin.Context) {
	req := requests.SystemBankRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListSystemBankService(c.Request.Context(), req)
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

func UpdateSystemBank(c *gin.Context) {
	id := requests.SystemBankIdRequest{}
	AdminID := c.GetInt("admin_id")
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.SystemBankUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateSystembankService(c, int64(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	logMessage := fmt.Sprintf("แอดมิน ID : %d แก้ไขบัญชีธนาคารของระบบ", AdminID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_SYSTEMBANK", logMessage)

	response.Success(c, data)
}
