package reviews

import (
	"context"
	"errors"

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
		ColumnExpr("r.description AS description").
		ColumnExpr("r.created_at AS created_at").
		ColumnExpr("r.updated_at AS updated_at").
		Join("LEFT JOIN products AS p ON p.id = r.product_id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id").
		GroupExpr("r.id, u.id, p.id")

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

	err = db.NewSelect().
		TableExpr("reviews AS r").
		ColumnExpr("r.id AS id").
		ColumnExpr("u.username AS \"user\"").
		ColumnExpr("p.name AS product").
		ColumnExpr("r.rating AS rating").
		ColumnExpr("r.description AS description").
		ColumnExpr("r.created_at AS created_at").
		ColumnExpr("r.updated_at AS updated_at").
		Join("LEFT JOIN products AS p ON p.id = r.product_id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id").
		GroupExpr("r.id, u.id, p.id").
		Where("r.id = ?", id).Scan(ctx, review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func CreateReviewService(ctx context.Context, req requests.ReviewCreateRequest) (*model.Reviews, error) {
	// ตรวจสอบค่าคะแนนรีวิว
	if req.Rating <= 0 || req.Rating > 5 {
		return nil, errors.New("คะแนนรีวิวต้องอยู่ระหว่าง 1 ถึง 5")
	}

	// ตรวจสอบว่าสินค้านี้มีอยู่ในคำสั่งซื้อของผู้ใช้หรือไม่
	orderExists, err := db.NewSelect().
		TableExpr("order_details AS od").
		Join("JOIN orders AS o ON o.id = od.order_id").
		Join("JOIN products AS p ON p.name = od.product_name").
		Where("o.user_id = ? AND p.id = ?", req.UserID, req.ProductID).
		Exists(ctx)

	if err != nil {
		return nil, errors.New("failed to check order details")
	}

	if !orderExists {
		return nil, errors.New("you can only review products from your order")
	}

	// ตรวจสอบว่าผู้ใช้เคยรีวิวสินค้านี้ไปแล้วหรือไม่
	reviewExists, err := db.NewSelect().
		Table("reviews").
		Where("user_id = ? AND product_id = ?", req.UserID, req.ProductID).
		Exists(ctx)

	if err != nil {
		return nil, errors.New("failed to check review status")
	}

	if reviewExists {
		return nil, errors.New("you have already reviewed this product")
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

	// บันทึกข้อมูลรีวิวลงฐานข้อมูล
	_, err = db.NewInsert().Model(review).Exec(ctx)
	if err != nil {
		return nil, errors.New("failed to insert review")
	}

	return review, nil
}

func UpdateReviewService(ctx context.Context, id int, req requests.ReviewUpdateRequest) (*model.Reviews, error) {
	// ตรวจสอบว่ามีรีวิวอยู่หรือไม่
	ex, err := db.NewSelect().Table("reviews").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("review not found")
	}

	// ดึงข้อมูลรีวิวปัจจุบัน
	review := &model.Reviews{}
	err = db.NewSelect().Model(review).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// อัปเดตข้อมูลรีวิว
	review.ProductID = req.ProductID
	review.UserID = req.UserID
	review.Description = req.Description
	review.Rating = req.Rating
	review.SetUpdateNow()

	// บันทึกการอัปเดต
	_, err = db.NewUpdate().Model(review).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func DeleteReviewService(ctx context.Context, id int64) error {
	// ตรวจสอบว่ารีวิวมีอยู่ในระบบหรือไม่
	exists, err := db.NewSelect().TableExpr("reviews").Where("id=?", id).Exists(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("ไม่พบรีวิวในระบบ")
	}

	// ลบรูปภาพที่เกี่ยวข้องกับรีวิว
	_, err = db.NewDelete().
		TableExpr("images").
		Where("ref_id = ? AND type = 'review_image'", id).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to delete")
	}

	// ลบข้อมูลรีวิว
	_, err = db.NewDelete().
		TableExpr("reviews").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return errors.New("failed to delete review")
	}

	return nil
}
