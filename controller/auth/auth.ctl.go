package auth

import (
	"fmt"
	"net/http"

	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils/jwt"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	req := requests.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		// บันทึก log กรณีข้อมูลไม่ถูกต้อง
		logMessage := fmt.Sprintf("เข้าสู่ระบบล้มเหลว - ข้อมูลไม่ถูกต้อง: %s", err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), 0, "LOGIN_FAILED", logMessage)

		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginUserService(c, req)
	if err != nil {
		// บันทึก log กรณีเข้าสู่ระบบไม่สำเร็จ
		logMessage := fmt.Sprintf("เข้าสู่ระบบล้มเหลว - อีเมล: %s, ข้อผิดพลาด: %s", req.Email, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), 0, "LOGIN_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenUser(c, data)
	if err != nil {
		// บันทึก log กรณีสร้าง token ไม่สำเร็จ
		logMessage := fmt.Sprintf("สร้าง Token ล้มเหลว - อีเมล: %s, ข้อผิดพลาด: %s", req.Email, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), 0, "LOGIN_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	// บันทึก log กรณีเข้าสู่ระบบสำเร็จ
	logMessage := fmt.Sprintf("เข้าสู่ระบบสำเร็จ - ผู้ใช้ ID: %d, อีเมล: %s", data.ID, req.Email)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), data.ID, "LOGIN_SUCCESS", logMessage)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}

func LoginAdmin(c *gin.Context) {
	req := requests.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		// บันทึก log กรณีข้อมูลไม่ถูกต้อง
		logMessage := fmt.Sprintf("แอดมินเข้าสู่ระบบล้มเหลว - ข้อมูลไม่ถูกต้อง: %s", err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), 0, "ADMIN_LOGIN_FAILED", logMessage)

		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginAdminService(c, req)
	if err != nil {
		// บันทึก log กรณีเข้าสู่ระบบไม่สำเร็จ
		logMessage := fmt.Sprintf("แอดมินเข้าสู่ระบบล้มเหลว - อีเมล: %s, ข้อผิดพลาด: %s", req.Email, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), 0, "ADMIN_LOGIN_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenAdmin(c, data)
	if err != nil {
		// บันทึก log กรณีสร้าง token ไม่สำเร็จ
		logMessage := fmt.Sprintf("สร้าง Token ล้มเหลว - แอดมินอีเมล: %s, ข้อผิดพลาด: %s", req.Email, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), 0, "ADMIN_LOGIN_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	// บันทึก log กรณีเข้าสู่ระบบสำเร็จ
	logMessage := fmt.Sprintf("แอดมินเข้าสู่ระบบสำเร็จ - แอดมิน ID: %d, อีเมล: %s", data.ID, req.Email)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), data.ID, "ADMIN_LOGIN_SUCCESS", logMessage)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}
