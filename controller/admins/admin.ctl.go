package admins

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func GetInfoAdmin(c *gin.Context) {
	admin := c.GetInt("admin_id")

	data, err := GetByIdAdminService(c, admin)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}

func AdmintList(c *gin.Context) {
	req := requests.AdminRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListAdminService(c.Request.Context(), req)
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

func GetAdminByID(c *gin.Context) {
	id := requests.AdminIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdAdminService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateAdmin(c *gin.Context) {
	AdminID := c.GetInt("admin_id")
	req := requests.AdminCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	newAdmin, err := CreateAdminService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// Log ว่ามีการสร้างแอดมิน
	_ = adminlogs.CreateAdminLog(
		c.Request.Context(),
		AdminID,
		"CREATE_ADMIN",
		fmt.Sprintf("แอดมิน ID : %d เพิ่มแอดมินคนใหม่ ID : %d ชื่อ : %s",
			AdminID, newAdmin.ID, newAdmin.Name),
	)

	response.Success(c, "Admin created successfully")
}

func DeleteAdmin(c *gin.Context) {
	AdminID := c.GetInt("admin_id")
	id := requests.AdminIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteAdminService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	_ = adminlogs.CreateAdminLog(
		c.Request.Context(),
		AdminID,
		"DELETE_ADMIN",
		fmt.Sprintf("แอดมิน ID : %d ลบแอดมิน ID : %d", AdminID, id.ID),
	)

	response.Success(c, "delete successfully")
}

func UpdateAdmin(c *gin.Context) {
	AdminID := c.GetInt("admin_id")
	id := requests.AdminIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.AdminUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := UpdateAdminService(c, int64(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	_ = adminlogs.CreateAdminLog(
		c.Request.Context(),
		AdminID,
		"UPDATE_ADMIN",
		fmt.Sprintf("แอดมิน ID : %d แก้ไขแอดมิน ID : %d", AdminID, id.ID),
	)

	response.Success(c, "Admin updated successfully")
}
