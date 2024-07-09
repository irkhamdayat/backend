package errmapper

import (
	"errors"
	"strings"
	"sync"

	"github.com/Halalins/backend/internal/common/constant"
	"github.com/Halalins/backend/internal/common/customerror"
	"github.com/Halalins/backend/internal/common/util"
	"github.com/Halalins/backend/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

var (
	errMapper *ErrorMapper
	once      sync.Once
)

type ErrorMapper struct {
	mapDuplicateConstraint map[string]*customerror.CustomError
	i18nBundle             *i18n.Bundle
	i18nPrefix             string
}

func Initialize() *ErrorMapper {
	once.Do(func() {
		errMapper = new(ErrorMapper)
	})

	return errMapper
}

func (e *ErrorMapper) WithI18nBundle(i18nBundle *i18n.Bundle, keyPrefix string) *ErrorMapper {
	e.i18nBundle = i18nBundle
	e.i18nPrefix = keyPrefix
	return e
}

func (e *ErrorMapper) WithMapConstraintError(mapErr map[string]*customerror.CustomError) *ErrorMapper {
	e.mapDuplicateConstraint = mapErr
	return e
}

// HandleError handles the error and sends an appropriate JSON response
func HandleError(c *gin.Context, err error) {
	var (
		errResponse = errMapper.determineError(c, err)
	)

	c.JSON(errResponse.HTTPCode, response.ErrorResp{
		Error: *errResponse,
	})
}

func (e *ErrorMapper) processValidationErr(ginErr validator.ValidationErrors) *customerror.CustomError {
	var validationErrors []customerror.ValidationError

	for _, fieldError := range ginErr {
		validationError := customerror.ValidationError{
			Field:   fieldError.Field(),
			Tag:     fieldError.Tag(),
			Message: fieldError.Error(),
		}
		validationErrors = append(validationErrors, validationError)
	}

	errorResponse := *constant.ErrBindingError
	errorResponse.ValidationError = validationErrors

	return &errorResponse
}

func (e *ErrorMapper) determineError(c *gin.Context, err error) *customerror.CustomError {
	switch customErr := err.(type) {
	case *customerror.CustomError:
		return e.processErrorTranslation(c, customErr)
	case validator.ValidationErrors:
		return e.processValidationErr(customErr)
	default:
		return e.checkErrorConstraint(err)
	}
}

func (e *ErrorMapper) processErrorTranslation(c *gin.Context, customErr *customerror.CustomError) *customerror.CustomError {
	if e.i18nBundle == nil {
		return customErr
	}
	loc := i18n.NewLocalizer(e.i18nBundle, util.GetAcceptLangFromGinContext(c))
	text, _ := loc.Localize(&i18n.LocalizeConfig{
		MessageID:    e.i18nPrefix + "." + customErr.ErrorCode,
		TemplateData: customErr.Placeholder,
	})
	if text != "" {
		customErr.Message = text
	}
	return customErr
}

func (e *ErrorMapper) checkErrorConstraint(err error) *customerror.CustomError {
	// check error record not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return constant.ErrDataNotFound
	}

	// check error constraint
	for key, constErr := range e.mapDuplicateConstraint {
		if strings.Contains(err.Error(), key) {
			return constErr
		}
	}

	if strings.Contains(err.Error(), "UUID") {
		return constant.ErrInvalidUUID
	}

	return constant.ErrInternalServerError
}
