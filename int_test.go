package convert_test

import (
	"math"
	"testing"

	"github.com/iostrovok/go-convert"
)

var testCases = []any{
	"3049", int32(3049), int64(3049), 3049, float32(3049.24343), float64(3049.24343),
}

func TestInt64(t *testing.T) {
	Equal(t, convert.Int64(""), int64(0))
	Equal(t, convert.Int64("wqeqwe"), int64(0))
	Equal(t, convert.Int64(true), int64(1))
	for _, v := range testCases {
		Equal(t, convert.Int64(v), int64(3049))
	}
}

func TestInt32(t *testing.T) {
	Equal(t, convert.Int32(""), int32(0))
	Equal(t, convert.Int32("wqeqwe"), int32(0))
	Equal(t, convert.Int32(true), int32(1))
	for _, v := range testCases {
		Equal(t, convert.Int32(v), int32(3049))
	}
}

func TestListOfInt64Err(t *testing.T) {
	res, err := convert.ListOfInt64Err([]any{
		"1", 12, 45, 123412323, -1, 0,
	}, false)

	Nil2(t, err)
	check := []int64{1, 12, 45, 123412323, -1, 0}
	Equal(t, len(res), len(check))
	for i, v := range check {
		Equal(t, res[i], v)
	}
}

func TestListOfInt64Err2(t *testing.T) {
	_, err := convert.ListOfInt64Err([]any{
		"1", 12, 45, nil, -1, 0,
	}, false)
	NotNil2(t, err)

	_, err = convert.ListOfInt64Err([]any{
		"", 12, 45, -1, 0,
	}, false)
	NotNil2(t, err)
}

func TestListOfInt64ErrEmpty(t *testing.T) {
	_, err := convert.ListOfInt64Err([]any{}, false)
	Nil2(t, err)

	_, err = convert.ListOfInt64Err([]any{}, true)
	NotNil2(t, err)
}

func TestListOfInt32Err(t *testing.T) {
	res, err := convert.ListOfInt32Err([]any{
		"1", 12, 45, 123412323, -1, 0,
	}, false)

	Nil2(t, err)
	check := []int32{1, 12, 45, 123412323, -1, 0}
	Equal(t, len(res), len(check))
	for i, v := range check {
		Equal(t, res[i], v)
	}
}

func TestListOfInt32Err2(t *testing.T) {
	_, err := convert.ListOfInt32Err([]any{
		"1", 12, 45, nil, -1, 0,
	}, false)
	NotNil2(t, err)

	_, err = convert.ListOfInt32Err([]any{
		"", 12, 45, -1, 0,
	}, false)
	NotNil2(t, err)

	_, err = convert.ListOfInt32Err([]any{
		"1", 12, 45, 123412323, -1, 0, math.MaxInt64,
	}, false)
	NotNil2(t, err)

	_, err = convert.ListOfInt32Err([]any{
		"1", 12, 45, 123412323, -1, 0, math.MinInt64,
	}, false)
	NotNil2(t, err)
}

func TestListOfInt32ErrEmpty(t *testing.T) {
	_, err := convert.ListOfInt32Err([]any{}, false)
	Nil2(t, err)

	_, err = convert.ListOfInt32Err([]any{}, true)
	NotNil2(t, err)
}
