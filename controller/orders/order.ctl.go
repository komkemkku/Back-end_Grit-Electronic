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

	// บันทึก Log สำหรับคำสั่งซื้อที่สร้างขึ้น
	logMessage := fmt.Sprintf("ผู้ใช้งาน ID: %d ได้ทำการสั่งซื้อสินค้า หมายเลขคำสั่งซื้อ #%d", req.UserID, data.ID)
	if err := adminlogs.CreateAdminLog(c.Request.Context(), user, "CREATE_ORDER", logMessage); err != nil {
		fmt.Println("Failed to create admin log:", err)
	}

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

	user := c.GetInt("user_id")

	data, err := GetByIdOrderService(c, int64(id.ID), int64(user))
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

func OrderUserPaidList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.OrderUserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListOrderUserPaidService(c.Request.Context(), req)
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

	// ไม่บันทึก Log หากสถานะถูกเปลี่ยนเป็น "success"
	if req.Status != "success" {
		logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสถานะคำสั่งซื้อหมายเลข #%d เป็น: %s", AdminID, id.ID, req.Status)
		if err := adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_ORDER", logMessage); err != nil {
			fmt.Println("Failed to create admin log:", err)
		}
	}

	response.Success(c, data)
}

func UpdateShipOrder(c *gin.Context) {
	id := requests.OrderIdRequest{}
	user := c.GetInt("user_id")

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.OrderUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateShipOrderService(c, int(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// ไม่บันทึก Log หากสถานะถูกเปลี่ยนเป็น "success"
	if req.Status != "success" {
		logMessage := fmt.Sprintf("ผู้ใช้งาน ID: %d แก้ไขช่องทางจัดส่งคำสั่งซื้อหมายเลข #%d", user, id.ID)
		if err := adminlogs.CreateAdminLog(c.Request.Context(), user, "UPDATE_ORDER", logMessage); err != nil {
			fmt.Println("Failed to create admin log:", err)
		}
	}

	response.Success(c, data)
}
