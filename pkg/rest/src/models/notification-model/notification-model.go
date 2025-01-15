package notificationmodel

import "firebase.google.com/go/messaging"

type SendNotificationRequest struct {
	Token            string                 `json:"token" validate:"required"`
	NotificationBody messaging.Notification `json:"notification" validate:"required"`
}

func (snr *SendNotificationRequest) Validate() error {
	return validateStruct(snr)
}
