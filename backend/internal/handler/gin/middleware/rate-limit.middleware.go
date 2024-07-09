package middleware

import (
	"fmt"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/cachekey"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/Halalins/backend/internal/common/errmapper"
	"github.com/gin-gonic/gin"
)

type RateLimitMiddleware struct {
	rdb *redis.Client
}

func NewRateLimitMiddleware() *RateLimitMiddleware {
	return new(RateLimitMiddleware)
}

func (m *RateLimitMiddleware) WithRedisDatabase(rdb *redis.Client) *RateLimitMiddleware {
	m.rdb = rdb
	return m
}

func (m *RateLimitMiddleware) Rate(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx = c.Request.Context()
			err error
		)

		cacheKey := cachekey.RateLimitCacheKey(key, c.ClientIP())
		limiter := redis_rate.NewLimiter(m.rdb)

		rateLimit, err := limiter.Allow(ctx, cacheKey, redis_rate.Limit{
			Rate:   1,
			Period: time.Minute * 5,
			Burst:  3,
		})
		if err != nil {
			err = constant.ErrToManyRequest
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		fmt.Println(rateLimit.Allowed, rateLimit.Remaining, rateLimit.ResetAfter, rateLimit.RetryAfter)

		if rateLimit.Allowed <= 0 {
			err = constant.ErrToManyRequest
			errmapper.HandleError(c, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
