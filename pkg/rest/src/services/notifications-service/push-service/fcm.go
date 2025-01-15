package pushservice

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/petmeds24/backend/config"
	notificationmodel "github.com/petmeds24/backend/pkg/rest/src/models/notification-model"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

type PushNotificationService struct {
	fcmClient *messaging.Client
}

func NewPushNotificationService() *PushNotificationService {
	app := config.NewFcmConfig()
	fcmClient, err := app.FcmClient.Messaging(ctx)
	if err != nil {
		log.Errorf("error initializing firebase messagingclient: %v", err)
		return nil
	}
	return &PushNotificationService{
		fcmClient: fcmClient,
	}
}

func (pns *PushNotificationService) SendToSingleDevice(notification notificationmodel.SendNotificationRequest) error {
	resp, err := pns.fcmClient.Send(ctx, &messaging.Message{
		Notification: &messaging.Notification{
			Title:    notification.NotificationBody.Title,
			Body:     notification.NotificationBody.Body,
			ImageURL: notification.NotificationBody.ImageURL,
		},
		Token: notification.Token,
	})
	if err != nil {
		return err
	}
	log.Infof("Successfully sent message: %v", resp)
	return nil
}
