package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// Get device by ID
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getDevice
func GetDevice(deviceID int) (device Device, err error) {
	err = request(http.MethodGet, fmt.Sprintf("device/%d", deviceID), nil, &device)
	return
}

// Returns list of devices for organization
//
// See https://app.ninjarmm.com/apidocs-beta/core-resources/operations/getOrganizationDevices
func ListOrganizationDevices(organizationID int) (devices []Device, err error) {
	err = request(http.MethodGet, fmt.Sprintf("organization/%d/devices", organizationID), nil, &devices)
	return
}

// List all devices with some filters
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getDevices
//
// For filter see
// https://eu.ninjarmm.com/apidocs-beta/core-resources/articles/devices/device-filters
func ListDevices(filter string, detailed bool, after, pageSize int) (devices []Device, err error) {

	urlValues := url.Values{}

	if filter != "" {
		urlValues.Set("df", filter)
	}

	if after != 0 {
		urlValues.Set("after", fmt.Sprint(after))
	}

	if pageSize != 0 {
		urlValues.Set("pageSize", fmt.Sprint(pageSize))
	}

	path := "devices"
	if detailed {
		path = "devices-detailed"
	}

	err = request(http.MethodGet, path+"?"+urlValues.Encode(), nil, &devices)
	return
}

// Find devices by search string
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/search
func FindDevices(search string, limit int) (devices []Device, err error) {

	urlValues := url.Values{}

	if search != "" {
		urlValues.Set("q", search)
	}

	if limit != 0 {
		urlValues.Set("limit", fmt.Sprint(limit))
	}

	err = request(http.MethodGet, "devices/search?"+urlValues.Encode(), nil, &devices)
	return
}

// List all device roles
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeRoles
func ListDeviceRoles() (deviceRoles []DeviceRole, err error) {
	err = request(http.MethodGet, "roles", nil, &deviceRoles)
	return
}

// List all device policies
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getPolicies
func ListDevicePolicies() (policies []Policy, err error) {
	err = request(http.MethodGet, "policies", nil, &policies)
	return
}

// Return device custom fields
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeCustomFields
func GetDeviceCustomFields(deviceID int) (customFields CustomFields, err error) {
	err = request(http.MethodGet, fmt.Sprintf("device/%d/custom-fields", deviceID), nil, &customFields)
	return
}

// Populate device custom fields
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getNodeCustomFields
func (device *Device) GetCustomFields() (err error) {
	var customFields CustomFields
	err = request(http.MethodGet, fmt.Sprintf("device/%d/custom-fields", device.ID), nil, &customFields)
	device.Fields = customFields
	return
}

// Update device custom field values
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/updateNodeAttributeValues

func SetDeviceCustomFields(deviceID int, customFields CustomFields) (err error) {
	err = request(http.MethodPatch, fmt.Sprintf("device/%d/custom-fields", deviceID), customFields, nil)
	return
}

type Device struct {
	ID             int            `json:"id"`
	ParentDeviceID int            `json:"parentDeviceId"`
	OrganizationID int            `json:"organizationId"`
	LocationID     int            `json:"locationId"`
	NodeClass      NodeClass      `json:"nodeClass"`
	NodeRoleID     int            `json:"nodeRoleId"`
	RolePolicyID   int            `json:"rolePolicyId"`
	PolicyID       int            `json:"policyId"`
	ApprovalStatus ApprovalStatus `json:"approvalStatus"`
	Offline        bool           `json:"offline"`
	DisplayName    string         `json:"displayName"`
	SystemName     string         `json:"systemName"`
	DNSName        string         `json:"dnsName"`
	NETBIOSName    string         `json:"netbiosName"`
	Created        Time           `json:"created"`
	LastContact    Time           `json:"lastContact"`
	LastUpdate     Time           `json:"lastUpdate"`
	UserData       CustomFields   `json:"userData"`
	Tags           []string       `json:"tags"`   // seems not implemented yet
	Fields         CustomFields   `json:"fields"` // seems not implemented yet
	Maintenance    struct {
		Status MaintenanceStatus `json:"status"`
		Start  Time              `json:"start"`
		End    Time              `json:"end"`
	} `json:"maintenance"`
	References struct {
		Organization Organization `json:"organization"`
		Location     Location     `json:"location"`
		RolePolicy   Policy       `json:"rolePolicy"`
		Policy       Policy       `json:"policy"`
		Role         DeviceRole   `json:"role"`
		// BackupUsage BackupUsage `json:"backupUsage"`
	} `json:"references"`

	// Only in detailed mode
	IPAddress []string `json:"ipAddress,omitempty"`
	PublicIP  string   `json:"publicIp,omitempty"`
	Notes     []struct {
		Text string `json:"text"`
	} `json:"notes,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
}

type Policy struct {
	ID               int          `json:"id"`
	ParentPolicyID   int          `json:"parentPolicyId"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	NodeClass        NodeClass    `json:"nodeClass"`
	Updated          Time         `json:"updated"`
	NodeClassDefault bool         `json:"nodeClassDefault"`
	Tags             []string     `json:"tags"`
	Fields           CustomFields `json:"fields"`
}

