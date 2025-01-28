package cartitems

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateCartItem(c *gin.Context) {

	user := c.GetInt("user_id")

	req := requests.CartItemCreateRequest{}

	req.UserID = user

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
	req := requests.CartItemDeleteRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.CartID == 0 || req.UserID == 0 || req.CartItemID == 0 {
		response.BadRequest(c, "Invalid cart_id, user_id, or cart_item_id")
		return
	}

	err := DeleteCartItemService(c.Request.Context(), req.CartID, req.UserID, req.CartItemID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, map[string]interface{}{
		"message":      "Cart item deleted successfully",
		"cart_id":      req.CartID,
		"user_id":      req.UserID,
		"cart_item_id": req.CartItemID,
	})
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
