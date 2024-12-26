package auth

import (
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := requests.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateToken(c, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, token)
}

