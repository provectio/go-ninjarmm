package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// Create an organization, optionally based on a template organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/createOrganization
func CreateOrganization(newOrganization OrganizationDetailed, model_id int) (createdOrganization OrganizationDetailed, err error) {

	path := "organizations"
	if model_id != 0 {
		values := url.Values{
			"templateOrganizationId": {fmt.Sprint(model_id)},
		}
		path += "?" + values.Encode()
	}

	err = request(http.MethodPost, path, newOrganization, &createdOrganization)

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
