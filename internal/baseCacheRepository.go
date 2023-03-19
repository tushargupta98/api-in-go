package internal

import (
	"encoding/json"

	"github.com/tushargupta98/api-in-go/cache"
)

type BaseCacheRepository struct {
	cache cache.RedisClient
}

func NewBaseCacheRepository(cache cache.RedisClient) *BaseCacheRepository {
	return &BaseCacheRepository{cache}
}

func (r *BaseCacheRepository) Get(key string, dest interface{}) error {
	cacheValue, err := r.cache.Get(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(cacheValue), dest)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseCacheRepository) Set(key string, value interface{}, expiration int) error {
	cacheValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.cache.Set(key, string(cacheValue))
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseCacheRepository) Delete(key string) error {
	err := r.cache.Delete(key)
	if err != nil {
		return err
	}
	return nil
}
