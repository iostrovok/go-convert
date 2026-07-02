package convert

import (
	"fmt"
	"reflect"
	"strings"
)

// Bool converts in to a bool using the following rules:
//   - nil            → false
//   - bool           → the value itself
//   - numeric        → false if zero, true otherwise
//   - string / []byte → false for "", "false", "f", "0" (case-insensitive); true otherwise
//   - slice / map    → false if empty, true otherwise
func Bool(in any) bool {
	if in == nil {
		return false
	}

	switch in.(type) {
	case bool:
		return in.(bool)
	case int8, int16, int32, int, int64, uint8, uint16, uint32, uint, uint64:
		return fmt.Sprintf("%d", in) != "0"
	case float32, float64:
		return fmt.Sprintf("%.10f", in) != "0.0000000000"
	case []byte:
		switch strings.ToLower(string(in.([]byte))) {
		case "", "false", "f", "0":
			return false
		default:
			return true
		}
	case string:
		switch strings.ToLower(fmt.Sprintf("%s", in)) {
		case "", "false", "f", "0":
			return false
		default:
			return true
		}
	}

	items := reflect.ValueOf(in)
	if items.Kind() == reflect.Slice || items.Kind() == reflect.Map {
		return items.Len() > 0
	}

	switch strings.ToLower(fmt.Sprintf("%s", in)) {
	case "", "false", "f", "0":
		return false
	default:
		return true
	}
}

// BoolErr converts in to a bool. It always succeeds; the error is always nil.
// See Bool for conversion rules.
func BoolErr(in any) (bool, error) {
	return Bool(in), nil
}
