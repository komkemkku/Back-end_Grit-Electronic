package orders

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateOrder(c *gin.Context) {

	user := c.GetInt("user_id")

	req := requests.OrderCreateRequest{}

	req.UserID = user

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateOrderService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	logMessage := fmt.Sprintf("ผู้ใช้งาน ID : %d สั่งซื้อสินค้า", user)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), user, "CREATE_CATEGORIES", logMessage)
	response.Success(c, data)
}

// func DeleteOrder(c *gin.Context) {
// 	id := requests.OrderIdRequest{}
// 	if err := c.BindUri(&id); err != nil {
// 		response.BadRequest(c, err.Error())
// 		return
// 	}
// 	err := DeleteOrderService(c, int64(id.ID))
// 	if err != nil {
// 		response.InternalError(c, err.Error())
// 		return
// 	}

// 	response.Success(c, "delete successfully")
// }

func GetOrderByID(c *gin.Context) {

	id := requests.OrderIdRequest{}

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdOrderService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func OrderList(c *gin.Context) {
	req := requests.OrderRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListOrderService(c.Request.Context(), req)
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

// func OrderUserList(c *gin.Context) {

// 	user := c.GetInt("user_id")
// 	req := requests.OrderUserRequest{}
// 	if err := c.BindQuery(&req); err != nil {
// 		response.BadRequest(c, err.Error())
// 		return
// 	}
// 	req.UserID = user

// 	data, total, err := ListOrderUserService(c.Request.Context(), req)
// 	if err != nil {
// 		response.InternalError(c, err.Error())
// 		return
// 	}

// 	paginate := model.Paginate{
// 		Page:  req.Page,
// 		Size:  req.Size,
// 		Total: int64(total),
// 	}

// 	response.SuccessWithPaginate(c, data, paginate)
// }

// Controller สำหรับสถานะ pending
func OrderUserPendingList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserPendingService(c.Request.Context(), req)
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

// Controller สำหรับสถานะ prepare
func OrderUserPrepareList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserPrepareService(c.Request.Context(), req)
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

// Controller สำหรับสถานะ ship
func OrderUserShipList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserShipService(c.Request.Context(), req)
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

// Controller สำหรับสถานะ success
func OrderUserSuccessList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserSuccessService(c.Request.Context(), req)
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

// Controller สำหรับสถานะ failed
func OrderUserFailedList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserFailedService(c.Request.Context(), req)
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

// Controller สำหรับสถานะ cancelled
func OrderUserCancelledList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserCancelledService(c.Request.Context(), req)
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

func OrderUserHistoryList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserHistoryService(c.Request.Context(), req)
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

func UpdateOrder(c *gin.Context) {
	id := requests.OrderIdRequest{}
	AdminID := c.GetInt("admin_id")
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.OrderUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateOrderService(c, int64(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	logMessage := fmt.Sprintf("แอดมิน ID : %d แก้ไขสถานะ : %s", AdminID, req.Status)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "CREATE_CATEGORIES", logMessage)

	response.Success(c, data)
}
