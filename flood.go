package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type floodControl struct {
	redisClient *redis.Client
	interval    time.Duration
	threshold   int
}

func NewFloodControl(redisAddr, redisPassword string, interval time.Duration, threshold int) *floodControl {
	return &floodControl{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       0,
		}),
		interval:  interval,
		threshold: threshold,
	}
}

func (fc *floodControl) Check(ctx context.Context, userID int64) (bool, error) {
	key := fmt.Sprintf("flood_control:%d", userID)

	// Получаем количество вызовов за последний интервал времени
	countCmd := fc.redisClient.ZCount(ctx, key, fmt.Sprintf("(%d", time.Now().Unix()-int64(fc.interval.Seconds())), "+inf")
	count, err := countCmd.Result()
	if err != nil {
		return false, err
	}

	// Проверяем, не превышает ли количество вызовов порог
	if count > int64(fc.threshold) {
		return false, nil
	}

	// Добавляем текущий вызов в хранилище с временной меткой
	now := time.Now().Unix()
	_, err = fc.redisClient.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	}).Result()
	if err != nil {
		return false, err
	}

	// Устанавливаем TTL на ключ
	_, err = fc.redisClient.Expire(ctx, key, fc.interval).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}
