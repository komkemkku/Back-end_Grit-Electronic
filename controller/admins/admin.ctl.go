package admins

import (
	"github.com/gin-gonic/gin"
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
	req := requests.AdminCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateAdminService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Admin Create successfully")
}

func DeleteAdmin(c *gin.Context) {
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

	response.Success(c, "delete successfully")
}

func UpdateAdmin(c *gin.Context) {
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

	// ส่งข้อความ Success โดยไม่ต้องส่งข้อมูลกลับ
	response.Success(c, "Admin updated successfully")
}
