package ninjarmm

import (
	"fmt"
	"net/http"
)

// List all users, can be filtered by user type
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getUsers
func ListUsers(userType UserType) (users []User, err error) {

	if userType != UserTypeTechnician && userType != UserTypeEndUser && userType != "" {
		err = fmt.Errorf("invalid user type '%s'", userType)
		return
	}

	err = request(http.MethodGet, "users?userType="+string(userType), nil, &users)

	return
}

// List all users from an organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getEndUsers
func ListOrganizationUsers(organizationID int) (users []User, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/end-users", organizationID), nil, &users)
	return
}

type UserType string

const (
	UserTypeTechnician UserType = "TECHNICIAN"
	UserTypeEndUser    UserType = "END_USER"
)

type InvitationStatus string

const (
	InvitationStatusPending    InvitationStatus = "PENDING"
	InvitationStatusRegistered InvitationStatus = "REGISTERED"
	InvitationStatusExpired    InvitationStatus = "EXPIRED"
)

type User struct {
	ID               int              `json:"id"`
	Firstname        string           `json:"firstname"`
	Lastname         string           `json:"lastname"`
	Email            string           `json:"email"`
	Phone            string           `json:"phone"`
	Enabled          bool             `json:"enabled"`
	Administrator    bool             `json:"administrator"`
	PermitAllClients bool             `json:"permitAllClients"`
	NotifyAllClients bool             `json:"notifyAllClients"`
	MustChangePw     bool             `json:"mustChangePw"`
	MFAConfigured    bool             `json:"mfaConfigured"`
	UserType         UserType         `json:"userType"`         // 'TECHNICIAN' or 'END_USER'
	InvitationStatus InvitationStatus `json:"invitationStatus"` // 'PENDING' or 'REGISTERED' or 'EXPIRED'
	OrganizationID   int              `json:"organizationId"`   // For END_USER only
	DeviceIDs        []int            `json:"deviceIds"`
	Tags             []string         `json:"tags"`
	Fields           CustomFields     `json:"fields"`
}
