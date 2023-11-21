package ninjarmm

import (
	"net/http"
)

func CreateOrganization(organization OrganizationDetailed) (err error) {
	err = request(http.MethodPost, "organizations", organization, &organization)
	return
}

func ListOrganizations() (organizations []Organization, err error) {
	err = request(http.MethodGet, "organizations", nil, &organizations)
	return
}

func ListOrganizationsDetailed() (organizations []OrganizationDetailed, err error) {
	err = request(http.MethodGet, "organizations-detailed", nil, &organizations)
	return
}

type Organization struct {
	ID               int          `json:"id,omitempty"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	UserData         any          `json:"userData"`
	NodeApprovalMode ApprovalMode `json:"nodeApprovalMode"`
	Tags             []string     `json:"tags"`
	Fields           Fields       `json:"fields"`
}

type OrganizationDetailed struct {
	ID               int                            `json:"id,omitempty"`
	Name             string                         `json:"name"`
	Description      string                         `json:"description"`
	UserData         any                            `json:"userData"`
	NodeApprovalMode ApprovalMode                   `json:"nodeApprovalMode"`
	Tags             []string                       `json:"tags"`
	Fields           Fields                         `json:"fields"`
	Locations        []Location                     `json:"locations,omitempty"`
	Policies         []OrganizationPolicyItem       `json:"policies"`
	Settings         map[string]OrganizationSetting `json:"settings,omitempty"` // 'trayicon', 'splashtop', 'teamviewer', 'backup' and 'psa'
}

type OrganizationPolicyItem struct {
	NodeRoleID int `json:"nodeRoleId"`
	PolicyID   int `json:"policyId"`
}

type ApprovalMode string

const (
	ApprovalModeAutomatic ApprovalMode = "AUTOMATIC"
	ApprovalModeManual    ApprovalMode = "MANUAL"
	ApprovalModeRefect    ApprovalMode = "REJECT"
)

type OrganizationSetting struct {
	Product string         `json:"product"`
	Enabled bool           `json:"enabled"`
	Targets []string       `json:"targets"`
	Options map[string]any `json:"options"`
}
