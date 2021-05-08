package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	Redis *redis.Client
}

type Cache interface {
	Set(string, string)
	SetJSON(string, interface{})
	Get(string) string
}

func New(address string) Cache {
	return &cache{
		Redis: redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}

func (c *cache) Set(key, val string) {
	if c.Redis == nil {
		return
	}

	if err := c.Redis.Set(context.TODO(), key, val, 0).Err(); err != nil {
		fmt.Println(err)
	}
}

func (c *cache) SetJSON(key string, i interface{}) {
	if c.Redis == nil {
		return
	}

	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	c.Set(key, string(b))
}

func (c *cache) Get(key string) string {
	if c.Redis == nil {
		return ""
	}

	val, err := c.Redis.Get(context.TODO(), key).Result()
	if err != nil && err != redis.Nil {
		fmt.Println(err)
		return ""
	}

	return val
}
