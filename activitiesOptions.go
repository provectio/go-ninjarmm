package ninjarmm

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

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
