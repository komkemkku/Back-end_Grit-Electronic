package cartitems

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateCartItem(c *gin.Context) {
	req := requests.CartItemCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateCartItemService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteCartItem(c *gin.Context) {
	id := requests.CartItemIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteCartItemService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}


func GetCartItemByID(c *gin.Context) {
	id := requests.CartItemIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdCartItemService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CartItemList(c *gin.Context) {
	req := requests.CartItemRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListCartItemService(c.Request.Context(), req)
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

func UpdateCartItem(c *gin.Context) {
	id := requests.CartItemIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.CartItemUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateCartItemService(c, int64(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}