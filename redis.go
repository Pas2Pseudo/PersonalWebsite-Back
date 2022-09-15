package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"strconv"
)

var ctx = context.Background()

type RedisManager struct {
	client *redis.Client
}

func (m RedisManager) set(key string, value string) {
	err := m.client.Set(ctx, key, value, 0).Err()

	if err != nil {
		log.Fatalf("Redis error: %v", err.Error())
		return
	}
}

func (m RedisManager) get(key string) string {
	val, err := m.client.Get(ctx, key).Result()

	if err != nil {
		log.Fatalf("Redis error: %v", err.Error())
		return ""
	}
	return val
}

func (m RedisManager) del(key string) {
	err := m.client.Del(ctx, key).Err()

	if err != nil {
		log.Fatalf("Redis error: %v", err.Error())
		return
	}
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