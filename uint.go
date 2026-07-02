package convert

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// Uint32 converts in to a uint32. Returns 0 on error.
// If debugKeys are provided and an error occurs, the error is logged.
// Optional debugKeys are embedded in error messages for tracing.
func Uint32(in any, debugKeys ...string) uint32 {
	out, err := Uint32Err(in, debugKeys...)
	if err != nil {
		if len(debugKeys) > 0 {
			log.Printf("Uint32 wrong value for '%+v' [keys: %+v], error: %v", in, debugKeys, err)
		}
		return 0
	}

	return out
}

// Uint32Err converts in to a uint32. Returns an error if the conversion fails
// or if the value exceeds math.MaxUint32.
// Optional debugKeys are embedded in error messages for tracing.
func Uint32Err(in any, debugKeys ...string) (uint32, error) {
	u64, err := Uint64Err(in, debugKeys...)
	if err != nil {
		return 0, err
	}

	if u64 > math.MaxUint32 {
		return 0, fmt.Errorf("Uint32Err wrong value for '%+v'", in)
	}

	return uint32(u64), nil
}

// Uint64 converts in to a uint64. Returns 0 on error.
// If debugKeys are provided and an error occurs, the error is logged.
// Optional debugKeys are embedded in error messages for tracing.
func Uint64(in any, debugKeys ...string) uint64 {

	out, err := Uint64Err(in, debugKeys...)
	if err != nil {
		if len(debugKeys) > 0 {
			log.Printf("Uint64 wrong value for '%+v' [keys: %+v]", in, debugKeys)
		}
	}

	return out
}

// Uint64Err converts in to a uint64. Accepted types: all integer and float types,
// bool (true→1, false→0), string, and []byte. nil is treated as 0 (no error).
// Returns an error for unconvertible values.
// Optional debugKeys are embedded in error messages for tracing.
func Uint64Err(in any, debugKeys ...string) (uint64, error) {

	if in == nil {
		return uint64(0), nil
	}

	switch in.(type) {
	case uint64:
		return in.(uint64), nil

	case bool:
		if in.(bool) {
			return uint64(1), nil
		}
		return uint64(0), nil

	case float32:
		return uint64(in.(float32)), nil
	case float64:
		return uint64(in.(float64)), nil

	case uint8:
		return uint64(in.(uint8)), nil
	case uint16:
		return uint64(in.(uint16)), nil
	case uint32:
		return uint64(in.(uint32)), nil
	case uint:
		return uint64(in.(uint)), nil

	case int8:
		return uint64(in.(int8)), nil
	case int16:
		return uint64(in.(int16)), nil
	case int32:
		return uint64(in.(int32)), nil
	case int:
		return uint64(in.(int)), nil
	case int64:
		return uint64(in.(int64)), nil

	case []byte:
		return strconv.ParseUint(string(in.([]byte)), 10, 64)
	case string:
		return strconv.ParseUint(in.(string), 10, 64)
	}

	return uint64(0), fmt.Errorf("Uint64Err wrong value for '%+v' [keys: %+v]", in, debugKeys)
}
