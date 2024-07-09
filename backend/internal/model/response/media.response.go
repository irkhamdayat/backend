package response

import "github.com/google/uuid"

type Media struct {
	ID  uuid.UUID `json:"id"`
	Url string    `json:"url"`
}
