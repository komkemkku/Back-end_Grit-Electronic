package wishlist

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func Wishlist(c *gin.Context) {
	user := c.GetInt("user_id")
	var req requests.WishlistsRequest
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, "Invalid query parameters: "+err.Error())
		return
	}
	req.UserID = user

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

	user := c.GetInt("user_id")

	var req requests.WishlistsAddRequest

	req.UserID = user

	// ตรวจสอบ JSON Input
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	if err := CreateWishlistsService(c.Request.Context(), req); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// ส่ง Response กลับ
	response.Success(c, "Wishlist created successfully")
}

// func GetWishlistByID(c *gin.Context) {
// 	user := c.GetInt("user_id")

// 	id := requests.WishlistsIdRequest{}
// 	if err := c.BindUri(&id); err != nil {
// 		response.BadRequest(c, err.Error())
// 		return
// 	}

// 	data, err := GetByIdWishlistsService(c, id.ID, user)
// 	if err != nil {
// 		response.InternalError(c, err.Error())
// 		return
// 	}

// 	response.Success(c, data)
// }

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
	userID := c.GetInt("user_id")

	req := requests.WishlistsUpdateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request: "+err.Error())
		return
	}

	wish, err := UpdateWishlistsService(c.Request.Context(), int64(userID), int64(req.ProductID))
	if err != nil {
		response.InternalError(c, "Failed to update wishlist: "+err.Error())
		return
	}

	if wish == nil {
		// ถ้าสินค้าถูกลบออกจาก Wishlist
		response.Success(c, "Product removed from wishlist")
	} else {
		// ถ้าสินค้าเพิ่มเข้า Wishlist
		response.Success(c, wish)
	}
}

