package response

import "github.com/google/uuid"

type MessageResponse struct {
	Message string `json:"message"`
}

type CommonWithDataResp struct {
	Data any `json:"data"`
}

type LoginResponse struct {
	ExpiredAt   string `json:"expiredAt"`
	AccessToken string `json:"accessToken"`
}

type YearMonthResp struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type IDResp struct {
	ID uuid.UUID `json:"id"`
}

type EmailResp struct {
	Email string `json:"email"`
}
