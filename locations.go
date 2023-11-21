package ninjarmm

import "net/http"

func ListLocations() (locations []Location, err error) {
	err = request(http.MethodGet, "locations", nil, &locations)
	return
}

type Location struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Address        string         `json:"address"`
	Description    string         `json:"description"`
	UserData       any            `json:"userData"`
	Tags           []string       `json:"tags"`
	Fields         map[string]any `json:"fields"`
	OrganizationID int            `json:"organizationId"` // only when list all locations
}
