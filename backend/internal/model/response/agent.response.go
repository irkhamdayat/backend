package response

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type GetAgentInfoResp struct {
	ID                uuid.UUID   `json:"id"`
	Code              string      `json:"code"`
	PhoneNumber       string      `json:"phoneNumber"`
	BirthDate         string      `json:"birthDate"`
	BirthPlace        string      `json:"birthPlace"`
	Address           string      `json:"address"`
	Location          string      `json:"location"`
	Photo             *Media      `json:"photo"`
	FirstName         string      `json:"firstName"`
	LastName          null.String `json:"lastName"`
	Email             string      `json:"email"`
	Status            string      `json:"status"`
	CodeReferral      string      `json:"codeReferral"`
	KtpDocument       *Media      `json:"ktpDocument"`
	KtpNumber         string      `json:"ktpNumber"`
	NpwpDocument      *Media      `json:"npwpDocument"`
	NpwpNumber        string      `json:"npwpNumber"`
	BankAccountNumber string      `json:"bankAccountNumber"`
	Bank              Bank        `json:"bank"`
	IsSubscribeNews   bool        `json:"isSubscribeNews"`
}

type PostVerifyEmailResp struct {
	Email string `json:"email"`
}

type PostVerifyPinResp struct {
	ID uuid.UUID `json:"id"`
}
