package users

import (
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils"
)

var db = config.Database()

const DefaultRoleID = "3"

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
	if req.RoleID == "" {
		req.RoleID = DefaultRoleID
	}

	ex, err := db.NewSelect().TableExpr("roles").Where("id =?", req.RoleID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("role not found")
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

	// แปลง req.Role_id จาก string เป็น int64
	roleID, err := strconv.ParseInt(req.RoleID, 10, 64)
	if err != nil {
		return nil, errors.New("invalid role_id format: " + err.Error())
	}

	// ใช้ roleID ที่แปลงแล้วใน struct
	userRole := &model.UserRole{
		UserID:    user.ID,
		RoleID:    roleID,
	}
	

	_, err = db.NewInsert().Model(userRole).Exec(ctx)
	if err != nil {
		return nil, errors.New("failed to assign role to user: " + err.Error())
	}

	return user, nil
}
