package image

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/requests"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func CreateImageBanner(c *gin.Context) {
	AdminID := c.GetInt("admin_id")

	var req requests.ImageCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// ตรวจสอบว่ามีแบนเนอร์ครบ 4 รูปแล้วหรือยัง
	data, err := CreateImageBannerService(c.Request.Context(), req)
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d เพิ่มแบนเนอร์ล้มเหลว: %s", AdminID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "ADD_BANNER_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d เพิ่มแบนเนอร์สำเร็จ ID: %d", AdminID, data.ID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "ADD_BANNER_SUCCESS", logMessage)

	response.Success(c, data)
}


func DeleteImageBanner(c *gin.Context) {
	id := requests.ImageIdRequest{}
	AdminID := c.GetInt("admin_id")

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := DeleteImageService(c, int64(id.ID))
	if err != nil {
		logMessage := fmt.Sprintf("แอดมิน ID: %d ล้มเหลวในการลบแบนเนอร์ ID: %d - ข้อผิดพลาด: %s", AdminID, id.ID, err.Error())
		_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_BANNER_FAILED", logMessage)

		response.InternalError(c, err.Error())
		return
	}

	logMessage := fmt.Sprintf("แอดมิน ID: %d ลบแบนเนอร์สำเร็จ ID: %d", AdminID, id.ID)
	_ = adminlogs.CreateAdminLog(c.Request.Context(), AdminID, "DELETE_BANNER_SUCCESS", logMessage)

	response.Success(c, "Delete successfully")
}

func BannerList(c *gin.Context) {
	req := requests.ImageRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListImageBannerService(c.Request.Context(), req)
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
