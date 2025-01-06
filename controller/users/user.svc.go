package users

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils"
)

var db = config.Database()

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

func CreateUsersService(ctx context.Context, req requests.UserCreateRequest) (*model.Users, error) {

	// ตรวจสอบว่า Role ID = 3 มีอยู่ในตาราง roles หรือไม่
	ex, err := db.NewSelect().TableExpr("roles").Where("id = ?", 3).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("default role not found (role_id = 3)")
	}

	hashpassword, _ := utils.HashPassword(req.Password)
	user := &model.Users{
		Username:    req.Username,
		Password:    hashpassword,
		Email:       req.Email,
		Phone:       req.Phone,
		BankNumber: req.BankNumber,
		BankName: req.BankName,
	}
	user.SetCreatedNow()
	user.SetUpdateNow()

	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// กำหนด Role เป็น 3 โดยอัตโนมัติ
	userRole := &model.UserRole{
		UserID: user.ID,
		RoleID: 3,
	}

	_, err = db.NewInsert().Model(userRole).Exec(ctx)
	if err != nil {
		return nil, errors.New("failed to assign role to user: " + err.Error())
	}

	return user, nil
}