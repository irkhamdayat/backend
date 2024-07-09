package middleware

import (
	"github.com/Halalins/backend/config"
	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/handler/gin/auth"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
)

type AuthBuilder struct {
	authHandler *auth.Handler
	realm       string
}

func NewAuthBuilder() *AuthBuilder {
	return new(AuthBuilder)
}

func (m *AuthBuilder) WithAuthHandler(authHandler *auth.Handler) *AuthBuilder {
	m.authHandler = authHandler
	return m
}

func (m *AuthBuilder) WithRealm(realm string) *AuthBuilder {
	m.realm = realm
	return m
}

func (m *AuthBuilder) Build() (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            m.realm,
		SigningAlgorithm: "HS512",
		Key:              []byte(config.Env.JWT.UserSecret),
		Timeout:          config.Env.JWT.Timeout,
		MaxRefresh:       config.Env.JWT.MaxRefresh,
		IdentityKey:      constant.IDKey,
		PayloadFunc:      m.authHandler.PayloadFunc,
		IdentityHandler:  m.authHandler.Identity,
		Authenticator:    m.authHandler.Authenticator(m.realm),
		Authorizator:     m.authHandler.Authorizator,
		Unauthorized:     m.authHandler.Unauthorized,
		LoginResponse:    m.authHandler.LoginOrRefreshResponse,
		RefreshResponse:  m.authHandler.LoginOrRefreshResponse,
		TokenLookup:      "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
	})

	return jwtMiddleware, err
}
