package util

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"

	"github.com/Halalins/backend/internal/common/constant"
)

func Password(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	re := regexp.MustCompile(`[a-zA-Z].*\d|\d.*[a-zA-Z]`)

	return len(value) >= 8 && re.MatchString(value)
}

func NotGreaterThanCurrentDate(fl validator.FieldLevel) bool {
	currentDate := time.Now().UTC()
	input, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	input = input.UTC()

	// extract date only
	input = GetDate(input)
	currentDate = GetDate(currentDate)

	return currentDate.After(input) || currentDate.Equal(input)
}

func TitleFormat(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	if strings.HasPrefix(value, "-") || strings.HasSuffix(value, "-") {
		return false
	}
	pattern := "^[a-zA-Z0-9 -]*[a-zA-Z][a-zA-Z0-9 -]*$"
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(value)
}

func NameFormat(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	pattern := `^[^\d\W_]+(?:[\w\s'.,()\[\]\-]*[^\W_)\].]+[)\.]?)?$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(value)
}

func AlphaNumSpace(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	pattern := "^[a-zA-Z0-9 ]+$"
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(value)
}

func AlphaNumSpaceHyphen(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	pattern := "^[a-zA-Z0-9 -]+$" // Updated pattern to allow hyphen '-'
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(value)
}

func AlphaSpace(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	pattern := "^[a-zA-Z ]+$"
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(value)
}

func MinIf(fl validator.FieldLevel) bool {
	tagParams := strings.Split(fl.Param(), " ")

	if len(tagParams) != 3 {
		return false
	}

	targetField := fl.Parent().FieldByName(tagParams[0])
	if !targetField.IsValid() {
		return false
	}

	targetValue := tagParams[1]

	minLen, err := strconv.Atoi(tagParams[2])
	if err != nil {
		return false
	}

	if targetField.String() == targetValue {
		switch k := fl.Field().Kind(); k {
		case reflect.Int64:
			return fl.Field().Int() >= int64(minLen)
		default:
			return fl.Field().Len() >= minLen
		}
	}

	return true
}

func MaxIf(fl validator.FieldLevel) bool {
	tagParams := strings.Split(fl.Param(), " ")

	if len(tagParams) != 3 {
		return false
	}

	targetField := fl.Parent().FieldByName(tagParams[0])
	if !targetField.IsValid() {
		return false
	}

	targetValue := tagParams[1]

	maxLen, err := strconv.Atoi(tagParams[2])
	if err != nil {
		return false
	}

	if targetField.String() == targetValue {
		switch k := fl.Field().Kind(); k {
		case reflect.Int64:
			return fl.Field().Int() <= int64(maxLen)
		default:
			return fl.Field().Len() <= maxLen
		}
	}

	return true
}

func UploadType(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	switch value {
	case constant.UploadTypeRichMedia, constant.UploadTypeEvidance, constant.UploadTypeProfilePicture:
		return true
	}
	return false
}

// AddValidation Register validator for struct
func AddValidation() {
	// register the validation in main function
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Password", Password)
		_ = v.RegisterValidation("Ngtcd", NotGreaterThanCurrentDate)
		_ = v.RegisterValidation("Title", TitleFormat)
		_ = v.RegisterValidation("Min_if", MinIf)
		_ = v.RegisterValidation("Max_if", MaxIf)
		_ = v.RegisterValidation("Alpha_space", AlphaSpace)
		_ = v.RegisterValidation("Alpha_num_space", AlphaNumSpace)
		_ = v.RegisterValidation("Alpha_num_space_hyphen", AlphaNumSpaceHyphen)
		_ = v.RegisterValidation("Name", NameFormat)
		_ = v.RegisterValidation("upload-type", UploadType)

		// register all null guregu value
		v.RegisterCustomTypeFunc(nullIntValidator, null.Int{})
		v.RegisterCustomTypeFunc(nullTimeValidator, null.Time{})
		v.RegisterCustomTypeFunc(nullFloatValidator, null.Float{})
		v.RegisterCustomTypeFunc(nullStringValidator, null.String{})
		v.RegisterCustomTypeFunc(nullBoolValidator, null.Bool{})
	}
}

func nullIntValidator(field reflect.Value) any {
	if valuer, ok := field.Interface().(null.Int); ok {
		if valuer.Valid {
			return valuer.Int64
		}
	}
	return nil
}

func nullTimeValidator(field reflect.Value) any {
	if valuer, ok := field.Interface().(null.Time); ok {
		if valuer.Valid {
			return valuer.Time
		}
	}
	return nil
}

func nullFloatValidator(field reflect.Value) any {
	if valuer, ok := field.Interface().(null.Float); ok {
		if valuer.Valid {
			return valuer.Float64
		}
	}
	return nil
}

func nullStringValidator(field reflect.Value) any {
	if valuer, ok := field.Interface().(null.String); ok {
		if valuer.Valid {
			return valuer.String
		}
	}
	return nil
}

func nullBoolValidator(field reflect.Value) any {
	if valuer, ok := field.Interface().(null.Bool); ok {
		if valuer.Valid {
			return valuer.Bool
		}
	}
	return nil
}
