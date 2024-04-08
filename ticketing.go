package ninjarmm

import (
	"errors"
	"fmt"
	"net/http"
)

// Create new ticket, does not accept files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/create
func (newTicket NewTicket) Create() (createdTicket Ticket, err error) {
	err = request(http.MethodPost, "ticketing/ticket", newTicket, &createdTicket)
	return
}

// Create new ticket, does not accept files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/create
func CreateTicket(newTicket NewTicket) (createdTicket Ticket, err error) {
	return newTicket.Create()
}

// [W.I.P] Add a new comment to a ticket, allows files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/createComment
func (ticket Ticket) AddComment(comment TicketComment) (err error) {
	return errors.New("add comment not implemented, feel free to contribute")
}

// Returns a ticket
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getTicketById
func GetTicket(ticketID int) (ticket Ticket, err error) {
	if ticketID == 0 {
		err = errors.New("ticket ID required")
	} else {
		err = request(http.MethodGet, fmt.Sprintf("ticketing/ticket/%d", ticketID), nil, &ticket)
	}
	return
}

// Change ticket fields. Does not accept comments or files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/update
func (ticket Ticket) Update() (updatedTicket Ticket, err error) {
	if ticket.ID == 0 {
		err = errors.New("ticket ID required")
	} else {
		ticketID := ticket.ID
		ticket.ID = 0
		err = request(http.MethodPut, fmt.Sprintf("ticketing/ticket/%d", ticketID), ticket, &updatedTicket)
	}

	return
}

// Change ticket fields. Does not accept comments or files
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/update
func UpdateTicket(ticket Ticket) (updatedTicket Ticket, err error) {
	return ticket.Update()
}

// Returns list of the ticket log entries for a ticket
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getTicketLogEntriesByTicketId
func (ticket Ticket) GetLog() (log []TicketLog, err error) {
	return GetTicketLog(ticket.ID)
}

// Returns list of the ticket log entries for a ticket
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getTicketLogEntriesByTicketId
func GetTicketLog(ticketID int) (log []TicketLog, err error) {
	if ticketID == 0 {
		err = errors.New("ticket ID required")
	} else {
		err = request(http.MethodGet, fmt.Sprintf("ticketing/ticket/%d/log-entry", ticketID), nil, &log)
	}
	return
}

// Returns list of contacts
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getContacts
func ListContacts() (contacts []Contact, err error) {
	err = request(http.MethodGet, "ticketing/contact/contacts", nil, &contacts)
	return
}

// Returns list of ticketing boards
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getBoards
func ListTicketingBoards() (boards []TicketingBoard, err error) {
	err = request(http.MethodGet, "ticketing/trigger/boards", nil, &boards)
	return
}

// Run a board. Returns list of tickets matching the board condition and filters. Allows pagination
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getTicketsByBoard
func ListTicketsByBoard(boardID int, options ListTicketsOptions) (tickets BoardTickets, err error) {
	if boardID == 0 {
		err = errors.New("board ID required")
	} else {
		err = request(http.MethodPost, fmt.Sprintf("ticketing/trigger/board/%d/run", boardID), options, &tickets)
	}
	return
}

type Ticket struct {
	ID                int                `json:"id,omitempty"`
	Version           int                `json:"version,omitempty"`
	NodeID            int                `json:"nodeId,omitempty"`
	ClientID          int                `json:"clientId,omitempty"`
	LocationID        int                `json:"locationId,omitempty"`
	AssignedAppUserID int                `json:"assignedAppUserId,omitempty"`
	RequesterUID      string             `json:"requesterUid,omitempty"`
	Subject           string             `json:"subject,omitempty"`
	Status            TicketStatus       `json:"status,omitempty"`
	Type              TicketType         `json:"type,omitempty"`
	TicketFormID      int                `json:"ticketFormId,omitempty"`
	Source            string             `json:"source,omitempty"`
	Tags              []string           `json:"tags,omitempty"`
	CcList            TicketCcList       `json:"ccList,omitempty"`
	CreateTime        Time               `json:"createTime,omitempty"`
	Deleted           bool               `json:"deleted,omitempty"`
	AttributeValues   []TicketAttributes `json:"attributeValues,omitempty"`

	Priority `json:"priority,omitempty"`
	Severity `json:"severity,omitempty"`
}

type TicketStatus struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ParentID    int    `json:"parentId,omitempty"`
	StatusID    int    `json:"statusId,omitempty"`
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
	ClientID          int                `json:"clientId"`             // REQUIRED / Client (Organization) identifier
	TicketFormID      int                `json:"ticketFormId"`         // REQUIRED / Ticket form identifier
	LocationID        int                `json:"locationId,omitempty"` // Location identifier
	NodeID            int                `json:"nodeId,omitempty"`     // Device identifier
	Subject           string             `json:"subject"`              // REQUIRED / Ticket subject (>= 0 characters <= 200 characters)
	Description       TicketDescription  `json:"description"`          // REQUIRED / Ticket description
	Status            string             `json:"status"`               // REQUIRED / Ticket status
	Type              TicketType         `json:"type,omitempty"`       // Allowed values: 'PROBLEM', 'QUESTION', 'INCIDENT' or 'TASK'
	Cc                TicketCcList       `json:"cc,omitempty"`
	AssignedAppUserID int                `json:"assignedAppUserId,omitempty"`
	RequesterUID      string             `json:"requesterUid,omitempty"`
	Severity          string             `json:"severity,omitempty"`
	Priority          string             `json:"priority,omitempty"`
	ParentTicketID    int                `json:"parentTicketId,omitempty"`
	Tags              []string           `json:"tags,omitempty"`
	Attributes        []TicketAttributes `json:"attributes,omitempty"`
}

