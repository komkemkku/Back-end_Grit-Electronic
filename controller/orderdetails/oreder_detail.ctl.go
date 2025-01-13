package orederdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func OrderDetailList(c *gin.Context) {
	req := requests.OrderDetailRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListOrderDetailService(c.Request.Context(), requests.OrderDetailRequest(req))
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

func CreateOrderDetail(c *gin.Context) {
	req := requests.OrderDetailCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateOrderDetailService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteOrderDetail(c *gin.Context) {
	// รับค่า id จาก request และทำการ bind
	id := requests.OrderDetailIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// เรียกใช้ DeleteOrderDetailService พร้อม id ที่แปลงแล้ว
	err := DeleteOrderDetailService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}

func GetOrderDetailByID(c *gin.Context) {
	id := requests.OrderDetailIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByOrderDetailService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UpdateOrderDetaill(c *gin.Context) {
	id := requests.OrderDetailIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.OrderDetailUpdateRequest{}
if err := c.ShouldBindJSON(&req); err != nil {
    response.BadRequest(c, err.Error())
    return
}

data, err := UpdateOrderDetailService(c, int64(id.ID), req)
if err != nil {
    response.InternalError(c, err.Error())
    return
}

response.Success(c, data)
}