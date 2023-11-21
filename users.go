package ninjarmm

import (
	"fmt"
	"net/http"
)

func ListUsers(userType UserType) (users []User, err error) {

	if userType != UserTypeTechnician && userType != UserTypeEndUser && userType != "" {
		err = fmt.Errorf("invalid user type '%s'", userType)
		return
	}

	err = request(http.MethodGet, "users?userType="+string(userType), nil, &users)

	return
}

type UserType string

const (
	UserTypeTechnician UserType = "TECHNICIAN"
	UserTypeEndUser    UserType = "END_USER"
)

type InvitationStatusType string

const (
	InvitationStatusPending    InvitationStatusType = "PENDING"
	InvitationStatusRegistered InvitationStatusType = "REGISTERED"
	InvitationStatusExpired    InvitationStatusType = "EXPIRED"
)

type User struct {
	ID               int                  `json:"id"`
	Firstname        string               `json:"firstname"`
	Lastname         string               `json:"lastname"`
	Email            string               `json:"email"`
	Phone            string               `json:"phone"`
	Enabled          bool                 `json:"enabled"`
	Administrator    bool                 `json:"administrator"`
	PermitAllClients bool                 `json:"permitAllClients"`
	NotifyAllClients bool                 `json:"notifyAllClients"`
	MustChangePw     bool                 `json:"mustChangePw"`
	MFAConfigured    bool                 `json:"mfaConfigured"`
	UserType         UserType             `json:"userType"`         // 'TECHNICIAN' or 'END_USER'
	InvitationStatus InvitationStatusType `json:"invitationStatus"` // 'PENDING' or 'REGISTERED' or 'EXPIRED'
	OrganizationID   int                  `json:"organizationId"`
	DeviceIDs        []int                `json:"deviceIds"`
	Tags             []string             `json:"tags"`
	Fields           map[string]any       `json:"fields"`
}
