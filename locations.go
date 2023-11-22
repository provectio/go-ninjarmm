package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getLocations
func ListLocations(after, pageSize int) (locations []Location, err error) {
	values := url.Values{
		"after":    []string{fmt.Sprint(after)},
		"pageSize": []string{fmt.Sprint(pageSize)},
	}

	err = request(http.MethodGet, "locations?"+values.Encode(), nil, &locations)
	return
}

type Location struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Address        string       `json:"address"`
	Description    string       `json:"description"`
	UserData       CustomFields `json:"userData"`
	Tags           []string     `json:"tags"`
	Fields         CustomFields `json:"fields"`
	OrganizationID int          `json:"organizationId,omitempty"` // only when list all locations
}
