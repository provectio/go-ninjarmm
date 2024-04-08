package ninjarmm

import "net/http"

// Create new ticket, does not accept files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/create
func CreateTicket(newTicket NewTicket) (createdTicket Ticket, err error) {
	return newTicket.Create()
}

// Create new ticket, does not accept files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/create
func (newTicket NewTicket) Create() (createdTicket Ticket, err error) {
	err = request(http.MethodPost, "ticketing/ticket", newTicket, &createdTicket)
	return
}

type Ticket struct {
	ID                int                `json:"id"`
	Version           int                `json:"version"`
	NodeID            int                `json:"nodeId"`
	ClientID          int                `json:"clientId"`
	LocationID        int                `json:"locationId"`
	AssignedAppUserID int                `json:"assignedAppUserId"`
	RequesterUID      string             `json:"requesterUid"`
	Subject           string             `json:"subject"`
	Status            TicketStatus       `json:"status"`
	Type              string             `json:"type"` // Allowed values: 'PROBLEM', 'QUESTION', 'INCIDENT' or 'TASK'
	TicketFormID      int                `json:"ticketFormId"`
	Source            string             `json:"source"`
	Tags              []string           `json:"tags"`
	CcList            TicketCcList       `json:"ccList"`
	CreateTime        int                `json:"createTime"`
	Deleted           bool               `json:"deleted"`
	AttributeValues   []TicketAttributes `json:"attributeValues"`

	Priority `json:"priority"`
	Severity `json:"severity"`
}

type TicketStatus struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ParentID    int    `json:"parentId"`
	StatusID    int    `json:"statusId"`
}

type TicketCcList struct {
	Uids   []string `json:"uids"`
	Emails []string `json:"emails"`
}

type TicketAttributes struct {
	ID          int `json:"id"`
	AttributeID int `json:"attributeId"`
	Value       any `json:"value"`
}

// NewTicket object that needs to be added to the store
type NewTicket struct {
	ClientID          int                  `json:"clientId"`     // REQUIRED / Client (Organization) identifier
	TicketFormID      int                  `json:"ticketFormId"` // REQUIRED / Ticket form identifier
	LocationID        int                  `json:"locationId"`   // Location identifier
	NodeID            int                  `json:"nodeId"`       // Device identifier
	Subject           string               `json:"subject"`      // REQUIRED / Ticket subject (>= 0 characters <= 200 characters)
	Description       NewTicketDescription `json:"description"`  // REQUIRED / Ticket description
	Status            string               `json:"status"`       // REQUIRED / Ticket status
	Type              string               `json:"type"`         // Allowed values: 'PROBLEM', 'QUESTION', 'INCIDENT' or 'TASK'
	Cc                TicketCcList         `json:"cc"`
	AssignedAppUserID int                  `json:"assignedAppUserId"`
	RequesterUID      string               `json:"requesterUid"`
	Severity          string               `json:"severity"`
	Priority          string               `json:"priority"`
	ParentTicketID    int                  `json:"parentTicketId"`
	Tags              []string             `json:"tags"`
	Attributes        []TicketAttributes   `json:"attributes"`
}

type NewTicketDescription struct {
	Public               bool   `json:"public"` // REQUIRED
	Body                 string `json:"body"`
	HTMLBody             string `json:"htmlBody"`
	TimeTracked          int    `json:"timeTracked"`
	DuplicateInIncidents bool   `json:"duplicateInIncidents"`
}
