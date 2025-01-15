package notificationscontroller

import pushservice "github.com/petmeds24/backend/pkg/rest/src/services/notifications-service/push-service"

type NotificationsController struct {
	pushSvc *pushservice.PushNotificationService
}

func NewNotificationsController() *NotificationsController {
	pushSvc := pushservice.NewPushNotificationService()
	return &NotificationsController{
		pushSvc: pushSvc,
	}
}
