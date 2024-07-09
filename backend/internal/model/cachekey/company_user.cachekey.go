package cachekey

import "fmt"

func AdminForgotPasswordTokenCacheKey(email string) string {
	return fmt.Sprintf("admin-forgot-password:%s", email)
}
