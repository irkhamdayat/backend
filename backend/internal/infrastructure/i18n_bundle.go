package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

func InitializeI18nBundle() *i18n.Bundle {
	i18nPath := "i18n"
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	err := filepath.Walk(i18nPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Failed to load i18n file template:", err)
		}
		if strings.HasSuffix(info.Name(), ".json") {
			bundle.MustLoadMessageFile(fmt.Sprintf("%s/%s", i18nPath, info.Name()))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return bundle
}
