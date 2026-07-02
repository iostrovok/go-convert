package convert

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// Float64 converts in to a float64. Returns 0 on error.
// Optional debugKeys are embedded in error messages for tracing.
func Float64(in any, debugKeys ...string) float64 {
	f64, err := Float64Err(in, debugKeys...)
	if err != nil {
		return 0
	}
	return f64
}

// Float32 converts in to a float32. Returns 0 on error.
// Optional debugKeys are embedded in error messages for tracing.
func Float32(in any, debugKeys ...string) float32 {
	f, err := Float32Err(in, debugKeys...)
	if err != nil {
		return 0
	}
	return f
}

// Float64Err converts in to a float64. It delegates to BaseFloat64Err and
// additionally rejects ±Inf results, returning 0 in that case.
// Optional debugKeys are embedded in error messages for tracing.
func Float64Err(in any, debugKeys ...string) (float64, error) {
	f, err := BaseFloat64Err(in, debugKeys...)
	if err != nil {
		return 0, err
	}

	if math.IsInf(f, -1) || math.IsInf(f, -1) {
		return 0, err
	}

	return f, nil
}

// BaseFloat64Err converts in to a float64 without filtering ±Inf.
// Accepted types: all numeric types, bool (true→1.0, false→0.0), string, and []byte.
// Returns an error for nil or unconvertible values.
// Optional debugKeys are embedded in error messages for tracing.
func BaseFloat64Err(in any, debugKeys ...string) (float64, error) {

	if in == nil {
		return .0, fmt.Errorf("Float64Err null value for '%+v' [keys: %+v]", in, debugKeys)
	}

	switch in.(type) {

	case reflect.Kind:
		return 0, fmt.Errorf("Float64Err wrong value for '%+v' [keys: %+v]", in, debugKeys)

	case float64:
		return in.(float64), nil
	case float32:
		return float64(in.(float32)), nil

	case bool:
		if in.(bool) {
			return 1.0, nil
		}
		return 0, nil

	case uint8:
		return float64(in.(uint8)), nil
	case uint16:
		return float64(in.(uint16)), nil
	case uint32:
		return float64(in.(uint32)), nil
	case uint:
		return float64(in.(uint)), nil
	case uint64:
		return float64(in.(uint64)), nil

	case int8:
		return float64(in.(int8)), nil
	case int16:
		return float64(in.(int16)), nil
	case int32:
		return float64(in.(int32)), nil
	case int:
		return float64(in.(int)), nil
	case int64:
		return float64(in.(int64)), nil
	case []byte:
		f, err := strconv.ParseFloat(string(in.([]byte)), 64)
		if err != nil {
			return 0, fmt.Errorf("Float64Err wrong common value for '%+v' [keys: %+v]", in, debugKeys)
		}
		return f, nil
	case string:
		f, err := strconv.ParseFloat(in.(string), 64)
		if err != nil {
			return 0, fmt.Errorf("Float64Err wrong common value for '%+v' [keys: %+v]", in, debugKeys)
		}
		return f, nil

	}

	return 0, fmt.Errorf("Float64Err wrong unreachable value for '%+v' [keys: %+v]", in, debugKeys)
}

// Float32Err converts in to a float32. Returns an error if the conversion fails
// or if a float64 input value exceeds ±math.MaxFloat32.
// Optional debugKeys are embedded in error messages for tracing.
func Float32Err(in any, debugKeys ...string) (float32, error) {

	if in == nil {
		return .0, fmt.Errorf("Float32Err null value for '%+v' [keys: %+v]", in, debugKeys)
	}

	switch in.(type) {

	case reflect.Kind:
		return 0, fmt.Errorf("Float32Err wrong value for '%+v' [keys: %+v]", in, debugKeys)

	case float64:
		f64 := in.(float64)
		if f64 > math.MaxFloat32 || f64 < -1.0*math.MaxFloat32 {
			return 0, fmt.Errorf("Float32Err Wrong value for '%+v' [keys: %+v]", in, debugKeys)
		}
		return float32(f64), nil
	case float32:
		return in.(float32), nil

	case bool:
		if in.(bool) {
			return 1.0, nil
		}
		return 0, nil
	case uint8:
		return float32(in.(uint8)), nil
	case uint16:
		return float32(in.(uint16)), nil
	case uint32:
		return float32(in.(uint32)), nil
	case uint:
		return float32(in.(uint)), nil
	case uint64:
		return float32(in.(uint64)), nil
	case int8:
		return float32(in.(int8)), nil
	case int16:
		return float32(in.(int16)), nil
	case int32:
		return float32(in.(int32)), nil
	case int:
		return float32(in.(int)), nil
	case int64:
		return float32(in.(int64)), nil
	case []byte:
		f, err := strconv.ParseFloat(string(in.([]byte)), 32)
		if err != nil {
			return 0, fmt.Errorf("Float32Err wrong common value for '%+v' [keys: %+v]", in, debugKeys)
		}
		return float32(f), nil
	case string:
		f, err := strconv.ParseFloat(in.(string), 32)
		if err != nil {
			return 0, fmt.Errorf("Float32Err wrong common value for '%+v' [keys: %+v]", in, debugKeys)
		}
		return float32(f), nil

	}

	return 0, fmt.Errorf("Float32Err wrong unreachable value for '%+v' [keys: %+v]", in, debugKeys)
}
