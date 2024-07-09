package cachekey

import (
	"fmt"
	"github.com/google/uuid"
)

func AgentVerificationEmailTokenCacheKey(email string) string {
	return fmt.Sprintf("agent-verification-email:%s", email)
}

func AgentVerificationPinCacheKey(id uuid.UUID) string {
	return fmt.Sprintf("agent-verification-pin:%s", id)
}
