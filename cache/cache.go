package cache

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/tushargupta98/api-in-go/config"
	"github.com/tushargupta98/api-in-go/logger"
)

var (
	redisClient *redis.Client
	expiration  time.Duration
)

func init() {
	// Initialize the Redis client once when the package is imported
	cfg := config.GetConfig().Cache

	db, err := stringToInt(cfg.DB)
	if err != nil {
		logger.Error("error parsing cache db: %w", err)
	}

	expiration = cfg.Expiration

	options := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           db,
		PoolSize:     cfg.PoolSize,
		MaxConnAge:   0,
		PoolTimeout:  cfg.PoolTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		MaxRetries:   cfg.MaxRetries,
	}

	if cfg.TLSEnabled {
		options.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	redisClient = redis.NewClient(options)

	// Test the connection to Redis
	if err := redisClient.Ping().Err(); err != nil {
		logger.Error("Error connecting to Redis: ", err)
	}
}

type RedisClientInterface interface {
	Close() error
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}

type RedisClient struct{}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (r *RedisClient) Close() error {
	return redisClient.Close()
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisClient) Set(key string, value string) error {
	return redisClient.Set(key, value, expiration).Err()
}

func (r *RedisClient) Delete(key string) error {
	return redisClient.Del(key).Err()
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
