package response

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Bank struct {
	ID      uuid.UUID
	Icon    uuid.NullUUID
	IconUrl null.String
	Code    string
}
