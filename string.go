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

func isScalar(in any) bool {
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

// StringTS returns spaces trimmed string
func StringTS(in any, debugKeys ...string) string {
	return strings.TrimSpace(String(in, debugKeys...))
}

func String(in any, _ ...string) string {
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

// ListOfStringsP ("List Of Strings Positive")
// is exactly the same as ListOfStringsPErr, but always returns a list (can be empty).
func ListOfStringsP(in any, debugKeys ...string) []string {
	if a, err := ListOfStringsPErr(in, false, debugKeys...); err == nil {
		return a
	}

	return []string{}
}

// ListOfStringsPErr ("List Of Strings Positive with Error")
// returns a list of no-empty strings extracted from the input interface.
// If the list contains an empty string, the function skips it.
// If checkLen is true and the result list is empty, the function returns an error.
func ListOfStringsPErr(in any, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStrings(in, checkLen, SkipEmpty, debugKeys...)
}

// ListOfStringsStrictPErr ("List Of Strings Strict Positive with Error")
// returns a list of no-empty strings extracted from the input interface.
// If the list has empty string function returns the error.
func ListOfStringsStrictPErr(in any, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStrings(in, checkLen, FallOnEmpty, debugKeys...)
}

func ListOfStrings(in any, debugKeys ...string) []string {
	if a, err := ListOfStringsErr(in, false, debugKeys...); err == nil {
		return a
	}

	return []string{}
}

func ListOfStringsErr(in any, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStrings(in, checkLen, PassAll, debugKeys...)
}

type StringListRunner string

const (
	PassAll     StringListRunner = "pass_all"
	FallOnEmpty StringListRunner = "fall_on_empty"
	SkipEmpty   StringListRunner = "skip_empty"
)

func _listOfStrings(in any, checkLen bool, howToRun StringListRunner, debugKeys ...string) ([]string, error) {
	debugKey := ""
	if len(debugKeys) > 0 {
		debugKey = debugKeys[0]
	}

	if in == nil {
		return nil, fmt.Errorf("null value for '%s'", debugKey)
	}

	it, err := Iterator(in, checkLen)
	if err != nil {
		return nil, fmt.Errorf("wrong iterator value for '%s'", debugKey)
	}

	out := make([]string, 0)
	for i := 0; i < it.Len(); i++ {
		next := it.NextNotNil()
		if next == nil {
			return nil, fmt.Errorf("wrong next value for '%+v' [debugKey: %s]", next, debugKey)
		}

		s := String(next)
		switch howToRun {
		case FallOnEmpty:
			if s == "" {
				return nil, fmt.Errorf("empty string value for '%+v' [debugKey: %s]", next, debugKey)
			}
			out = append(out, s)
		case SkipEmpty:
			if s != "" {
				out = append(out, s)
			}
		default:
			// PassAll
			out = append(out, s)
		}
	}

	if checkLen && len(out) == 0 {
		return out, fmt.Errorf("the result list is empty for '%s'", debugKey)
	}

	return out, nil
}
