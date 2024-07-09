package cachekey

import (
	"fmt"
)

func RateLimitCacheKey(key, ip string) string {
	return fmt.Sprintf("%s-%s", key, ip)
}
