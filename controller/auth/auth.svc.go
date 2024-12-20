package auth

import (
	"context"
	"errors"

	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/utils"

	model "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
)

var db = config.Database()

func LoginService(ctx context.Context, req requests.LoginRequest) (*model.Users, error) {
	ex, err := db.NewSelect().TableExpr("users").Where("email = ?", req.Email).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("email or password not found")
	}

	user := &model.Users{}

	err = db.NewSelect().Model(user).Where("email =?", req.Email).Scan(ctx)
	if err != nil {
		return nil, err
	}

	bool := utils.CheckPasswordHash(req.Password, user.Password)

	if !bool {
		return nil, errors.New("email or password not found")
	}

	return user, nil
}
