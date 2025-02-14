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

	if AdminID == 0 {
		response.Unauthorized(c, "Unauthorized access")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateProductService(c, req)
	if err != nil {
		logMessage := fmt.Sprintf("สร้างสินค้าล้มเหลว - สินค้า: %s, ข้อผิดพลาด: %s", req.Name, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "ADD_PRODUCT_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d เพิ่มสินค้า: %s", AdminID, req.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "ADD_PRODUCT_SUCCESS", logMessage)

	response.Success(c, data)
}


func DeleteProduct(c *gin.Context) {
	id := requests.ProductIdRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.BindUri(&id); err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ลบสินค้าล้มเหลว - ข้อมูลไม่ถูกต้อง: %s", AdminID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_PRODUCT_FAILED", logMessage)

		response.BadRequest(c, err.Error())
		return
	}

	// ดึงข้อมูลสินค้าก่อนลบ (เพื่อให้สามารถบันทึก log ได้ถูกต้อง)
	product, err := GetByIdProductService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ลบสินค้าล้มเหลว - ไม่พบสินค้า ID: %d, ข้อผิดพลาด: %s", AdminID, id.ID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_PRODUCT_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	// ลบสินค้า
	err = DeleteProductService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ลบสินค้าล้มเหลว - สินค้า: %s, ข้อผิดพลาด: %s", AdminID, product.Name, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_PRODUCT_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d ลบสินค้า: %s (ID: %d)", AdminID, product.Name, id.ID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_PRODUCT_SUCCESS", logMessage)

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
		// บันทึก log กรณีที่ URI ไม่ถูกต้อง
		logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสินค้าล้มเหลว - ข้อมูลไม่ถูกต้อง: %s", AdminID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_PRODUCT_FAILED", logMessage)

		response.BadRequest(c, err.Error())
		return
	}

	req := requests.ProductUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสินค้าล้มเหลว - JSON ไม่ถูกต้อง: %s", AdminID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_PRODUCT_FAILED", logMessage)

		response.BadRequest(c, err.Error())
		return
	}

	// ดึงข้อมูลสินค้าก่อนอัปเดต (เพื่อให้แสดงชื่อใน log)
	product, err := GetByIdProductService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสินค้าล้มเหลว - ไม่พบสินค้า ID: %d, ข้อผิดพลาด: %s", AdminID, id.ID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_PRODUCT_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	data, err := UpdateProductService(c, int(id.ID), req)
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสินค้าล้มเหลว - สินค้า: %s, ข้อผิดพลาด: %s", AdminID, product.Name, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_PRODUCT_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขสินค้า: %s (ID: %d)", AdminID, product.Name, id.ID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_PRODUCT_SUCCESS", logMessage)

	response.Success(c, data)
}
