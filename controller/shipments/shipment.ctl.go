package shipments

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateShipment(c *gin.Context) {
	req := requests.ShipmentCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateShipmentService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteShipment(c *gin.Context) {
	id := requests.ShipmentIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteShipmentService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}

func GetShipmentByID(c *gin.Context) {
	id := requests.ShipmentIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdShipmentService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func ShipmenttList(c *gin.Context) {
	req := requests.ShipmentRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListShipmentService(c.Request.Context(), req)
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

func UpdateShipment(c *gin.Context) {
	id := requests.ShipmentIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.ShipmentUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateShipmentService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}