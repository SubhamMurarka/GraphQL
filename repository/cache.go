package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/subhammurarka/GraphQL/graph/model"
)

type cache struct {
	redisClient *redis.Client
}

type CacheInterface interface {
	GetMaterialAndSupplierCache(materialType string, price float64, location string) (model.UserResponse, error)
	CacheResponse(req *model.UserRequest, res *model.UserResponse) error
}

func NewCache(redisClient *redis.Client) CacheInterface {
	return &cache{
		redisClient: redisClient,
	}
}

func (c *cache) GetMaterialAndSupplierCache(materialType string, price float64, location string) (model.UserResponse, error) {
	redisKey := fmt.Sprintf("%s:%f:%s", materialType, price, location)

	jsonData, err := c.redisClient.Get(context.TODO(), redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("key not found : %v", err)
			return model.UserResponse{}, err
		}

		log.Printf("Redis error: %v", err)
		return model.UserResponse{}, err
	}

	var response model.UserResponse

	err = json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		log.Printf("Failed to unmarshal cached data: %v", err)
		return model.UserResponse{}, err
	}

	return response, nil
}

func (c *cache) CacheResponse(req *model.UserRequest, resp *model.UserResponse) error {
	redisKey := fmt.Sprintf("%s:%f:%s", req.MaterialType, req.Price, req.Locality)

	jsonData, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	err = c.redisClient.Set(context.TODO(), redisKey, jsonData, 40*time.Second).Err()
	if err != nil {
		log.Printf("error caching response : %v", err)
		return err
	}

	return nil
}
