package request

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type AdminLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminRegisterReq struct {
	FirstName        string        `json:"firstName" binding:"required,max=255"`
	LastName         null.String   `json:"lastName" binding:"required"`
	Email            string        `json:"email" binding:"required,email"`
	Username         string        `json:"username" binding:"required,min=4,max=50"`
	RoleID           uuid.UUID     `json:"roleId" binding:"required,uuid"`
	Password         string        `json:"password" binding:"required,Password"`
	ConfirmPassword  string        `json:"confirmPassword" binding:"required,eqfield=Password"`
	InsuranceBrandID null.String   `json:"insuranceBrandId"`
	Photo            uuid.NullUUID `json:"photo"`
}

type AdminActivationReq struct {
	ActivationToken string `json:"activationToken" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
}

type AdminForgotPasswordReq struct {
	Email string `json:"email" binding:"required,email"`
}

type AdminChangePasswordReq struct {
	Token           string `json:"token" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,Password"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

type GetAdminInfoReq struct {
	ID uuid.UUID `json:"-"`
}
