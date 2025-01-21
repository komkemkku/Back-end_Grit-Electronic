package wishlist

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func Wishlist(c *gin.Context) {
	var req requests.WishlistsRequest
	// ตรวจสอบ Query Parameters
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, "Invalid query parameters: "+err.Error())
		return
	}

	// เรียกใช้ Service
	data, total, err := ListWishlistsService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "Failed to fetch wishlists: "+err.Error())
		return
	}

	// จัดการ Pagination
	paginate := model.Paginate{
		Page:  req.Page,
		Size:  req.Size,
		Total: int64(total),
	}

	// ส่ง Response กลับ
	response.SuccessWithPaginate(c, data, paginate)
}

func CreateWishlist(c *gin.Context) {
	var req requests.WishlistsAddRequest

	// ตรวจสอบ JSON Input
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	// เรียกใช้ Service
	if err := CreateWishlistsService(c.Request.Context(), req); err != nil {
		response.InternalError(c, "Failed to create wishlist: "+err.Error())
		return
	}

	// ส่ง Response กลับ
	response.Success(c, "Wishlist created successfully")
}

func GetWishlistByID(c *gin.Context) {
	id := requests.WishlistsIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdWishlistsService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}
func DeleteWishlists(c *gin.Context) {
	// รับค่า id จาก request และทำการ bind
	id := requests.WishlistsIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// เรียกใช้ DeleteWishlistsService พร้อม id ที่แปลงแล้ว
	err := DeleteWishlistsService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}

func UpdateWishlists(c *gin.Context) {
	id := requests.WishlistsIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.WishlistsUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := UpdateWishlistsService(c, int64(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return

	}
	response.Success(c, data)
}
