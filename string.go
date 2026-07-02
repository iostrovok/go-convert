package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// isScalar reports whether in is one of the scalar types that String() can
// format via fmt.Sprintf("%s", ...): bool, all int/uint variants, float32/64,
// and string. Using a positive whitelist avoids panicking on a nil input and
// makes the intent explicit.
func isScalar(in any) bool {
	/*
		Goods types as scalar:
		   Bool Int8 Int Int16 Int32 Int64
		   Uint Uint8 Uint16 Uint32 Uint64
		   String Float32 Float64
	*/

	t := reflect.TypeOf(in)
	switch t.Kind() {
	case reflect.Bool, reflect.Int8, reflect.Int,
		reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// StringTS returns the string representation of in with leading and trailing
// whitespace removed. See String for conversion rules.
func StringTS(in any, debugKeys ...string) string {
	return strings.TrimSpace(String(in, debugKeys...))
}

// String converts in to its string representation.
// Scalar types (bool, numeric, string, []byte) are converted via strconv for
// performance. Non-scalar types and nil return "".
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

// ListOfStringsP ("List Of Strings Positive") behaves exactly like ListOfStringsPErr,
// but always returns a list (which may be empty) and never returns an error.
func ListOfStringsP(in any, debugKeys ...string) []string {
	if a, err := ListOfStringsPErr(in, false, debugKeys...); err == nil {
		return a
	}

	return []string{}
}

// ListOfStringsPErr ("List Of Strings Positive with Error") returns a list of
// non-empty strings converted from in. Empty strings are silently skipped.
// If checkLen is true and the resulting list is empty, an error is returned.
func ListOfStringsPErr(in any, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStrings(in, checkLen, SkipEmpty, debugKeys...)
}

// ListOfStringsStrictPErr ("List Of Strings Strict Positive with Error") returns a
// list of non-empty strings converted from in. If any element converts to an empty
// string, an error is returned immediately.
// If checkLen is true and the resulting list is empty, an error is also returned.
func ListOfStringsStrictPErr(in any, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStrings(in, checkLen, FallOnEmpty, debugKeys...)
}

// ListOfStrings converts in (which must be a slice) into a []string, including
// empty strings. It always returns a list (which may be empty) and never returns
// an error. See ListOfStringsErr for the error-returning variant.
func ListOfStrings(in any, debugKeys ...string) []string {
	if a, err := ListOfStringsErr(in, false, debugKeys...); err == nil {
		return a
	}

	return []string{}
}

// ListOfStringsErr converts in (which must be a slice) into a []string, including
// empty strings. If checkLen is true and the resulting list is empty, an error is
// returned. Optional debugKeys are embedded in error messages for tracing.
func ListOfStringsErr(in any, checkLen bool, debugKeys ...string) ([]string, error) {
	return _listOfStrings(in, checkLen, PassAll, debugKeys...)
}

// StringListRunner controls how _listOfStrings handles empty string elements.
type StringListRunner string

const (
	// PassAll includes every element in the result, even empty strings.
	PassAll StringListRunner = "pass_all"
	// FallOnEmpty returns an error as soon as an empty string element is encountered.
	FallOnEmpty StringListRunner = "fall_on_empty"
	// SkipEmpty silently omits empty string elements from the result.
	SkipEmpty StringListRunner = "skip_empty"
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
