package reviews

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

func ListReviewService(ctx context.Context, req requests.ReviewRequest) ([]response.ReviewResponses, int, error) {
	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.ReviewResponses{}

	query := db.NewSelect().
		TableExpr("reviews AS r").
		ColumnExpr("r.id AS id").
		ColumnExpr("u.username AS \"user\"").
		ColumnExpr("p.name AS product").
		ColumnExpr("r.rating AS rating").
		ColumnExpr("r.description AS text_review").
		ColumnExpr("r.description AS description").
		ColumnExpr("r.created_at AS created_at").
		ColumnExpr("r.updated_at AS updated_at").
		Join("LEFT JOIN products AS p ON p.id = r.product_id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id")

	if req.Search != "" {
		query.Where("p.name ILIKE ?", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	for i := range resp {
		resp[i].ImageReview = []string{}
	}

	return resp, total, nil
}

func GetByIdReviewService(ctx context.Context, id int64) (*response.ReviewResponses, error) {
	ex, err := db.NewSelect().TableExpr("reviews").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("review not found")
	}
	review := &response.ReviewResponses{}

	err = db.NewSelect().	TableExpr("reviews AS r").
	ColumnExpr("r.id AS id").
	ColumnExpr("u.username AS \"user\"").
	ColumnExpr("p.name AS product").
	ColumnExpr("r.rating AS rating").
	ColumnExpr("r.description AS description").
	ColumnExpr("r.created_at AS created_at").
	ColumnExpr("r.updated_at AS updated_at").
	Join("LEFT JOIN products AS p ON p.id = r.product_id").
	Join("LEFT JOIN users AS u ON u.id = r.user_id").Where("r.id = ?", id).Scan(ctx, review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func CreateReviewService(ctx context.Context, req requests.ReviewCreateRequest) (*model.Reviews, error) {
	// ตรวจสอบค่าที่ส่งมา
	if req.Rating <= 0 || req.Rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	if req.ProductID <= 0 {
		return nil, errors.New("invalid product ID")
	}

	// ตรวจสอบว่าสินค้ามีอยู่ในฐานข้อมูล
	exists, err := db.NewSelect().
		Table("products").
		Where("id = ?", req.ProductID).
		Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check product existence: %w", err)
	}
	if !exists {
		return nil, errors.New("product not found")
	}

	// เพิ่มรีวิวใหม่
	review := &model.Reviews{
		Rating:      req.Rating,
		ProductID:   req.ProductID,
		UserID:      req.UserID,
		Description: req.Description,
	}

	review.SetCreatedNow()
	review.SetUpdateNow()

	_, err = db.NewInsert().Model(review).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create review: %w", err)
	}

	img := requests.ImageCreateRequest{
		RefID:       review.ID,
		Type:        "review_image",
		Description: req.ImageReview,
	}

	_, err = image.CreateImageService(ctx, img)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func UpdateReviewService(ctx context.Context, id int64, req requests.ReviewUpdateRequest) (*model.Reviews, error) {
	ex, err := db.NewSelect().TableExpr("reviews").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("review not found")
	}

	review := &model.Reviews{}

	err = db.NewSelect().Model(review).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	review.ProductID = req.ProductID
	review.UserID = req.UserID
	review.Description = req.Description
	review.Rating = req.Rating
	review.SetUpdateNow()

	_, err = db.NewUpdate().Model(review).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func DeleteReviewService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("reviews").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("review not found")
	}

	_, err = db.NewDelete().TableExpr("reviews").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
