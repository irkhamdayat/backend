package infrastructure

import (
	"context"
	"encoding/base64"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"github.com/Halalins/backend/config"
)

func InitializeStorageCloud() (client *storage.Client) {
	credentialDecoded, err := base64.StdEncoding.DecodeString(config.Env.GCP.Credential)
	if err != nil {
		logrus.Fatalf("Failed to decode gcp credential: %v", err)
	}

	client, err = storage.NewClient(context.Background(), option.WithCredentialsJSON(credentialDecoded))
	if err != nil {
		logrus.Fatalf("Failed to create client: %v", err)
	}

	return
}
