package convert

import (
	"fmt"
	"math"
	"strconv"
)

// Int32 converts in to an int32. Returns 0 on error or if the value overflows int32.
// Optional debugKeys are embedded in error messages for tracing.
func Int32(in any, debugKeys ...string) int32 {
	i64, err := Int64Err(in, debugKeys...)
	if err != nil || i64 > math.MaxInt32 || i64 < math.MinInt32 {
		return int32(0)
	}

	return int32(i64)
}

// Int converts in to an int. Returns 0 on error.
// Optional debugKeys are embedded in error messages for tracing.
func Int(in any, debugKeys ...string) int {
	i64, err := Int64Err(in, debugKeys...)
	if err != nil {
		return 0
	}
	return int(i64)
}

// Int64 converts in to an int64. Returns 0 on error.
// Optional debugKeys are embedded in error messages for tracing.
func Int64(in any, debugKeys ...string) int64 {
	i64, err := Int64Err(in, debugKeys...)
	if err != nil {
		return int64(0)
	}
	return i64
}

// Int32Err converts in to an int32. Returns an error if the conversion fails
// or if the value overflows the int32 range (math.MinInt32…math.MaxInt32).
// Optional debugKeys are embedded in error messages for tracing.
func Int32Err(in any, debugKeys ...string) (int32, error) {
	i64, err := Int64Err(in, debugKeys...)
	if err != nil {
		return int32(0), err
	}

	if i64 > math.MaxInt32 || i64 < math.MinInt32 {
		return int32(0), fmt.Errorf("Int32Err wrong value for '%+v' [keys: %+v]", in, debugKeys)
	}

	return int32(i64), err
}

// Int64Err converts in to an int64. Accepted types: all integer and float types,
// bool (true→1, false→0), string, and []byte.
// Returns an error for nil or unconvertible values.
// Optional debugKeys are embedded in error messages for tracing.
func Int64Err(in any, debugKeys ...string) (int64, error) {
	if in == nil {
		return int64(0), fmt.Errorf("Int64Err null value for '%+v' [keys: %+v]", in, debugKeys)
	}

	switch in.(type) {
	case int64:
		return in.(int64), nil

	case bool:
		if in.(bool) {
			return int64(1), nil
		}
		return int64(0), nil

	case float32:
		return int64(in.(float32)), nil
	case float64:
		return int64(in.(float64)), nil

	case uint8:
		return int64(in.(uint8)), nil
	case uint16:
		return int64(in.(uint16)), nil
	case uint32:
		return int64(in.(uint32)), nil
	case uint:
		return int64(in.(uint)), nil
	case uint64:
		return int64(in.(uint64)), nil

	case int8:
		return int64(in.(int8)), nil
	case int16:
		return int64(in.(int16)), nil
	case int:
		return int64(in.(int)), nil
	case int32:
		return int64(in.(int32)), nil
	case []byte:
		return strconv.ParseInt(string(in.([]byte)), 10, 64)
	case string:
		return strconv.ParseInt(in.(string), 10, 64)
	}

	return int64(0), fmt.Errorf("Int64Err wrong value for '%+v' [keys: %+v]", in, debugKeys)
}

// ListOfInt32Err converts in (which must be a slice) into a []int32.
// Each element is converted via Int64Err and then range-checked against int32 bounds.
// If checkLen is true and the input slice is empty, an error is returned.
// Optional debugKeys are embedded in error messages for tracing.
func ListOfInt32Err(in any, checkLen bool, debugKeys ...string) ([]int32, error) {

	debugKey := ""
	if len(debugKeys) > 0 {
		debugKey = debugKeys[0]
	}

	int32List, err := ListOfInt64Err(in, checkLen, debugKeys...)
	if err != nil {
		return nil, fmt.Errorf("ListOfInt32Err wrong Int64 value for '%+v' [debugKey: %s]", in, debugKey)
	}

	out := make([]int32, len(int32List), len(int32List))
	for i, v := range int32List {
		if v > math.MaxInt32 || v < math.MinInt32 {
			return nil, fmt.Errorf("ListOfInt32Err Wrong value for '%+v' [keys: %s]", in, debugKey+"/"+strconv.Itoa(i))
		}
		out[i] = int32(v)
	}

	return out, nil
}

// ListOfInt64Err converts in (which must be a slice) into a []int64.
// Each element is converted via Int64Err.
// If checkLen is true and the input slice is empty, an error is returned.
// Optional debugKeys are embedded in error messages for tracing.
func ListOfInt64Err(in any, checkLen bool, debugKeys ...string) ([]int64, error) {

	debugKey := ""
	if len(debugKeys) > 0 {
		debugKey = debugKeys[0]
	}

	if in == nil {
		return nil, fmt.Errorf("ListOfInt64Err null value for '%+v' [debugKey: %s]", in, debugKey)
	}

	it, err := Iterator(in, checkLen)
	if err != nil {
		return nil, fmt.Errorf("ListOfInt64Err wrong iterator value for '%+v' [debugKey: %s]", in, debugKey)
	}

	out := make([]int64, 0)

	for i := 0; i < it.Len(); i++ {
		s := it.NextNotNil()
		if s == nil {
			return nil, fmt.Errorf("ListOfInt64Err wrong next value for '%+v' [debugKey: %s]", in, debugKey)
		}

		res, err := Int64Err(s, debugKey+"/"+strconv.Itoa(len(out)))
		if err != nil {
			return nil, fmt.Errorf("ListOfInt64Err wrong value for '%+v' [keys: %s]", in, debugKey+"/"+strconv.Itoa(len(out)))
		}

		out = append(out, res)
	}

	return out, nil
}
