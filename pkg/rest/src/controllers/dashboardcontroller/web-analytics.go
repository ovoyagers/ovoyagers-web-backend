package dashboardcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// GetWebAnalytics handles the HTTP request for fetching web analytics data.
//
//	@Summary		Retrieve web analytics data
//	@Description	Fetches web analytics data from the service and returns it in the response.
//	@Tags			analytics
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		404	{object}	models.Error	"Status Not Found"
//	@Router			/dashboard/web-analytics [get]
//	@Security		BearerAuth
func (dc *DashboardController) GetWebAnalytics(c *gin.Context) {
	data, err := dc.dashboardSvc.GetWebAnalytics()
	if err != nil {
		utils.HTTPErrorHandler(c, err, 404, "Status Not Found")
		return
	}
	utils.HTTPResponseHandler(c, data, 200, "Success")
}
