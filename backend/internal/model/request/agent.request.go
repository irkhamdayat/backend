package request

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type AgentRegisterReq struct {
	PhoneNumber       string        `json:"phoneNumber" binding:"required,e164"`
	BirthDate         string        `json:"birthDate" binding:"required"`
	BirthPlace        string        `json:"birthPlace" binding:"required"`
	Address           string        `json:"address" binding:"required"`
	Location          string        `json:"location" binding:"required"`
	KtpDocument       uuid.UUID     `json:"ktpDocument" binding:"required,uuid"`
	KtpNumber         string        `json:"ktpNumber" binding:"required"`
	NpwpDocument      uuid.UUID     `json:"npwpDocument" binding:"required,uuid"`
	NpwpNumber        string        `json:"npwpNumber" binding:"required"`
	BankAccountNumber string        `json:"bankAccountNumber" binding:"required"`
	BankId            uuid.UUID     `json:"bankId" binding:"required,uuid"`
	Pin               string        `json:"pin" binding:"required,min=6,max=6"`
	IsSubscribeNews   bool          `json:"isSubscribeNews"`
	FirstName         string        `json:"firstName" binding:"required,max=255"`
	LastName          null.String   `json:"lastName" binding:"required"`
	Email             string        `json:"email" binding:"required,email"`
	Username          string        `json:"username" binding:"required,min=4,max=50"`
	Password          string        `json:"password" binding:"required,Password"`
	ConfirmPassword   string        `json:"confirmPassword" binding:"required,eqfield=Password"`
	Photo             uuid.NullUUID `json:"photo"`
	OTP               string        `json:"otp" binding:"required"`
}

type AgentLoginReq struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required_without=Username"`
	Password string `json:"password" binding:"required"`
}

type GetAgentInfoReq struct {
	ID uuid.UUID `json:"-"`
}

type PostVerifyEmailReq struct {
	Email string `json:"email" binding:"required"`
}

type PostVerifyPinReq struct {
	Pin string    `json:"pin" binding:"numeric,required"`
	ID  uuid.UUID `json:"-"`
}
