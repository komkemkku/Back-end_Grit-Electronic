package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header id requird"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		token := parts[1]
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		type user struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
		}

		type admin struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		tt, _ := json.Marshal(claims["sub"])
		var usr *user
		_ = json.Unmarshal(tt, &usr)
		ctx.Set("user_id", usr.ID)

		tt, _ = json.Marshal(claims["sub"])
		var adm *admin
		_ = json.Unmarshal(tt, &adm)
		ctx.Set("admin_id", adm.ID)

		ctx.Next()
	}
}
