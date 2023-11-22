package ninjarmm

import (
	"net/http"
)

// Get activity log in reverse chronological order
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getActivities
func GetActivityLog(options ActivityLogOptions) (activityLog ActivityLog, err error) {
	err = request(http.MethodGet, "activities?"+options.queryString(), nil, &activityLog)
	return
}

type ActivityLog struct {
	LastActivityID int        `json:"lastActivityId"`
	Activities     []Activity `json:"activities"`
}

type Activity struct {
	ID              int         `json:"id"`
	ActivityTime    Time        `json:"activityTime"`
	DeviceID        int         `json:"deviceId"`
	SeriesUID       string      `json:"seriesUid"`
	StatusCode      string      `json:"statusCode"`
	Status          string      `json:"status"`
	SourceConfigUID string      `json:"sourceConfigUid"`
	SourceName      string      `json:"sourceName"`
	Subject         string      `json:"subject"`
	UserID          int         `json:"userId"`
	Message         string      `json:"message"`
	Type            string      `json:"type"`
	Data            interface{} `json:"data"`

	Severity       `json:"severity"`
	Priority       `json:"priority"`
	ActivityType   `json:"activityType"`
	ActivityResult `json:"activityResult"`
}

type Priority string

const (
	PriorityNone   Priority = "NONE"
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

type Severity string

const (
	SeverityNone     Severity = "NONE"
	SeverityMinor    Severity = "MINOR"
	SeverityModerate Severity = "MODERATE"
	SeverityMajor    Severity = "MAJOR"
	SeverityCritical Severity = "CRITICAL"
)

type ActivityResult string

const (
	ActivityResultSuccess     ActivityResult = "SUCCESS"
	ActivityResultFailure     ActivityResult = "FAILURE"
	ActivityResultUnsupported ActivityResult = "UNSUPPORTED"
	ActivityResultUncompleted ActivityResult = "UNCOMPLETED"
)

type ActivityType string

const (
	ActivityTypeActionSet         ActivityType = "ACTIONSET"
	ActivityTypeAction            ActivityType = "ACTION"
	ActivityTypeCondition         ActivityType = "CONDITION"
	ActvityTypeConditionActionSet ActivityType = ActivityTypeCondition + "_ACTIONSET"
	ActivityTypeConditionAction   ActivityType = ActivityTypeCondition + "_ACTION"
	ActivityTypeAntivirus         ActivityType = "ANTIVIRUS"
	ActivityTypePatchManagement   ActivityType = "PATCH_MANAGEMENT"
	ActivityTypeTeamViewer        ActivityType = "TEAMVIEWER"
	ActivityTypeMonitor           ActivityType = "MONITOR"
	ActivityTypeSystem            ActivityType = "SYSTEM"
	ActivityTypeComment           ActivityType = "COMMENT"
	ActivityTypeShadowProtect     ActivityType = "SHADOWPROTECT"
	ActivityTypeImageManager      ActivityType = "IMAGEMANAGER"
	ActivityTypeHelpRequest       ActivityType = "HELP_REQUEST"
	ActivityTypeSoftwarePatch     ActivityType = "SOFTWARE_PATCH_MANAGEMENT"
	ActivityTypeSplashtop         ActivityType = "SPLASHTOP"
	ActivityTypeCloudBerry        ActivityType = "CLOUDBERRY"
	ActivityTypeCloudBerryBackup  ActivityType = ActivityTypeCloudBerry + "_BACKUP"
	ActivityTypeScheduledTask     ActivityType = "SCHEDULED_TASK"
	ActivityTypeRDP               ActivityType = "RDP"
	ActivityTypeScripting         ActivityType = "SCRIPTING"
	ActivityTypeSecurity          ActivityType = "SECURITY"
	ActivityTypeRemoteTools       ActivityType = "REMOTE_TOOLS"
	ActivityTypeVirtualization    ActivityType = "VIRTUALIZATION"
	ActivityTypePSA               ActivityType = "PSA"
	ActivityTypeMDM               ActivityType = "MDM"
	ActivityTypeNinjaRemote       ActivityType = "NINJA_REMOTE"
	ActivityTypeNinjaQuickConnect ActivityType = "NINJA_QUICK_CONNECT"
	ActivityTypeNinjaDiscovery    ActivityType = "NINJA_NETWORK_DISCOVERY"
)
