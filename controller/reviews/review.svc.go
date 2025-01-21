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
    var reviews []response.ReviewResponses
    query := db.NewSelect().
    TableExpr("reviews r").
    Join("LEFT JOIN products p ON r.product_id = p.id").
    Join("LEFT JOIN users u ON r.user_id = u.id").
    Column("r.id", "r.rating", "r.description", "r.created_at", "r.updated_at", "u.email AS user_name", "p.name AS product_name")



    if req.Search != "" {
        query.Where("p.name LIKE ?", "%"+req.Search+"%")
    }

    err := query.Scan(ctx, &reviews)
    if err != nil {
        return nil, 0, err
    }

    total, err := query.Count(ctx)
    if err != nil {
        return nil, 0, err
    }

    return reviews, total, nil
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

	err = db.NewSelect().TableExpr("reviews AS r").
		Column("r.id", "r.text_review", "r.rating", "r.image_review", "r.created_at", "r.updated_at").
		ColumnExpr("p.id AS product__id").
		ColumnExpr("p.name AS product__name").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.username AS user__name").
		Join("LEFT JOIN products as p ON p.id = r.product_id").
		Join("LEFT JOIN users as u ON u.id = r.user_id").Where("r.id = ?", id).Scan(ctx, review)
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
		Rating:    int64(req.Rating),
		ProductID: int64(req.ProductID),
		UserID:    int64(req.UserID),
	}

	review.SetCreatedNow()
	review.SetUpdateNow()

	_, err = db.NewInsert().Model(review).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create review: %w", err)
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
