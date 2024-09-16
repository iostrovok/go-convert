package convert_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestUint64(t *testing.T) {
	Equal(t, convert.Uint64(""), uint64(0))
	Equal(t, convert.Uint64(nil), uint64(0))
	Equal(t, convert.Uint64("3049"), uint64(3049))
	Equal(t, convert.Uint64(int32(3049)), uint64(3049))
	Equal(t, convert.Uint64(int64(3049)), uint64(3049))
	Equal(t, convert.Uint64(3049), uint64(3049))
	Equal(t, convert.Uint64(3049.24343), uint64(3049))
	Equal(t, convert.Uint64(math.MaxUint32), uint64(math.MaxUint32))
	Equal(t, convert.Uint64(math.MaxInt64), uint64(math.MaxInt64))
	Equal(t, convert.Uint64(reflect.Invalid), uint64(0))
}

func TestConvertUint32(t *testing.T) {
	Equal(t, convert.Uint32(""), uint32(0))
	Equal(t, convert.Uint32(nil), uint32(0))
	Equal(t, convert.Uint32("3049"), uint32(3049))
	Equal(t, convert.Uint32(int32(3049)), uint32(3049))
	Equal(t, convert.Uint32(int64(3049)), uint32(3049))
	Equal(t, convert.Uint32(3049), uint32(3049))
	Equal(t, convert.Uint32(3049.24343), uint32(3049))
	Equal(t, convert.Uint32(3032432434323449.24343), uint32(0))
	Equal(t, convert.Uint32(math.MaxUint32), uint32(math.MaxUint32))
	Equal(t, convert.Uint32(math.MaxInt64), uint32(0))
	Equal(t, convert.Uint32(reflect.Invalid), uint32(0))
}
