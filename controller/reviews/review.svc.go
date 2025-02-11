package reviews

import (
	"context"
	"errors"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/image"
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
		ColumnExpr("COALESCE(json_agg(json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description)) FILTER (WHERE i.id IS NOT NULL), '[]') AS image").
		Join("LEFT JOIN products AS p ON p.id = r.product_id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id").
		Join("LEFT JOIN images AS i ON i.ref_id = r.id AND i.type = 'review_image'").
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

	// for i := range resp {
	// 	resp[i].ImageReview = []string{}
	// }

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

	err = db.NewSelect().TableExpr("reviews AS r").
		ColumnExpr("r.id AS id").
		ColumnExpr("u.username AS \"user\"").
		ColumnExpr("p.name AS product").
		ColumnExpr("r.rating AS rating").
		ColumnExpr("r.description AS description").
		ColumnExpr("r.created_at AS created_at").
		ColumnExpr("r.updated_at AS updated_at").
		ColumnExpr("COALESCE(json_agg(json_build_object('id', i.id, 'ref_id', i.ref_id, 'type', i.type, 'description', i.description)) FILTER (WHERE i.id IS NOT NULL), '[]') AS image").
		Join("LEFT JOIN products AS p ON p.id = r.product_id").
		Join("LEFT JOIN users AS u ON u.id = r.user_id").
		Join("LEFT JOIN images AS i ON i.ref_id = r.id AND i.type = 'review_image'").
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

	// ตรวจสอบ `ProductID`
	if req.ProductID <= 0 {
		return nil, errors.New("product not found")
	}

	// ตรวจสอบว่าสินค้ามีอยู่ในระบบหรือไม่
	exists, err := db.NewSelect().
		Table("products").
		Where("id = ?", req.ProductID).
		Exists(ctx)
	if err != nil {
		return nil, errors.New("error")
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

	// บันทึกข้อมูลรีวิวลงฐานข้อมูล
	_, err = db.NewInsert().Model(review).Exec(ctx)
	if err != nil {
		return nil, errors.New("failed to insert")
	}

	// เพิ่มรูปภาพสำหรับรีวิว
	for _, image := range req.ImageReview {
		img := &model.Images{
			RefID:       review.ID,
			Type:        "review_image",
			Description: image,
		}

		_, err := db.NewInsert().Model(img).Exec(ctx)
		if err != nil {
			return nil, errors.New("failed to insert image")
		}
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

	// ตรวจสอบว่ารูปภาพที่เกี่ยวข้องมีอยู่หรือไม่
	exists, err := db.NewSelect().
		TableExpr("images").
		Where("ref_id = ? AND type = 'review_image'", review.ID).
		Exists(ctx)
	if err != nil {
		return nil, err
	}

	if exists {
		// ถ้ามีรูปภาพอยู่แล้ว ให้ลบภาพเก่าออกก่อน
		_, err = db.NewDelete().
			TableExpr("images").
			Where("ref_id = ? AND type = 'review_image'", review.ID).
			Exec(ctx)
		if err != nil {
			return nil, errors.New("failed to delete old review images")
		}

		// เพิ่มรูปภาพใหม่
		for _, image := range req.ImageReview {
			img := &model.Images{
				RefID:       review.ID,
				Type:        "review_image",
				Description: image,
			}

			_, err := db.NewInsert().Model(img).Exec(ctx)
			if err != nil {
				return nil, errors.New("failed to insert image")
			}
		}
	} else {
		// ถ้าไม่มีรูปภาพ ให้สร้างใหม่
		for _, images := range req.ImageReview {
			img := requests.ImageCreateRequest{
				RefID:       review.ID,
				Type:        "review_image",
				Description: images,
			}

			_, err = image.CreateImageService(ctx, img)
			if err != nil {
				return nil, errors.New("failed to create review slip")
			}
		}
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
