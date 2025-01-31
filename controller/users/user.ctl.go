package users

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func GetInfoUser(c *gin.Context) {
	user := c.GetInt("user_id")

	data, err := GetByIdUserService(c, user)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}

func UserList(c *gin.Context) {
	req := requests.UserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListUserService(c.Request.Context(), req)
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

func GetUserByID(c *gin.Context) {
	id := requests.UserIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdUserService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateUser(c *gin.Context) {
	req := requests.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateUsersService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "users created successfully")
}

// func CreateUser(c *gin.Context) {
//   var req requests.UserCreateRequest
//   if err := c.ShouldBindJSON(&req); err != nil {
//     response.BadRequest(c, err.Error())
//     return
//   }

//   // เรียกใช้ service โดยรับค่า return แค่ error
//   err := CreateUsersService(c, req)
//   if err != nil {
//     response.InternalError(c, err.Error())
//     return
//   }

// 	// ถ้าทำงานสำเร็จ
// 	response.Success(c, http.StatusOK)
// }

// func Success(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{
//     "status": gin.H{
//       "code":    http.StatusOK,
//       "message": "Success",
//     },
//   })
// }

func UpdateUser(c *gin.Context) {
	id := requests.UserIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.UserUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := UpdateUserService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Updated successfully")
}

func DeleteUser(c *gin.Context) {
	id := requests.UserIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteUserService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}
