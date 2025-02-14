package categories

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateCategory(c *gin.Context) {
	req := requests.CategoryCreateRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateCategoryService(c, req)
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ล้มเหลวในการเพิ่มประเภทสินค้า: %s เนื่องจากข้อผิดพลาด: %v", AdminID, req.Name, err)
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "CREATE_CATEGORY_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d เพิ่มประเภทสินค้า: %s สำเร็จ", AdminID, req.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "CREATE_CATEGORY_SUCCESS", logMessage)

	response.Success(c, data)
}

func DeleteCategory(c *gin.Context) {
	id := requests.CategoryIdRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	category, err := GetByIdCategoryService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ล้มเหลวในการลบประเภทสินค้า ID: %d เนื่องจากไม่พบข้อมูล", AdminID, id.ID)
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_CATEGORY_FAILED", logMessage)

		response.InternalError(c, "Category not found")
		return
	}

	err = DeleteCategoryService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ล้มเหลวในการลบประเภทสินค้า: %s เนื่องจากข้อผิดพลาด: %v", AdminID, category.Name, err)
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_CATEGORY_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d ลบประเภทสินค้า: %s สำเร็จ", AdminID, category.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_CATEGORY_SUCCESS", logMessage)

	response.Success(c, "Delete successfully")
}

func GetCategoryByID(c *gin.Context) {
	id := requests.CategoryIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdCategoryService(c, int64(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CategoryList(c *gin.Context) {
	req := requests.CategoryRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListCategoryService(c.Request.Context(), req)
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

func UpdateCategory(c *gin.Context) {
	id := requests.CategoryIdRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.CategoryUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// ดึงข้อมูลประเภทสินค้าก่อนอัปเดต
	category, err := GetByIdCategoryService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ล้มเหลวในการแก้ไขประเภทสินค้า ID: %d เนื่องจากไม่พบข้อมูล", AdminID, id.ID)
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_CATEGORY_FAILED", logMessage)

		response.InternalError(c, "Category not found")
		return
	}

	// อัปเดตประเภทสินค้า
	data, err := UpdateCategoryService(c, int64(id.ID), req)
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ล้มเหลวในการแก้ไขประเภทสินค้า: %s เนื่องจากข้อผิดพลาด: %v", AdminID, category.Name, err)
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_CATEGORY_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d แก้ไขประเภทสินค้า: %s เป็น: %s", AdminID, category.Name, req.Name)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "UPDATE_CATEGORY_SUCCESS", logMessage)

	response.Success(c, data)
}
