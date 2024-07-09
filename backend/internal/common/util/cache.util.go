package util

import (
	"context"
	"errors"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

func GetCache[T any](ctx context.Context, rdb *redis.Client, key string) (*T, error) {
	var result T

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SetCache(ctx context.Context, rdb *redis.Client, key string, expDuration time.Duration, value any) error {
	cacheData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = rdb.Set(ctx, key, cacheData, expDuration).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetOrSetCache[T any](ctx context.Context, rdb *redis.Client, key string, expDuration time.Duration, query func() (*T, error)) (*T, error) {
	var result *T

	result, err := GetCache[T](ctx, rdb, key)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if result != nil {
		return result, nil
	}

	// failover get data from db
	resultDB, err := query()
	if err != nil {
		return nil, err
	}

	// set cache
	err = SetCache(ctx, rdb, key, expDuration, &resultDB)
	if err != nil {
		return nil, err
	}

	return resultDB, err
}

func DelCacheByKeyPattern(ctx context.Context, rdb *redis.Client, keys ...string) {
	errGroup, errGroupCtx := errgroup.WithContext(ctx)
	for _, key := range keys {
		currentKey := key
		errGroup.Go(func() error {
			iter := rdb.Scan(errGroupCtx, 0, currentKey, 0).Iterator()
			for iter.Next(errGroupCtx) {
				_, _ = rdb.Del(errGroupCtx, iter.Val()).Result()
			}
			return nil
		})
	}
	_ = errGroup.Wait()
}

func BatchSetCache(ctx context.Context, rdb *redis.Client, expiration time.Duration, dataMap map[string]any) error {
	for key, value := range dataMap {
		err := SetCache(ctx, rdb, key, expiration, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetCacheAndDelete[T any](ctx context.Context, rdb *redis.Client, key string) (*T, error) {
	var result T

	val, err := rdb.GetDel(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
