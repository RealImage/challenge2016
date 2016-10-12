// generated by jsonenums -type=Permission; DO NOT EDIT

package domain

import (
	"encoding/json"
	"fmt"
)

var (
	_PermissionNameToValue = map[string]Permission{
		"Granted":    Granted,
		"Denied":     Denied,
		"NotDefined": NotDefined,
	}

	_PermissionValueToName = map[Permission]string{
		Granted:    "Granted",
		Denied:     "Denied",
		NotDefined: "NotDefined",
	}
)

func init() {
	var v Permission
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_PermissionNameToValue = map[string]Permission{
			interface{}(Granted).(fmt.Stringer).String():    Granted,
			interface{}(Denied).(fmt.Stringer).String():     Denied,
			interface{}(NotDefined).(fmt.Stringer).String(): NotDefined,
		}
	}
}

// MarshalJSON is generated so Permission satisfies json.Marshaler.
func (r Permission) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PermissionValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Permission: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Permission satisfies json.Unmarshaler.
func (r *Permission) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Permission should be a string, got %s", data)
	}
	v, ok := _PermissionNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Permission %q", s)
	}
	*r = v
	return nil
}