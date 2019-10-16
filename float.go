package convert

import (
	"fmt"
	"math"
	"strconv"
)

func Float64(in interface{}, debugKeys ...string) float64 {
	f64, err := Float64Err(in, debugKeys...)
	if err != nil {
		return 0
	}
	return f64
}

func Float32(in interface{}, debugKeys ...string) float32 {
	f, err := Float32Err(in, debugKeys...)
	if err != nil {
		return 0
	}
	return f
}

func Float32Err(in interface{}, debugKeys ...string) (float32, error) {
	f64, err := Float64Err(in, debugKeys...)

	if err != nil {
		return 0, err
	}

	if f64 > math.MaxFloat32 || f64 < -1.0*math.MaxFloat32 {
		return 0, fmt.Errorf("Float32Err Wrong value for '%+v' [keys: %+v]", in, debugKeys)
	}

	return float32(f64), err
}

func Float64Err(in interface{}, debugKeys ...string) (float64, error) {
	f, err := _float64Err(in, debugKeys...)
	if err != nil {
		return 0, err
	}

	if math.IsInf(f, -1) || math.IsInf(f, -1) {
		return 0, err
	}

	return f, nil
}

func _float64Err(in interface{}, debugKeys ...string) (float64, error) {

	if in == nil {
		return .0, fmt.Errorf("Float64Err Wrong value for '%+v' [keys: %+v]", in, debugKeys)
	}

	switch in.(type) {

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
	case string:
		f, err := strconv.ParseFloat(in.(string), 64)
		if err != nil {
			return 0, fmt.Errorf("Float64Err Wrong value for '%+v' [keys: %+v]", in, debugKeys)
		}
		return f, nil

	}

	return 0, nil
}
