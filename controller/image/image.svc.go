package image

import (
	"context"
	"errors"
	"fmt"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListImageBannerService(ctx context.Context, req requests.ImageRequest) ([]response.ImageBanner, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.ImageBanner{}

	query := db.NewSelect().
		TableExpr("images AS i").
		Column("i.id", "i.type", "i.banner", "i.created_at").
		OrderExpr("i.created_at DESC")

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Execute query
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func CreateImageBannerService(ctx context.Context, req requests.ImageCreateRequest) (*model.Images, error) {

	var imageCount int
	imageCount, err := db.NewSelect().
		Table("images").
		Where("type = ?", "Banner").
		Count(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to count images: %v", err)
	}

	// ถ้ามีครบ 4 รูปแล้ว ไม่ให้เพิ่มใหม่
	if imageCount >= 6 {
		return nil, errors.New("cannot add more than 6 images")
	}

	image := &model.Images{
		Type:   "Banner",
		Banner: req.Banner,
	}
	image.SetCreatedNow()

	_, err = db.NewInsert().Model(image).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to insert image: %v", err)
	}

	return image, nil
}

func DeleteImageService(ctx context.Context, ID int64) error {
	banner := &model.Images{}
	err := db.NewSelect().Model(banner).Where("id = ?", ID).Scan(ctx)
	if err != nil {
		return errors.New("product not found")
	}

	_, err = db.NewDelete().
		TableExpr("images").
		Where("id = ?", ID).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to delete image banner")
	}

	return nil
}
