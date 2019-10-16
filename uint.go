package convert

/*
	It's long code but quick! ))
*/

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Uint32(in interface{}, debugKeys ...string) uint32 {
	out, err := Uint32Err(in, debugKeys...)
	if err != nil {
		if len(debugKeys) > 0 {
			log.Printf("Uint64 Wrong value for '%+v' [keys: %+v], error: %v", in, debugKeys, err)
		}
		return 0
	}

	return out
}

func Uint32Err(in interface{}, debugKeys ...string) (uint32, error) {
	u64, err := Uint64Err(in, debugKeys...)
	if err != nil {
		return 0, err
	}

	if u64 > math.MaxUint32 {
		return 0, fmt.Errorf("Uint64Err Wrong value for '%+v'", in)
	}

	return uint32(u64), nil
}

func Uint64(in interface{}, debugKeys ...string) uint64 {

	out, err := Uint64Err(in, debugKeys...)
	if err != nil {
		if len(debugKeys) > 0 {
			log.Printf("Uint64 Wrong value for '%+v' [keys: %+v]", in, debugKeys)
		}
	}

	return out
}

func Uint64Err(in interface{}, debugKeys ...string) (uint64, error) {

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

	case string:
		return strconv.ParseUint(in.(string), 10, 64)
	}

	return uint64(0), fmt.Errorf("Int64Err Wrong value for '%+v' [keys: %+v]", in, debugKeys)
}
