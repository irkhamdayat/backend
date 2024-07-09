package util

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	mrand "math/rand"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateRandomString(length int, base string) string {
	token := make([]byte, length)
	alphabetLen := big.NewInt(int64(len(base)))

	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, alphabetLen)
		token[i] = base[index.Int64()]
	}

	return string(token)
}

// TruncateString truncates a string to a specified maximum length and appending appendingStr
func TruncateString(input string, maxLength int, appendingStr string) string {
	// Check if the input string length is already within the limit
	if utf8.RuneCountInString(input) <= maxLength {
		return input
	}

	appendingStrLenght := utf8.RuneCountInString(appendingStr)

	// If not, truncate the string and append "..."
	runes := []rune(input)
	truncatedRunes := runes[:maxLength-appendingStrLenght]
	return string(truncatedRunes) + appendingStr
}

func TranslateWithPlaceholder(ctx context.Context, bundle *i18n.Bundle, messageID string, placeholder map[string]any) (string, error) {
	if bundle == nil {
		return "", errors.New("nil bundle")
	}
	loc := i18n.NewLocalizer(bundle, GetAcceptLanguageFromContext(ctx))
	text, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: placeholder,
	})
	return text, err
}

func TranslateWithLangAndPlaceholder(lang string, bundle *i18n.Bundle, messageID string, placeholder map[string]any) (string, error) {
	if bundle == nil {
		return "", errors.New("nil bundle")
	}
	loc := i18n.NewLocalizer(bundle, lang)
	text, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: placeholder,
	})
	return text, err
}

func GenerateRandomNumber(length int) string {
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	randomNumber := ""
	for i := 0; i < length; i++ {
		randomDigit := r.Intn(10) // generates a random digit between 0 and 9
		randomNumber += strconv.Itoa(randomDigit)
	}
	return randomNumber
}
