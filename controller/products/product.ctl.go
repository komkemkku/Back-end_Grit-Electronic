package products

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateProduct(c *gin.Context) {
	req := requests.ProductCreateRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateProductService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// บันทึก Log
	logMessage := fmt.Sprintf("แอดมิน ID : %d เพิ่มสินค้า : %s", AdminID, req.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "ADD_PRODUCT", logMessage)

	response.Success(c, data)
}

func DeleteProduct(c *gin.Context) {
	id := requests.ProductIdRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// ดึงข้อมูลสินค้าก่อนลบ (เพื่อบันทึก log ที่ถูกต้อง)
	product, err := GetByIdProductService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	err = DeleteProductService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// บันทึก Log
	logMessage := fmt.Sprintf("แอดมิน ID: %d ลบสินค้า: %s", AdminID, product.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_PRODUCT", logMessage)

	response.Success(c, "delete successfully")
}


func GetProductByID(c *gin.Context) {
	id := requests.ProductIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdProductService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func ProductList(c *gin.Context) {
	req := requests.ProductRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListProductService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	paginate := model.Paginate{
		Page:  req.Page,
		Size:  req.Size,
		Total: int64(total),
	}

	response.SuccessWithPaginate(c, data, paginate)
}

func UpdateProduct(c *gin.Context) {
	id := requests.ProductIdRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.ProductUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// ดึงข้อมูลสินค้าก่อนอัปเดต (เพื่อให้แสดงชื่อใน log)
	product, err := GetByIdProductService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	data, err := UpdateProductService(c, int(id.ID), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// บันทึก Log
	logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสินค้า: %s", AdminID, product.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_PRODUCT", logMessage)

	response.Success(c, data)
}
