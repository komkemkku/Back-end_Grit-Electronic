package jwt

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
)

func VerifyToken(raw string) (map[string]any, error) {
	godotenv.Load()
	token, err := jwt.Parse(raw, func(token *jwt.Token) (
		interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token singing method")
		}
		secret := []byte(os.Getenv("TOKEN_SECRET_USER"))
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}

func GenerateToken(ctx context.Context, user *model.Users) (string, error) {
	godotenv.Load()
	tokenDurationStr := os.Getenv("TOKEN_DURATION_USER")
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": jwt.MapClaims{
			"id":     user.Id,
			"username":   user.Username,
			"password": user.Password,
			"email":  user.Email,
			"phone": user.Phone,
			"bank_number": user.Bank_number,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(tokenDuration).Unix(),
	})

	secret := []byte(os.Getenv("TOKEN_SECRET_USER"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}
