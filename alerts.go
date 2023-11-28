package ninjarmm

import (
	"fmt"
	"net/http"
	"net/url"
)

// List all alerts with some filters
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getAlerts
func ListAlerts(filter string, sourceType AlertOrigin, lang string, tz string) (alerts []Alert, err error) {
	values := url.Values{}

	if filter != "" {
		values.Set("filter", filter)
	}

	if sourceType != "" && sourceType != AlertOriginAll {
		values.Set("sourceType", string(sourceType))
	}

	if lang != "" {
		values.Set("lang", lang)
	}

	if tz != "" {
		values.Set("tz", tz)
	}

	err = request(http.MethodGet, "alerts?"+values.Encode(), nil, &alerts)
	return
}

// List all alerts for a given device ID
//
// See https://eu.ninjarmm.com/apidocs-beta/core-resources/operations/getDeviceAlerts
func ListAlertsDevice(devideID int, lang string, tz string) (alerts []Alert, err error) {
	values := url.Values{}

	if lang != "" {
		values.Set("lang", lang)
	}

	if tz != "" {
		values.Set("tz", tz)
	}

	err = request(http.MethodGet, fmt.Sprintf("device/%d/alerts?%s", devideID, values.Encode()), nil, &alerts)
	return
}

type Alert struct {
	UID              string          `json:"uid"`              // Alert UID (activity series UID)
	DeviceID         int             `json:"deviceId"`         // Device identifier
	Message          string          `json:"message"`          // Alert message
	CreateTime       Time            `json:"createTime"`       // Alert creation timestamp
	UpdateTime       Time            `json:"updateTime"`       // Alert last updated
	SourceType       AlertOrigin     `json:"sourceType"`       // Alert origin
	SourceConfigUID  string          `json:"sourceConfigUid"`  // Source configuration/policy element reference
	SourceName       string          `json:"sourceName"`       // Source configuration/policy element name
	Subject          string          `json:"subject"`          // Alert subject
	UserID           int             `json:"userId"`           // User identifier
	PSATicketID      int             `json:"psaTicketId"`      // related PSA ticket identifier
	TicketTemplateID int             `json:"ticketTemplateId"` // related ticket template identifier
	Data             any             `json:"data"`             // Alert data
	Device           `json:"device"` // ☣️ seems not implemented
}

type AlertOrigin string

