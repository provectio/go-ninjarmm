package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// Returns organization details (policy mappings, locations)
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/getOrganization
func GetOrganization(organizationID int) (organization OrganizationDetailed, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d", organizationID), nil, &organization)
	return
}

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

// Populate custom fields for an organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeCustomField_2
func (organization *Organization) GetCustomFields() (err error) {
	var customFields CustomFields
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/custom-fields", organization.ID), nil, &customFields)
	organization.Fields = customFields
	return
}

// Populate custom fields for a detailed organization
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeCustomField_2
func (organization *OrganizationDetailed) GetCustomFields() (err error) {
	var customFields CustomFields
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/custom-fields", organization.ID), nil, &customFields)
	organization.Fields = customFields
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

// Change organization policy mappings
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateNodeRolePolicyAssignmentForOrganization
func (organization *OrganizationDetailed) UpdatePolicies(policies []OrganizationPolicyItem) (affectedDevicesIDs []int, err error) {
	err = request(http.MethodPut, fmt.Sprintf("organization/%d/policies", organization.ID), policies, &affectedDevicesIDs)
	return
}

// Change organization policy mappings
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateNodeRolePolicyAssignmentForOrganization
func (organization *Organization) UpdatePolicies(policies []OrganizationPolicyItem) (affectedDevicesIDs []int, err error) {
	err = request(http.MethodPut, fmt.Sprintf("organization/%d/policies", organization.ID), policies, &affectedDevicesIDs)
	return
}

// Change organization policy mappings
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateNodeRolePolicyAssignmentForOrganization
func UpdateOrganizationPolicies(organizationID int, policies []OrganizationPolicyItem) (affectedDevicesIDs []int, err error) {
	err = request(http.MethodPut, fmt.Sprintf("organization/%d/policies", organizationID), policies, &affectedDevicesIDs)
	return
}

type Organization struct {
	ID               int          `json:"id,omitempty"`
	Name             string       `json:"name,omitempty"`
	Description      string       `json:"description,omitempty"`
	UserData         CustomFields `json:"userData,omitempty"`
	NodeApprovalMode ApprovalMode `json:"nodeApprovalMode,omitempty"`
	Tags             []string     `json:"tags,omitempty"`   // seems not implemented yet
	Fields           CustomFields `json:"fields,omitempty"` // seems not implemented yet
}

type OrganizationDetailed struct {
	ID               int                            `json:"id,omitempty"`
	Name             string                         `json:"name,omitempty"`
	Description      string                         `json:"description,omitempty"`
	UserData         CustomFields                   `json:"userData,omitempty"`
	NodeApprovalMode ApprovalMode                   `json:"nodeApprovalMode,omitempty"`
	Tags             []string                       `json:"tags,omitempty"`   // seems not implemented yet
	Fields           CustomFields                   `json:"fields,omitempty"` // seems not implemented yet
	Locations        []Location                     `json:"locations,omitempty"`
	Policies         []OrganizationPolicyItem       `json:"policies,omitempty"`
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
