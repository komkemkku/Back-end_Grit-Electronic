package users

import (
	"context"
	"errors"

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
	req := requests.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateUsersService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}