package wishlist

import (
	"strconv"

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
	productID, err := strconv.ParseInt(c.Param("product_id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "Invalid product ID")
		return
	}

	data, message, isFavorite, err := UpdateWishlistsService(c.Request.Context(), userID, int(productID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"message":     message,
		"is_favorite": isFavorite,
		"data":        data,
	})
}


func GetWishlistStatus(c *gin.Context) {
	userID := c.GetInt("user_id")
	productID, err := strconv.ParseInt(c.Param("product_id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "Invalid product ID")
		return
	}

	status, err := GetWishlistStatusService(c.Request.Context(), int64(userID), productID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"is_favorite": status})
}


