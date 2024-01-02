package convert_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/iostrovok/go-convert"
)

type oneFloat64Test struct {
	val     interface{}
	result  float64
	isError bool
}

func TestFloat64Err(t *testing.T) {
	testList := []oneFloat64Test{
		{"", 0, true},
		{"wqeqwe", 0, true},
		{"3049", 3049, false},

		{nil, 0, true},

		{math.Inf(1), math.Inf(1), false},
		{math.Inf(-1), math.Inf(-1), false},
		{math.MaxFloat64, math.MaxFloat64, false},
		{math.MaxFloat32, math.MaxFloat32, false},
		{math.MaxInt64, math.MaxInt64, false},
		{math.MaxInt32, math.MaxInt32, false},
		{reflect.Invalid, 0, true},

		{true, 1, false},
		{false, 0, false},

		{float32(3049.24343), float64(float32(3049.24343)), false},
		{3049.24343, 3049.24343, false},

		{3049, 3049, false},
		{int8(32), 32, false},
		{int16(32), 32, false},
		{int32(3049), 3049, false},
		{int64(3049), 3049, false},

		{uint(32), 32, false},
		{uint32(32), 32, false},
		{uint8(32), 32, false},
		{uint16(32), 32, false},
		{uint64(32), 32, false},
	}

	for _, one := range testList {
		f64, err := convert.BaseFloat64Err(one.val)
		if one.isError {
			NotNil2(t, err)
		} else {
			Nil2(t, err)
			Equal(t, f64, one.result)
		}
	}
}

func TestFloat64(t *testing.T) {
	testList := []oneFloat64Test{
		{"", 0, false},
		{nil, 0, false},
		{"3049", 3049, false},
		{int32(3049), 3049, false},
		{int64(3049), 3049, false},
		{3049, 3049, false},
		{3049.24343, 3049.24343, false},
		{math.MaxFloat64, math.MaxFloat64, false},
		{math.MaxFloat32, math.MaxFloat32, false},
		{math.MaxInt64, math.MaxInt64, false},
		{math.MaxInt32, math.MaxInt32, false},
		{reflect.Invalid, 0, false},
	}

	for _, one := range testList {
		Equal(t, convert.Float64(one.val), one.result)
	}
}

type oneFloat32Test struct {
	val     interface{}
	result  float32
	isError bool
}

func TestFloat32Err(t *testing.T) {

	testList := []oneFloat32Test{
		{math.MaxFloat64, 0, true},
		{math.MaxFloat32, math.MaxFloat32, false},
		{math.MaxInt64, float32(math.MaxInt64), false},
		{reflect.Invalid, 0, true},
		{"", 0, true},
		{nil, 0, true},
		{"3049", 3049, false},
		{int32(3049), 3049, false},
		{int64(3049), 3049, false},
		{3049, 3049, false},
		{3049.24343, 3049.24343, false},
	}

	for _, one := range testList {
		f32, err := convert.Float32Err(one.val)
		if one.isError {
			NotNil2(t, err)
		} else {
			Nil2(t, err)
			Equal(t, f32, one.result)
		}
	}
}

func TestFloat32(t *testing.T) {
	testList := []oneFloat32Test{
		{"", 0, false},
		{"wqeqwe", 0, true},
		{"3049", 3049, false},

		{nil, 0, false},

		{math.MaxFloat64, 0, false},
		{math.MaxFloat32, math.MaxFloat32, false},
		{math.MaxInt64, float32(math.MaxInt64), false},
		{reflect.Invalid, 0, false},

		{true, 1, false},
		{false, 0, false},

		{float32(3049.24343), 3049.24343, false},
		{3049.24343, 3049.24343, false},

		{3049, 3049, false},
		{int8(32), 32, false},
		{int16(32), 32, false},
		{int32(3049), 3049, false},
		{int64(3049), 3049, false},

		{uint(32), 32, false},
		{uint32(32), 32, false},
		{uint8(32), 32, false},
		{uint16(32), 32, false},
		{uint64(32), 32, false},
	}

	for _, one := range testList {
		Equal(t, convert.Float32(one.val), one.result)
	}
}
