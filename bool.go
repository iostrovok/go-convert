package convert

import (
	"fmt"
	"reflect"
	"strings"
)

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

func BoolErr(in any) (bool, error) {
	return Bool(in), nil
}
