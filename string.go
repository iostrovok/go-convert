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

// String Trim spaces
func StringTS(in interface{}, debugKeys ...string) string {
	return strings.TrimSpace(String(in, debugKeys...))
}

func String(in interface{}, debugKeys ...string) string {

	if in == nil {
		return ""
	}

	switch in.(type) {
	case []byte:
		return string(in.([]byte))
	case string:
		return in.(string)
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

func ListOfStringsP(in interface{}, checkLen bool, debugKeys ...string) []string {
	a, err := _listOfStringsErr(in, checkLen, false, true, debugKeys...)
	if err == nil {
		return a
	}
	return []string{}
}

func ListOfStrings(in interface{}, checkLen bool, debugKeys ...string) []string {
	a, err := _listOfStringsErr(in, checkLen, false, false, debugKeys...)
	if err == nil {
		return a
	}
	return []string{}
}

func ListOfStringsPErr(in interface{}, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStringsErr(in, checkLen, true, false, debugKeys...)
}

func ListOfStringsErr(in interface{}, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStringsErr(in, checkLen, false, false, debugKeys...)
}

func _listOfStringsErr(in interface{}, checkLen, checkEmpty, missEmpty bool, debugKeys ...string) ([]string, error) {
	debugKey := ""
	if len(debugKeys) > 0 {
		debugKey = debugKeys[0]
	}

	if in == nil {
		return nil, fmt.Errorf("ListOfStringsErr null value for '%+v' [debugKey: %s]", in, debugKey)
	}

	it, err := Iterator(in, checkLen)
	if err != nil {
		return nil, fmt.Errorf("ListOfStringsErr wrong iterator value for '%+v' [debugKey: %s]", in, debugKey)
	}

	out := make([]string, 0)
	for i := 0; i < it.Len(); i++ {
		s := it.NextNotNil()
		if s == nil {
			return nil, fmt.Errorf("ListOfStringsErr wrong next value for '%+v' [debugKey: %s]", in, debugKey)
		}

		if checkEmpty {
			s := String(s, debugKey+"/"+strconv.Itoa(len(out)))
			if s == "" {
				return nil, fmt.Errorf("ListOfStringsErr empty string value for '%+v' [debugKey: %s]", in, debugKey)
			}
			out = append(out, s)
		} else if missEmpty {
			s := String(s, debugKey+"/"+strconv.Itoa(len(out)))
			if s != "" {
				out = append(out, s)
			}
		} else {
			out = append(out, String(s, debugKey+"/"+strconv.Itoa(len(out))))
		}
	}

	return out, nil
}
