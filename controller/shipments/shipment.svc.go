package shipments

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	configs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

var db = configs.Database()

func ListShipmentService(ctx context.Context, req requests.ShipmentRequest) ([]response.ShipmentResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.ShipmentResponses{}

	// สร้าง query
	query := db.NewSelect().
    TableExpr("shipments AS s").
    Column("s.id", "s.firstname", "s.lastname", "s.address", "s.zip_code", "s.sub_district", "s.district", "s.province", "s.status", "s.created_at", "s.updated_at")

	if req.Search != "" {
		query.Where("s.zip_code LIKE ?", "%"+req.Search+"%")
	}

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

func GetByIdShipmentService(ctx context.Context, id int64) (*response.ShipmentResponses, error) {
	ex, err := db.NewSelect().TableExpr("shipments").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("not found")
	}
	shipment := &response.ShipmentResponses{}

	err = db.NewSelect().
		TableExpr("shipments AS s").
		Column("s.id", "s.firstname", "s.lastname", "s.address", "s.zip_code", "s.sub_district", "s.district", "s.province", "s.status", "s.created_at", "s.updated_at").
		Where("s.id = ?", id).Scan(ctx, shipment)

	if err != nil {
		return nil, err
	}
	return shipment, nil
}

func CreateShipmentService(ctx context.Context, req requests.ShipmentCreateRequest) (*model.Shipments, error) {

	statusInt, err := strconv.Atoi(req.Status)
	if err != nil {
        return nil, fmt.Errorf("invalid status value: %v", err) // จัดการข้อผิดพลาด
    }
	zipcodeInt, err := strconv.Atoi(req.ZipCode)
	if err != nil {
        return nil, fmt.Errorf("invalid zipcode value: %v", err) // จัดการข้อผิดพลาด
    }

	// เพิ่ม
	shipment := &model.Shipments{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address:     req.Address,
		ZipCode:     zipcodeInt,
		SubDistrict: req.SubDistrict,
		District:    req.District,
		Province:    req.Province,
		Status:      statusInt,
	}
	shipment.SetCreatedNow()
	shipment.SetUpdateNow()

	_, err = db.NewInsert().Model(shipment).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return shipment, nil

}

func UpdateShipmentService(ctx context.Context, id int64, req requests.ShipmentUpdateRequest) (*model.Shipments, error) {
	ex, err := db.NewSelect().TableExpr("shipments").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("not found")
	}

	shipment := &model.Shipments{}

	statusInt, err := strconv.Atoi(req.Status)
	if err!= nil {
        return nil, fmt.Errorf("invalid status value: %v", err)
    }
	zipcodeInt, err := strconv.Atoi(req.ZipCode)
	if err!= nil {
        return nil, fmt.Errorf("invalid status value: %v", err)
    }

	err = db.NewSelect().Model(shipment).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	shipment.Firstname = req.Firstname
	shipment.Lastname = req.Lastname
	shipment.Address = req.Address
	shipment.ZipCode = zipcodeInt
	shipment.SubDistrict = req.SubDistrict
	shipment.District = req.District
	shipment.Province = req.Province
	shipment.Status = statusInt
	shipment.SetUpdateNow()

	_, err = db.NewUpdate().Model(shipment).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func DeleteShipmentService(ctx context.Context, id int64) error {
	ex, err := db.NewSelect().TableExpr("shipments").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("not found")
	}

	_, err = db.NewDelete().TableExpr("shipments").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
