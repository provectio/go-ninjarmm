package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// Creates new location for organization
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/createLocationForOrganization
func CreateLocation(organizationID int, location Location) (createdLocation Location, err error) {
	err = request(http.MethodPost, fmt.Sprintf("organization/%d/locations", organizationID), location, &createdLocation)
	return
}

// Change location name, address, description, custom data
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/updateLocation
func UpdateLocation(organizationID, locationID int, location Location) (err error) {
	err = request(http.MethodPatch, fmt.Sprintf("organization/%d/locations/%d", organizationID, locationID), location, nil)
	return
}

// Returns list of locations for organization
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/getOrganizationLocations
func ListOrganizationLocations(organizationID int) (locations []Location, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/locations", organizationID), nil, &locations)
	return
}

// Returns flat list of all locations for all organizations
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getLocations
func ListLocations(after, pageSize int) (locations []Location, err error) {
	values := url.Values{}

	if after > 0 {
		values.Set("after", fmt.Sprint(after))
	}

	if pageSize > 0 {
		values.Set("pageSize", fmt.Sprint(pageSize))
	}

	err = request(http.MethodGet, "locations?"+values.Encode(), nil, &locations)
	return
}

// Returns list of location custom fields
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeCustomFields_1
func GetLocationCustomFields(organizationID, locationID int) (customFields CustomFields, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/location/%d/custom-fields", organizationID, locationID), nil, &customFields)
	return
}

// Update location custom field values
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/updateNodeAttributeValues_2
func SetLocationCustomFields(organizationID, locationID int, customFields CustomFields) (err error) {
	err = request(http.MethodPatch, fmt.Sprintf("organization/%d/location/%d/custom-fields", organizationID, locationID), customFields, nil)
	return
}

type Location struct {
	ID             int          `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	Address        string       `json:"address,omitempty"`
	Description    string       `json:"description,omitempty"`
	UserData       CustomFields `json:"userData,omitempty"`
	Tags           []string     `json:"tags,omitempty"`           // seems not implemented yet
	Fields         CustomFields `json:"fields,omitempty"`         // seems not implemented yet
	OrganizationID int          `json:"organizationId,omitempty"` // only when list all locations
}
