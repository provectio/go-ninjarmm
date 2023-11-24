package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

type ActivityLogOptions struct {
	// Return activities recorded after to specified date
	AfterDate string `url:"after,omitempty"`

	// Return activities recorded prior to specified date
	BeforeDate string `url:"before,omitempty"`

	// Activity Class (System/Device) filter (allowed: SYSTEM, DEVICE, USER or ALL) (default: ALL)
	Class string `url:"class,omitempty"`

	// Device filter (See https://eu.ninjarmm.com/apidocs-beta/core-resources/articles/devices/device-filters)
	DeviceFilter string `url:"df,omitempty"`

	// Language tag
	Language string `url:"lang,omitempty"`

	// Return activities recorded that are older than specified activity ID
	NewerThanActivityID int `url:"newerThan,omitempty"`

	// Return activities recorded that are newer than specified activity ID
	OlderThanActivityID int `url:"olderThan,omitempty"`

	// Limit number of activities to return (10 >= v <= 1000) (default: 200)
	PageSize int `url:"pageSize,omitempty"`

	// Return activities related to alert (series)
	SeriesUID string `url:"seriesUid,omitempty"`

	// Directed to a specific script
	SourceConfigUID string `url:"sourceConfigUid,omitempty"`

	// Return activities with status(es)
	Status string `url:"status,omitempty"`

	// Return activities of type
	Type string `url:"type,omitempty"`

	// Time Zone
	TimeZone string `url:"tz,omitempty"`

	// Return activities for user(s)
	User string `url:"user,omitempty"`
}

// Internal function to convert struct `ActivityLogOptions` to url query string
func (options *ActivityLogOptions) queryString() string {

	values := url.Values{}

	typeOf := reflect.TypeOf(*options)

	// Parse struct tag `url`
	for i := 0; i < typeOf.NumField(); i++ {
		// Get tag `url`
		tag := typeOf.Field(i).Tag.Get("url")

		// Skip if tag is empty
		if tag == "" {
			continue
		}

		// Split tag by comma
		splitByComma := strings.Split(tag, ",")
		tag = splitByComma[0]
		omitempty := false

		// If last element is omitempty, set omitempty to true
		if len(splitByComma) > 1 {
			omitempty = splitByComma[len(splitByComma)-1] == "omitempty"
		}

		valueOfField := reflect.ValueOf(*options).Field(i)

		// Skip omitempty if value is empty
		if omitempty && valueOfField.IsZero() {
			continue
		}

		values.Set(tag, fmt.Sprint(valueOfField.Interface()))
	}

	return values.Encode()
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
