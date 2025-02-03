package carts

import (

	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

// ไม่ได้ใช้
func AddCart(c *gin.Context) {
	req := requests.CartAddItemRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateCartService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteCart(c *gin.Context) {
	id := requests.CartIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteCartService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}

func GetCartByID(c *gin.Context) {

		userID := c.GetInt("user_id")
		if userID == 0 {
			response.BadRequest(c, "user ID is required")
			return
		}
	
		data, err := GetByIdCartService(c, int64(userID))
		if err != nil {
			response.InternalError(c, err.Error())
			return
		}
	
		response.Success(c, data)
	}
	


func UpdateCart(c *gin.Context) {
	id := requests.CartIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.CartUpdateItemRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateCartService(c, int64(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}