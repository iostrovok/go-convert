package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
	It's long code but quick! ))
*/

func isScalar(in interface{}) bool {
	/*
		Goods types:
		   Bool Int8 Int Int16 Int32 Int64
		   Uint Uint8 Uint16 Uint32 Uint64
		   String Float32 Float64
	*/

	t := reflect.TypeOf(in)
	switch t.Kind() {
	case reflect.Invalid, reflect.Uintptr, reflect.Complex64, reflect.Complex128, reflect.Struct, reflect.UnsafePointer:
		return false
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return false
	}

	return true
}

func String(in interface{}, debugKeys ...string) string {

	if in == nil {
		return ""
	}

	switch in.(type) {
	case []byte:
		return strings.TrimSpace(in.(string))
	case string:
		return strings.TrimSpace(in.(string))
	case bool:
		return strconv.FormatBool(in.(bool))

	case float32:
		return strconv.FormatFloat(float64(in.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(in.(float64), 'f', -1, 64)

	case int8:
		return strconv.Itoa(int(in.(int8)))
	case int16:
		return strconv.Itoa(int(in.(int16)))
	case int32:
		return strconv.Itoa(int(in.(int32)))
	case int:
		return strconv.Itoa(in.(int))
	case int64:
		return strconv.FormatInt(in.(int64), 10)

	case uint8:
		return strconv.FormatUint(uint64(in.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(in.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(in.(uint32)), 10)
	case uint:
		return strconv.FormatUint(uint64(in.(uint)), 10)
	case uint64:
		return strconv.FormatUint(in.(uint64), 10)
	}

	if isScalar(in) {
		return strings.TrimSpace(fmt.Sprintf("%s", in))
	}

	return ""
}
