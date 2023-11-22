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

// Update an organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateOrganization
func UpdateOrganization(organization Organization) (err error) {
	if organization.ID == 0 {
		err = fmt.Errorf("organization ID required")
	} else {
		id := organization.ID
		organization.ID = 0
		err = request(http.MethodPatch, fmt.Sprintf("organization/%d", id), organization, nil)
	}
	return
}

// Update a set of custom fields for an organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateNodeAttributeValues_1
func SetOrganizationCustomFields(organizationID int, customFields CustomFields) (err error) {
	err = request(http.MethodPatch, fmt.Sprintf("organization/%d/custom-fields", organizationID), customFields, nil)
	return
}

// Getting custom fields for an organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeCustomFields_2
func GetOrganizationCustomFields(organizationID int) (customFields CustomFields, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/custom-fields", organizationID), nil, &customFields)
	return
}

// List all organizations
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getOrganizations
func ListOrganizations() (organizations []Organization, err error) {
	err = request(http.MethodGet, "organizations", nil, &organizations)
	return
}

// List all organizations with detailed information
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getOrganizationsDetailed
func ListOrganizationsDetailed() (organizations []OrganizationDetailed, err error) {
	err = request(http.MethodGet, "organizations-detailed", nil, &organizations)
	return
}

type Organization struct {
	ID               int          `json:"id,omitempty"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	UserData         CustomFields `json:"userData"`
	NodeApprovalMode ApprovalMode `json:"nodeApprovalMode"`
	Tags             []string     `json:"tags"`   // seems not implemented yet
	Fields           CustomFields `json:"fields"` // seems not implemented yet
}

type OrganizationDetailed struct {
	ID               int                            `json:"id,omitempty"`
	Name             string                         `json:"name"`
	Description      string                         `json:"description"`
	UserData         CustomFields                   `json:"userData"`
	NodeApprovalMode ApprovalMode                   `json:"nodeApprovalMode"`
	Tags             []string                       `json:"tags"`   // seems not implemented yet
	Fields           CustomFields                   `json:"fields"` // seems not implemented yet
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
