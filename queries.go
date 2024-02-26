package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// Query computer systems device informations
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getComputerSystems
func QueryComputerSystems(filter string, pageSize int) (report ComputerSystemReport, err error) {
	urlValues := url.Values{}

	if filter != "" {
		urlValues.Set("df", filter)
	}

	if pageSize != 0 {
		urlValues.Set("pageSize", fmt.Sprint(pageSize))
	}

	err = request(http.MethodGet, "queries/computer-systems?"+urlValues.Encode(), nil, &report)
	return
}

type ComputerSystemReport struct {
	Cursor  ReportCursor     `json:"cursor"`
	Results []ComputerSystem `json:"results"`
}

type ComputerSystem struct {
	Name                string `json:"name"`
	Manufacturer        string `json:"manufacturer"`
	Model               string `json:"model"`
	BiosSerialNumber    string `json:"biosSerialNumber"`
	SerialNumber        string `json:"serialNumber"`
	Domain              string `json:"domain"`
	DomainRole          string `json:"domainRole"`
	NumberOfProcessors  int    `json:"numberOfProcessors"`
	TotalPhysicalMemory int    `json:"totalPhysicalMemory"`
	VirtualMachine      bool   `json:"virtualMachine"`
	ChassisType         string `json:"chassisType"`
	DeviceID            int    `json:"deviceId"`
	Timestamp           Time   `json:"timestamp"`
}

// Query operating systems device informations
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getOperatingSystems
func QueryOperatingSystems(filter string, pageSize int) (report OperatingSystemReport, err error) {
	urlValues := url.Values{}

	if filter != "" {
		urlValues.Set("df", filter)
	}

	if pageSize != 0 {
		urlValues.Set("pageSize", fmt.Sprint(pageSize))
	}

	err = request(http.MethodGet, "queries/operating-systems?"+urlValues.Encode(), nil, &report)
	return
}

type OperatingSystemReport struct {
	Cursor  ReportCursor      `json:"cursor"`
	Results []OperatingSystem `json:"results"`
}

type OperatingSystem struct {
	Name                    string `json:"name"`
	Manufacturer            string `json:"manufacturer"`
	Architecture            string `json:"architecture"`
	LastBootTime            Time   `json:"lastBootTime"`
	BuildNumber             string `json:"buildNumber"`
	ReleaseID               string `json:"releaseId"`
	ServicePackMajorVersion int    `json:"servicePackMajorVersion"`
	ServicePackMinorVersion int    `json:"servicePackMinorVersion"`
	Locale                  string `json:"locale"`
	Language                string `json:"language"`
	NeedsReboot             bool   `json:"needsReboot"`
	DeviceID                int    `json:"deviceId"`
	Timestamp               Time   `json:"timestamp"`
}

// Global query cursor
type ReportCursor struct {
	Name    string `json:"name"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
	Expires Time   `json:"expires"`
}
