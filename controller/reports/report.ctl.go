package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/response"
)

func Dashboard(c *gin.Context) {

	data, err := GetDashboard(c)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}