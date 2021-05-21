package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	Redis   *redis.Client
	Enabled bool
}

type Cache interface {
	Set(string, string)
	SetJSON(string, interface{})
	Get(string) string
}

func New(address string) Cache {
	r := redis.NewClient(&redis.Options{
		Addr: address,
	})

	if r.Ping(context.TODO()).Err() != nil {
		return &cache{
			Enabled: false,
		}
	}
	return &cache{
		Redis:   r,
		Enabled: true,
	}
}

func NewDisabled() Cache {
	return &cache{
		Enabled: false,
	}
}

func (c *cache) Set(key, val string) {
	if !c.Enabled {
		return
	}

	if err := c.Redis.Set(context.TODO(), key, val, 0).Err(); err != nil {
		fmt.Println(err)
	}
}

func (c *cache) SetJSON(key string, i interface{}) {
	if !c.Enabled {
		return
	}

	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	c.Set(key, string(b))
}

func (c *cache) Get(key string) string {
	if !c.Enabled {
		return ""
	}

	val, err := c.Redis.Get(context.TODO(), key).Result()
	if err != nil && err != redis.Nil {
		fmt.Println(err)
		return ""
	}

	return val
}