type DeviceRole struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	NodeClass   NodeClass    `json:"nodeClass"`
	Custom      bool         `json:"custom"`
	ChassisType Chassis      `json:"chassisType"`
	Created     Time         `json:"created"`
	Tags        []string     `json:"tags"`
	Fields      CustomFields `json:"fields"`
}

type Chassis string

const (
	ChassisDesktop Chassis = "DESKTOP"
	ChassisLaptop  Chassis = "LAPTOP"
	ChassisMobile  Chassis = "MOBILE"
	ChassisUnknown Chassis = "UNKNOWN"
)

type ApprovalStatus string

const (
	ApprovalStatusPending  ApprovalStatus = "PENDING"
	ApprovalStatusApproved ApprovalStatus = "APPROVED"
)

type MaintenanceStatus string

const (
	MaintenanceStatusFailed       MaintenanceStatus = "FAILED"
	MaintenanceStatusPending      MaintenanceStatus = "PENDING"
	MaintenanceStatusInMaintenace MaintenanceStatus = "IN_MAINTENANCE"
)

type NodeClass string

const (
	NodeClassWindowsServer             NodeClass = "WINDOWS_SERVER"
	NodeClassWindowsWorkstation        NodeClass = "WINDOWS_WORKSTATION"
	NodeClassLinuxWorkstation          NodeClass = "LINUX_WORKSTATION"
	NodeClassMac                       NodeClass = "MAC"
	NodeClassAndroid                   NodeClass = "ANDROID"
	NodeClassAppleIos                  NodeClass = "APPLE_IOS"
	NodeClassAppleIpadOs               NodeClass = "APPLE_IPADOS"
	NodeClassVmwareVmHost              NodeClass = "VMWARE_VM_HOST"
	NodeClassVmwareVmGuest             NodeClass = "VMWARE_VM_GUEST"
	NodeClassHypervVmmHost             NodeClass = "HYPERV_VMM_HOST"
	NodeClassHypervVmmGuest            NodeClass = "HYPERV_VMM_GUEST"
	NodeClassLinuxServer               NodeClass = "LINUX_SERVER"
	NodeClassMacServer                 NodeClass = "MAC_SERVER"
	NodeClassCloudMonitorTarget        NodeClass = "CLOUD_MONITOR_TARGET"
	NodeClassNmsSwitch                 NodeClass = "NMS_SWITCH"
	NodeClassNmsRouter                 NodeClass = "NMS_ROUTER"
	NodeClassNmsFirewall               NodeClass = "NMS_FIREWALL"
	NodeClassNmsPrivateNetworkGateway  NodeClass = "NMS_PRIVATE_NETWORK_GATEWAY"
	NodeClassNmsPrinter                NodeClass = "NMS_PRINTER"
	NodeClassNmsScanner                NodeClass = "NMS_SCANNER"
	NodeClassNmsDialManager            NodeClass = "NMS_DIAL_MANAGER"
	NodeClassNmsWap                    NodeClass = "NMS_WAP"
	NodeClassNmsIpsla                  NodeClass = "NMS_IPSLA"
	NodeClassNmsComputer               NodeClass = "NMS_COMPUTER"
	NodeClassNmsVmHost                 NodeClass = "NMS_VM_HOST"
	NodeClassNmsAppliance              NodeClass = "NMS_APPLIANCE"
	NodeClassNmsOther                  NodeClass = "NMS_OTHER"
	NodeClassNmsServer                 NodeClass = "NMS_SERVER"
	NodeClassNmsPhone                  NodeClass = "NMS_PHONE"
	NodeClassNmsVirtualMachine         NodeClass = "NMS_VIRTUAL_MACHINE"
	NodeClassNmsNetworkManagementAgent NodeClass = "NMS_NETWORK_MANAGEMENT_AGENT"
)
