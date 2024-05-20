package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/2marks/go-exchangerate-api/config"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Cache struct {
	rd *redis.Client
}

func NewCache(r *redis.Client) *Cache {
	return &Cache{rd: r}
}

func (c *Cache) Get(key string, obj any) (any, error) {
	fmt.Printf("about to get cache value for: %s \n", key)

	value, err := c.rd.Get(ctx, key).Result()

	if err == redis.Nil {
		fmt.Printf("error from cache key: %s does not exist \n", key)
		return nil, fmt.Errorf("key does not exist")
	}

	if err != nil {
		fmt.Printf("error occured while getting cache value for: %s. err:%s \n", key, err.Error())
		return nil, err
	}

	fmt.Printf("retrieved cache value for: %s \n", key)

	if obj == nil {
		return value, err
	}

	if err := json.Unmarshal([]byte(value), obj); err != nil {
		fmt.Printf("error while unmarshalling value for key : %s", key)

		return nil, err
	}

	return obj, nil
}

func (c *Cache) Set(key string, value interface{}) error {
	valueToSet := value

	if _, ok := value.(string); !ok {
		bVal, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("error occured while marshalling value to json. key: %s", err.Error())

			return fmt.Errorf("error occured while marshalling value to json. key: %s", err.Error())
		}

		valueToSet = string(bVal)
	}

	fmt.Printf("about to set value for key: %s to cache \n", key)

	if err := c.rd.Set(ctx, key, valueToSet, 24*time.Hour).Err(); err != nil {
		fmt.Printf("error occured while setting value for : %s. err: %s \n", key, err.Error())

		return fmt.Errorf("error occured while saving value to cache. key: %s", err.Error())
	}

	return nil
}

func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Envs.RedisUrl,
		Password: config.Envs.RedisPassword,
		DB:       0,
	})
}
