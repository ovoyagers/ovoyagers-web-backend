package config

import (
	"context"

	firebase "firebase.google.com/go"
	log "github.com/sirupsen/logrus"

	"google.golang.org/api/option"
)

var ctx context.Context = context.Background()

type FcmConfig struct {
	FcmClient *firebase.App
}

func NewFcmConfig() *FcmConfig {
	opt := option.WithCredentialsFile("config/fcm-testing.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Errorf("error initializing app: %v\n", err)
	}
	return &FcmConfig{
		FcmClient: app,
	}
}