// W.I.P
type TicketComment struct {
	Comment TicketDescription `form:"comment"`
	Files   map[string][]byte `form:"files"`
}

type TicketDescription struct {
	Public               bool   `json:"public"` // REQUIRED
	Body                 string `json:"body"`
	HTMLBody             string `json:"htmlBody"`
	TimeTracked          int    `json:"timeTracked"`
	DuplicateInIncidents bool   `json:"duplicateInIncidents"`
}

type TicketLog struct {
	ID                        int                 `json:"id"`
	AppUserContactUID         string              `json:"appUserContactUid"`
	AppUserContactID          int                 `json:"appUserContactId"`
	AppUserContactType        UserType            `json:"appUserContactType"`
	Type                      TicketLogType       `json:"type"`
	Body                      string              `json:"body"`
	HTMLBody                  string              `json:"htmlBody"`
	FullEmailBody             string              `json:"fullEmailBody"`
	PublicEntry               bool                `json:"publicEntry"`
	System                    bool                `json:"system"`
	CreateTime                Time                `json:"createTime"`
	ChangeDiff                any                 `json:"changeDiff"`
	ActivityID                int                 `json:"activityId"`
	TimeTracked               int                 `json:"timeTracked"`
	TechnicianTagged          []int               `json:"technicianTagged"`
	TechniciansTaggedMetadata []TechniciansTagged `json:"techniciansTaggedMetadata"`
	Automation                TicketLogAutomation `json:"automation"`
	BlockedByInvoice          bool                `json:"blockedByInvoice"`
	EmailResponse             bool                `json:"emailResponse"`
}

type TechniciansTagged struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Deleted     bool   `json:"deleted"`
	Permitted   bool   `json:"permitted"`
}

type TicketLogAutomation struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	System bool   `json:"system"`
	Type   string `json:"type"`
}

type TicketLogType string

const (
	TicketLogTypeComment     TicketLogType = "COMMENT"
	TicketLogTypeDescription TicketLogType = "DESCRIPTION"
	TicketLogTypeCondition   TicketLogType = "CONDITION"
	TicketLogTypeSave        TicketLogType = "SAVE"
	TicketLogTypeDelete      TicketLogType = "DELETE"
)

type TicketType string

const (
	TicketTypeNone     TicketType = ""
	TicketTypeProblem  TicketType = "PROBLEM"
	TicketTypeQuestion TicketType = "QUESTION"
	TicketTypeIncident TicketType = "INCIDENT"
	TicketTypeTask     TicketType = "TASK"
)

type Contact struct {
	ID          int    `json:"id"`
	ClientID    int    `json:"clientId"`
	ClientName  string `json:"clientName"`
	UID         string `json:"uid"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	JobTitle    string `json:"jobTitle"`
	DisplayName string `json:"displayName"`
	CreateTime  Time   `json:"createTime"`
	UpdateTime  Time   `json:"updateTime"`
}

type TicketingBoard struct {
	ID          int             `json:"id"`
	UID         string          `json:"uid"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Conditions  BoardConditions `json:"conditions"`
	CreateTime  int             `json:"createTime"`
	UpdateTime  int             `json:"updateTime"`
	System      bool            `json:"system"`
	Columns     []string        `json:"columns"`
	SortBy      map[string]any  `json:"sortBy"`
	TicketCount int             `json:"ticketCount"`
}

type BoardCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type BoardConditions struct {
	Any []BoardCondition `json:"any"`
	All []BoardCondition `json:"all"`
}

type ListTicketsOptions struct {
	SortBy         []SortBy         `json:"sortBy,omitempty"`
	Filters        []BoardCondition `json:"filters,omitempty"`
	PageSize       int              `json:"pageSize,omitempty"`
	SearchCriteria string           `json:"searchCriteria,omitempty"`
	IncludeColumns []string         `json:"includeColumns,omitempty"`
	LastCursorID   int              `json:"lastCursorId,omitempty"`
}

type SortBy struct {
	Field     string `json:"field"`
	Direction string `json:"direction"` // 'ASC' or 'DESC'
}

type Filters struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type BoardTickets struct {
	Data     []BoardTicket `json:"data"`
	Metadata BoardMetadata `json:"metadata"`
}

type BoardTicket struct {
	AssignedAppUser    string                        `json:"assignedAppUser"`
	CreateTime         Time                          `json:"createTime"`
	Description        string                        `json:"description"`
	Device             string                        `json:"device"`
	ID                 int                           `json:"id"`
	Organization       string                        `json:"organization"`
	Priority           string                        `json:"priority"`
	Severity           string                        `json:"severity"`
	Status             BoardTicketStatus             `json:"status"`
	Summary            string                        `json:"summary"`
	Tags               []string                      `json:"tags"`
	TicketForm         string                        `json:"ticketForm"`
	TotalTimeTracked   int                           `json:"totalTimeTracked"`
	TriggeredCondition BoardTicketTriggeredCondition `json:"triggeredCondition"`
}

type BoardMetadata struct {
	Columns                 []string `json:"columns"`
	SortBy                  []SortBy `json:"sortBy"`
	Attributes              any      `json:"attributes"`
	Filters                 any      `json:"filters"`
	LastCursorID            int      `json:"lastCursorId"`
	AllRequiredColumns      []string `json:"allRequiredColumns"`
	AllColumns              []string `json:"allColumns"`
	ColumnNamesForExporting []string `json:"columnNamesForExporting"`
}

type BoardTicketStatus struct {
	StatusID    int    `json:"statusId"`
	DisplayName string `json:"displayName"`
	ParentID    int    `json:"parentId"`
}

type BoardTicketTriggeredCondition struct {
	UID        string `json:"uid"`
	Message    string `json:"message"`
	CreateTime Time   `json:"createTime"`
}
