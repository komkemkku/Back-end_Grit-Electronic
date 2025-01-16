package users

import (
  "context"
  "errors"
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
  "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
  "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func GetByIdUserService(ctx context.Context, id int64) (*model.Users, error) {
  ex, err := db.NewSelect().TableExpr("users").Where("id=?", id).Exists(ctx)
  if err != nil {
    return nil, err
  }
  if !ex {
    return nil, errors.New("user not found")
  }
  user := &model.Users{}

  err = db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
  if err != nil {
    return nil, err
  }
  return user, nil
}

func CreateUser(c *gin.Context) {
  var req requests.UserCreateRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    response.BadRequest(c, err.Error())
    return
  }

  // เรียกใช้ service โดยรับค่า return แค่ error
  err := CreateUsersService(c, req)
  if err != nil {
    response.InternalError(c, err.Error())
    return
  }

  // ถ้าทำงานสำเร็จ
  response.Success(c, http.StatusOK)
}

func Success(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "status": gin.H{
      "code":    http.StatusOK,
      "message": "Success",
    },
  })
}

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

  data, err := UpdateUserService(c, int64(id.ID), req)
  if err != nil {
    response.InternalError(c, err.Error())
    return
  }
  response.Success(c, data)
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
