package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"strconv"
)

var ctx = context.Background()

type RedisManager struct {
	client *redis.Client
}

func (m RedisManager) set(key string, value string) {
	m.client.Set(ctx, key, value, 0)
}

func (m RedisManager) get(key string) string {
	val, _ := m.client.Get(ctx, key).Result()
	return val
}

func (m RedisManager) del(key string) {
	m.client.Del(ctx, key)
}

func newRedisManager(addr string, port int, password string, db int) *RedisManager {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + strconv.Itoa(port),
		Password: password,
		DB:       db,
	})

	if rdb == nil {
		return nil
	}

	return &RedisManager{client: rdb}
}