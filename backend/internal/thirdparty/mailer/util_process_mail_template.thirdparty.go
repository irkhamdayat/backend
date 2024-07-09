package mailer

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
)

func (rq *ThirdParty) processMailTemplate(mailTemplate constant.MailerTemplate, body interface{}) (*bytes.Buffer, error) {
	logger := logrus.WithField("template", mailTemplate)

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		logger.Errorf("failed get the current working directory: %v", err)
		return nil, err
	}

	// Construct the file path for the template
	templatePath := filepath.Join(wd, "internal/thirdparty/mailer/templates", string(mailTemplate))

	// Load and parse HTML template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		logger.Errorf("failed load and parse html template: %v", err)
		return nil, err
	}

	// Render template into a buffer
	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, body)
	if err != nil {
		logger.Errorf("failed render template into a buffer: %v", err)
		return nil, err
	}

	return &bodyContent, nil
}
