package storagecloud

import (
	"cloud.google.com/go/storage"
)

type ThirdParty struct {
	storageClient *storage.Client
}

func New() *ThirdParty {
	return new(ThirdParty)
}

func (tp *ThirdParty) WithStorageClient(client *storage.Client) *ThirdParty {
	tp.storageClient = client
	return tp
}
