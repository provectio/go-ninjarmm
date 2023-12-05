package ninjarmm

import (
	"fmt"
	"strconv"
)

// Shortcuts for map[string]interface{}
type CustomFields map[string]interface{}

// Set sets a custom field.
// Don't forget to call make(CustomFields) before using this function.
//
// Example:
//
//	customFields := make(CustomFields)
//	customFields.Set("myField", "myValue")
func (c CustomFields) Set(key string, value interface{}) {
	c[key] = value
}

// Get returns the interface{} value from the custom field.
func (c CustomFields) Get(key string) interface{} {
	return c[key]
}

// InterfaceField returns the interface{} value from the custom field.
func (c CustomFields) InterfaceField(key string) interface{} {
	return c[key]
}

// StringField returns a string value from the custom field.
func (c CustomFields) StringField(key string) string {
	switch i := c[key].(type) {
	case string:
		return i
	case nil:
		return ""
	default:
		return fmt.Sprint(i)
	}
}

// IntField returns a parsed int value from the custom field.
func (c CustomFields) IntField(key string) int {
	switch i := c[key].(type) {
	case int:
		return i
	case float64:
		return int(i)
	case string:
		conv, _ := strconv.Atoi(i)
		return conv
	default:
		return 0
	}
}

// FloatField returns a parsed float64 value from the custom field.
func (c CustomFields) FloatField(key string) float64 {
	switch i := c[key].(type) {
	case float64:
		return i
	case int:
		return float64(i)
	case string:
		conv, _ := strconv.ParseFloat(i, 64)
		return conv
	default:
		return 0
	}
}

// BoolField returns a parsed boolean value from the custom field.
func (c CustomFields) BoolField(key string) bool {
	switch i := c[key].(type) {
	case bool:
		return i
	case string:
		conv, _ := strconv.ParseBool(i)
		return conv
	case int:
		return i > 0
	case float64:
		return i > 0
	default:
		return false
	}
}
