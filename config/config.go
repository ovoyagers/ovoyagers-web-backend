package config

import (
	"context"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type GlobalConfig struct {
	ctx         context.Context
	config      *Config
	redisclient *redis.Client
	// Define other configuration fields as needed
}

func NewGlobalConfig(ctx context.Context) *GlobalConfig {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	redisclient := NewRedisConfig().GetRedisClient()
	return &GlobalConfig{
		ctx:         ctx,
		config:      config,
		redisclient: redisclient,
	}
}

func (globalCfg *GlobalConfig) GetConfig() *Config {
	return globalCfg.config
}

func (globalCfg *GlobalConfig) GetRedisClient() *redis.Client {
	return globalCfg.redisclient
}

func (globalCfg *GlobalConfig) GetContext() context.Context {
	return globalCfg.ctx
}
