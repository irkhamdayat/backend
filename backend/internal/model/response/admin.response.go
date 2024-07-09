package response

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type CompanyUserRoleResp struct {
	ID   uuid.NullUUID `json:"id"`
	Name string        `json:"name"`
}

type GetAdminInfoResp struct {
	ID        uuid.UUID   `json:"id"`
	Photo     *Media      `json:"photo"`
	FirstName string      `json:"firstName"`
	LastName  null.String `json:"lastName"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Role      Role        `json:"role"`
	//TODO: Fix this response, should be insurance with insurance response
	InsuranceBrandID null.String `json:"insuranceBrandId"`
	Status           string      `json:"status"`
}
