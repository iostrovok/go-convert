package convert

import (
	"fmt"
	"strings"
)

func Bool(in interface{}, debugKeys ...string) bool {
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
	}

	s := strings.ToLower(fmt.Sprintf("%s", in))
	return s != "" && s != "false" && s != "f" && s != "0"
}

func BoolErr(in interface{}, debugKeys ...string) (bool, error) {
	return Bool(in), nil
}
