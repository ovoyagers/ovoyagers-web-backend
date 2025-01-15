package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type RedisConfig struct {
	redisClient *redis.Client
}

func NewRedisConfig() *RedisConfig {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	url := fmt.Sprintf("rediss://default:%s@%s:6379", cfg.UPSTASH_REDIS_REST_TOKEN, cfg.UPSTASH_REDIS_REST_URL)
	// url := "rediss://default:AY3FAAIjcDE0ODc4NjgxY2Q2ZDY0ZGI0ODEwY2IzNDQ1ZmVlMzllMXAxMA@funny-frog-36293.upstash.io:6379"
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}
	redisClient := redis.NewClient(opts)
	return &RedisConfig{
		redisClient: redisClient,
	}
}

func (rc *RedisConfig) GetRedisClient() *redis.Client {
	return rc.redisClient
}

func (rc *RedisConfig) CloseRedisClient() {
	rc.redisClient.Close()
}
