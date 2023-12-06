package ninjarmm

import (
	"fmt"
	"net/http"
)

// Returns organisation documents
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getOrganizationDocuments

func GetOrganizationDocuments(organizationID int) (documents []Document, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/documents", organizationID), nil, &documents)
	return
}

// Update organization document
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateOrganizationDocument

func UpdateOrganizationDocument(organizationID int, document Document) (err error) {
	err = request(http.MethodPost, fmt.Sprintf("organization/%d/document/%d", organizationID, document.ClientDocumentID), nil, nil)
	return
}

type Document struct {
	ClientDocumentID          int             `json:"clientDocumentId,omitempty"`
	ClientDocumentName        string          `json:"clientDocumentName"`
	ClientDocumentDescription string          `json:"clientDocumentDescription"`
	ClientDocumentUpdateTime  int             `json:"clientDocumentUpdateTime"`
	AttributeValues           []DocumentValue `json:"attributeValues,omitempty"`
}

type DocumentValue struct {
	Value           interface{} `json:"value"`
	ValueUpdateTime int         `json:"valueUpdateTime"`
	AttributeName   string      `json:"attributeName"`
}