const (
	AlertOriginAll                                    AlertOrigin = "" // Use for no filters alerts list
	AlertOriginAgentOffline                           AlertOrigin = "AGENT_OFFLINE"
	AlertOriginAgentCPU                               AlertOrigin = "CONDITION_AGENT_CPU"
	AlertOriginAgentMemory                            AlertOrigin = "CONDITION_AGENT_MEMORY"
	AlertOriginAgentNetwork                           AlertOrigin = "CONDITION_AGENT_NETWORK"
	AlertOriginAgentDiskIO                            AlertOrigin = "CONDITION_AGENT_DISK_IO"
	AlertOriginAgentDiskFreeSpace                     AlertOrigin = "CONDITION_AGENT_DISK_FREE_SPACE"
	AlertOriginAgentDiskUsage                         AlertOrigin = "CONDITION_AGENT_DISK_USAGE"
	AlertOriginAgentCVSSScore                         AlertOrigin = "CONDITION_AGENT_CVSS_SCORE"
	AlertOriginAgentPatchLastInstalled                AlertOrigin = "CONDITION_AGENT_PATCH_LAST_INSTALLED"
	AlertOriginNMSCPU                                 AlertOrigin = "CONDITION_NMS_CPU"
	AlertOriginNMSMemory                              AlertOrigin = "CONDITION_NMS_MEMORY"
	AlertOriginNMSNetworkTrafficBits                  AlertOrigin = "CONDITION_NMS_NETWORK_TRAFFIC_BITS"
	AlertOriginNMSNetworkTrafficPercent               AlertOrigin = "CONDITION_NMS_NETWORK_TRAFFIC_PERCENT"
	AlertOriginNMSNetworkStatus                       AlertOrigin = "CONDITION_NMS_NETWORK_STATUS"
	AlertOriginNMSNetworkStatusChange                 AlertOrigin = "CONDITION_NMS_NETWORK_STATUS_CHANGE"
	AlertOriginPing                                   AlertOrigin = "CONDITION_PING"
	AlertOriginPingLatency                            AlertOrigin = "CONDITION_PING_LATENCY"
	AlertOriginPingPacketLoss                         AlertOrigin = "CONDITION_PING_PACKET_LOSS"
	AlertOriginPingResponse                           AlertOrigin = "CONDITION_PING_RESPONSE"
	AlertOriginSystemUptime                           AlertOrigin = "CONDITION_SYSTEM_UPTIME"
	AlertOriginSmartStatusDegraded                    AlertOrigin = "CONDITION_SMART_STATUS_DEGRATED"
	AlertOriginRaidHealthStatus                       AlertOrigin = "CONDITION_RAID_HEALTH_STATUS"
	AlertOriginScriptResult                           AlertOrigin = "CONDITION_SCRIPT_RESULT"
	AlertOriginHTTP                                   AlertOrigin = "CONDITION_HTTP"
	AlertOriginHTTPResponse                           AlertOrigin = "CONDITION_HTTP_RESPONSE"
	AlertOriginPort                                   AlertOrigin = "CONDITION_PORT"
	AlertOriginPortScan                               AlertOrigin = "CONDITION_PORT_SCAN"
	AlertOriginSyslog                                 AlertOrigin = "CONDITION_SYSLOG"
	AlertOriginConfigurationFile                      AlertOrigin = "CONDITION_CONFIGURATION_FILE"
	AlertOriginSNMPTrap                               AlertOrigin = "CONDITION_SNMPTRAP"
	AlertOriginCriticalEvent                          AlertOrigin = "CONDITION_CRITICAL_EVENT"
	AlertOriginDNS                                    AlertOrigin = "CONDITION_DNS"
	AlertOriginEmail                                  AlertOrigin = "CONDITION_EMAIL"
	AlertOriginCustomSNMP                             AlertOrigin = "CONDITION_CUSTOM_SNMP"
	AlertOriginShadowProtectBackupJobCreate           AlertOrigin = "SHADOWPROTECT_BACKUPJOB_CREATE"
	AlertOriginShadowProtectBackupJobUpdate           AlertOrigin = "SHADOWPROTECT_BACKUPJOB_UPDATE"
	AlertOriginShadowProtectBackupJobDelete           AlertOrigin = "SHADOWPROTECT_BACKUPJOB_DELETE"
	AlertOriginShadowProtectBackupJobExecute          AlertOrigin = "SHADOWPROTECT_BACKUPJOB_EXECUTE"
	AlertOriginImageManagerManagedFolderCreate        AlertOrigin = "IMAGEMANAGER_MANAGEDFOLDER_CREATE"
	AlertOriginImageManagerManagedFolderUpdate        AlertOrigin = "IMAGEMANAGER_MANAGEDFOLDER_UPDATE"
	AlertOriginImageManagerManagedFolderDelete        AlertOrigin = "IMAGEMANAGER_MANAGEDFOLDER_DELETE"
	AlertOriginImageManagerManagedFolderExecute       AlertOrigin = "IMAGEMANAGER_MANAGEDFOLDER_EXECUTE"
	AlertOriginTeamViewerConnection                   AlertOrigin = "TEAMVIEWER_CONNECTION"
	AlertOriginRetrieveAgentLogs                      AlertOrigin = "RETRIEVE_AGENT_LOGS"
	AlertOriginScheduledTask                          AlertOrigin = "SCHEDULED_TASK"
	AlertOriginWindowsEventLogTriggered               AlertOrigin = "CONDITION_WINDOWS_EVENT_LOG_TRIGGERED"
	AlertOriginWindowsServiceStateChanged             AlertOrigin = "CONDITION_WINDOWS_SERVICE_STATE_CHANGED"
	AlertOriginUIMessageActionReboot                  AlertOrigin = "UI_MESSAGE_ACTION_REBOOT"
	AlertOriginUIMessageBDInstallationIssues          AlertOrigin = "UI_MESSAGE_BD_INSTALLATION_ISSUES"
	AlertOriginGravityZoneUIMessageInstallationIssues AlertOrigin = "GRAVITYZONE_UI_MESSAGE_INSTALLATION_ISSUES"
	AlertOriginAVQuarantineThreat                     AlertOrigin = "AV_QUARANTINE_THREAT"
	AlertOriginAVRestoreThreat                        AlertOrigin = "AV_RESTORE_THREAT"
	AlertOriginAVDeleteThreat                         AlertOrigin = "AV_DELETE_THREAT"
	AlertOriginAVRemoveThreat                         AlertOrigin = "AV_REMOVE_THREAT"
	AlertOriginBitdefenderRestoreThreat               AlertOrigin = "BITDEFENDER_RESTORE_THREAT"
	AlertOriginBitdefenderDeleteThreat                AlertOrigin = "BITDEFENDER_DELETE_THREAT"
	AlertOriginConditionBitlockerStatus               AlertOrigin = "CONDITION_BITLOCKER_STATUS"
	AlertOriginConditionFilevaultStatus               AlertOrigin = "CONDITION_FILEVAULT_STATUS"
	AlertOriginConditionLinuxProcess                  AlertOrigin = "CONDITION_LINUX_PROCESS"
	AlertOriginConditionLinuxDaemon                   AlertOrigin = "CONDITION_LINUX_Daemon"
	AlertOriginConditionLinuxProcessResource          AlertOrigin = "CONDITION_LINUX_PROCESS_RESOURCE"
	AlertOriginConditionLinuxProcessResourceCPU       AlertOrigin = "CONDITION_LINUX_PROCESS_RESOURCE_CPU"
	AlertOriginConditionLinuxProcessResourceMemory    AlertOrigin = "CONDITION_LINUX_PROCESS_RESOURCE_MEMORY"
	AlertOriginConditionLinuxDiskFreeSpace            AlertOrigin = "CONDITION_LINUX_DISK_FREE_SPACE"
	AlertOriginConditionLinuxDiskUsage                AlertOrigin = "CONDITION_LINUX_DISK_USAGE"
	AlertOriginConditionVMAggregateCPUUsage           AlertOrigin = "CONDITION_VM_AGGREGATE_CPU_USAGE"
	AlertOriginConditionVMDiskUsage                   AlertOrigin = "CONDITION_VM_DISK_USAGE"
	AlertOriginConditionVMHostDatastore               AlertOrigin = "CONDITION_VM_HOST_DATASTORE"
	AlertOriginConditionVMHostUptime                  AlertOrigin = "CONDITION_VM_HOST_UPTIME"
	AlertOriginConditionVMHostDeviceDown              AlertOrigin = "CONDITION_VM_HOST_DEVICE_DOWN"
	AlertOriginConditionVMHostBadSensors              AlertOrigin = "CONDITION_VM_HOST_BAD_SENSORS"
	AlertOriginConditionVMHostSensorHealth            AlertOrigin = "CONDITION_VM_HOST_SENSOR_HEALTH"
	AlertOriginConditionVMGuestGuestOperationalMode   AlertOrigin = "CONDITION_VM_GUEST_GUEST_OPERATIONAL_MODE"
	AlertOriginConditionVMGuestSnapshotSize           AlertOrigin = "CONDITION_VM_GUEST_SNAPSHOT_SIZE"
	AlertOriginConditionVMGuestSnapshotLifespan       AlertOrigin = "CONDITION_VM_GUEST_SNAPSHOT_LIFESPAN"
	AlertOriginConditionVMGuestToolsNotRunning        AlertOrigin = "CONDITION_VM_GUEST_TOOLS_NOT_RUNNING"
	AlertOriginConditionHVGuestCheckpointSize         AlertOrigin = "CONDITION_HV_GUEST_CHECKPOINT_SIZE"
	AlertOriginConditionHVGuestCheckpointLifespan     AlertOrigin = "CONDITION_HV_GUEST_CHECKPOINT_LIFESPAN"
	AlertOriginConditionSoftware                      AlertOrigin = "CONDITION_SOFTWARE"
	AlertOriginConditionWindowsProcessState           AlertOrigin = "CONDITION_WINDOWS_PROCESS_STATE"
	AlertOriginConditionWindowsProcessResourceCPU     AlertOrigin = "CONDITION_WINDOWS_PROCESS_RESOURCE_CPU"
	AlertOriginConditionWindowsProcessResourceMemory  AlertOrigin = "CONDITION_WINDOWS_PROCESS_RESOURCE_MEMORY"
	AlertOriginConditionMacProcessState               AlertOrigin = "CONDITION_MAC_PROCESS_STATE"
	AlertOriginConditionMacProcessResourceCPU         AlertOrigin = "CONDITION_MAC_PROCESS_RESOURCE_CPU"
	AlertOriginConditionMacProcessResourceMemory      AlertOrigin = "CONDITION_MAC_PROCESS_RESOURCE_MEMORY"
	AlertOriginConditionMacDeamon                     AlertOrigin = "CONDITION_MAC_DEAMON"
	AlertOriginConditionCustomField                   AlertOrigin = "CONDITION_CUSTOM_FIELD"
	AlertOriginConditionPendingReboot                 AlertOrigin = "CONDITION_PENDING_REBOOT"
)
