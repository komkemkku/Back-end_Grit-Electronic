package orders

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateOrder(c *gin.Context) {
	req := requests.OrderCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateOrderService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteOrder(c *gin.Context) {
	id := requests.OrderIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteOrderService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}

func GetOrderByID(c *gin.Context) {
	id := requests.OrderIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdOrderService(c, id.ID)
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

func UpdateOrder(c *gin.Context) {
	id := requests.OrderIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.OrderUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateOrderService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}
