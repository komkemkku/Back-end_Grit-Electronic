package users

import (
	"context"
	"errors"
	"net/http"

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

	data, err := GetByIdUserService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateUsersService(ctx context.Context, req requests.UserCreateRequest) error {
	// ตรวจสอบว่า Role ID = 3 มีอยู่หรือไม่
	ex, err := db.NewSelect().TableExpr("roles").Where("id = ?", 3).Exists(ctx)
	if err != nil {
		return err
	}
	if !ex {
		return errors.New("default role not found (role_id = 3)")
	}

	// แฮชรหัสผ่าน
	hashpassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.New("failed to hash password: " + err.Error())
	}

	// สร้างผู้ใช้ใหม่
	user := &model.Users{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Username:  req.Username,
		Password:  hashpassword, // แฮชแล้ว
		Email:     req.Email,
		Phone:     req.Phone,
	}
	user.SetCreatedNow()
	user.SetUpdateNow()

	// เพิ่มผู้ใช้ในฐานข้อมูล
	if _, err := db.NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}

	// กำหนด Role เป็น 3 โดยอัตโนมัติ
	userRole := &model.UserRole{
		UserID: user.ID,
		RoleID: 3,
	}
	if _, err := db.NewInsert().Model(userRole).Exec(ctx); err != nil {
		return errors.New("failed to assign role to user: " + err.Error())
	}

	// ถ้าถึงตรงนี้แปลว่าสำเร็จ ไม่มี error
	return nil
}

func CreateUserHandler(c *gin.Context) {
	var req requests.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{
				"code":    400,
				"message": err.Error(),
			},
		})
		return
	}

	// เรียก Service ที่แก้ไขแล้ว (ไม่ return user แล้ว)
	if err := CreateUsersService(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": gin.H{
				"code":    500,
				"message": err.Error(),
			},
		})
		return
	}

	// ตรงนี้ส่งกลับเฉพาะ status ตามต้องการ
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    200,
			"message": "Success",
		},
	})
}

func UpdateUserService(ctx context.Context, id int64, req requests.UserUpdateRequest) (*model.Users, error) {

	ex, err := db.NewSelect().TableExpr("roles").Where("id = ?", 3).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("default role not found (role_id = 3)")
	}

	user := &model.Users{}

	err = db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	user.FirstName = req.Firstname
	user.LastName = req.Lastname
	user.Username = req.Username
	user.Email = req.Email
	user.Phone = req.Phone
	user.SetUpdateNow()

	_, err = db.NewUpdate().Model(user).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUserService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("users").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("user not found")
	}

	_, err = db.NewDelete().TableExpr("users").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
