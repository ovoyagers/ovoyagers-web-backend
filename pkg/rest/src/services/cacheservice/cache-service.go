package cacheservice

import (
	"encoding/json"
	"time"

	"github.com/petmeds24/backend/config"
)

type CacheService struct {
	globalCfg *config.GlobalConfig
}

func NewCacheService(globalCfg *config.GlobalConfig) *CacheService {
	return &CacheService{
		globalCfg: globalCfg,
	}
}

func (c *CacheService) SetRedisKey(key string, value interface{}, expiration int) error {
	redisClient := c.globalCfg.GetRedisClient()
	err := redisClient.Set(c.globalCfg.GetContext(), key, value, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheService) GetRedisKey(key string) (map[string]interface{}, error) {
	redisClient := c.globalCfg.GetRedisClient()
	value, err := redisClient.Get(c.globalCfg.GetContext(), key).Result()
	if err != nil {
		return nil, err
	}
	// convert value to map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
