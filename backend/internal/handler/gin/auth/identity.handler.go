package auth

import (
	"context"
	"github.com/Halalins/backend/internal/model/response"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
)

func (h *Handler) Identity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)

	IDstr, ok := claims[constant.IDKey].(string)
	if !ok {
		return nil
	}

	entityKey, ok := claims[constant.EntityKey].(string)
	if !ok {
		return nil
	}

	ID, err := uuid.Parse(IDstr)
	if err != nil {
		logrus.Error(err)
		return err
	}

	reqCtx := context.WithValue(c.Request.Context(), constant.IDKey, ID) //nolint
	reqCtx = context.WithValue(reqCtx, constant.EntityKey, entityKey)    //nolint
	c.Request = c.Request.WithContext(reqCtx)

	return &response.IDResp{
		ID: ID,
	}
}
