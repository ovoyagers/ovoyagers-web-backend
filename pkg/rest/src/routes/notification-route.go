package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	notificationscontroller "github.com/petmeds24/backend/pkg/rest/src/controllers/notifications-controller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type NotificationRoute struct {
	notificationController *notificationscontroller.NotificationsController
}

func NewNotificationRoute(globalCfg *config.GlobalConfig) *NotificationRoute {
	notificationController := notificationscontroller.NewNotificationsController()
	return &NotificationRoute{
		notificationController: notificationController,
	}
}

func (nr *NotificationRoute) SetupNotificationRoute(rg *gin.RouterGroup) {
	router := rg.Group("/notifications")
	router.Use(middlewares.DeserializeUser())

	router.POST("/send", nr.notificationController.SendNotification)
}
