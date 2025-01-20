package image

import (
	"context"

	
	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
)

var db = configs.Database()

func CreateImageService(ctx context.Context, req requests.ImageCreateRequest) (*model.Images, error) {

	image := &model.Images{
		RefID:       req.RefID,
		Type:        req.Type,
		Description: req.Description,
	}
	image.SetCreatedNow()
	image.SetUpdateNow()

	_, err := db.NewInsert().Model(image).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return image, nil

}
