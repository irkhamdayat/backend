package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/model/entity"
)

func (h *Handler) PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*entity.UserClaim); ok {
		return jwt.MapClaims{
			constant.IDKey:     v.ID,
			constant.EntityKey: v.Entity,
		}
	}

	return jwt.MapClaims{}
}
