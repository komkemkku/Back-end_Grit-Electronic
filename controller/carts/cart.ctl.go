package carts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

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
	id := requests.CartIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdCartService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CartList(c *gin.Context) {
	req := requests.CartRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// เพิ่มตัวแปรเพื่อรับ grandTotal
	data, total, grandTotal, err := ListCartService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	fmt.Printf("Grand Total for Cart: %.2f\n", grandTotal)


	// เพิ่ม grandTotal เข้าไปใน Response หากต้องการ
	paginate := model.Paginate{
		Page:       req.Page,
		Size:       req.Size,
		Total:      int64(total),
	}

	response.SuccessWithPaginate(c, data, paginate)
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