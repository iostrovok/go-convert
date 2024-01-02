package convert_test

import (
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestBool_new(t *testing.T) {
	EqualBool(t, convert.Bool(1), true)
	EqualBool(t, convert.Bool(1), true)
	EqualBool(t, convert.Bool("1"), true)
	EqualBool(t, convert.Bool("-1"), true)
	EqualBool(t, convert.Bool(true), true)
	EqualBool(t, convert.Bool("True"), true)
	EqualBool(t, convert.Bool("true"), true)
	EqualBool(t, convert.Bool("t"), true)
	EqualBool(t, convert.Bool(0.1), true)

	EqualBool(t, convert.Bool("0"), false)
	EqualBool(t, convert.Bool(false), false)
	EqualBool(t, convert.Bool(""), false)
	EqualBool(t, convert.Bool("false"), false)
	EqualBool(t, convert.Bool("False"), false)
	EqualBool(t, convert.Bool("f"), false)
	EqualBool(t, convert.Bool(0), false)
	EqualBool(t, convert.Bool(0.0), false)
}

func TestBool(t *testing.T) {
	EqualBool(t, convert.Bool(1), true)
	EqualBool(t, convert.Bool("1"), true)
	EqualBool(t, convert.Bool("-1"), true)
	EqualBool(t, convert.Bool(true), true)
	EqualBool(t, convert.Bool("True"), true)
	EqualBool(t, convert.Bool("true"), true)
	EqualBool(t, convert.Bool("t"), true)
	EqualBool(t, convert.Bool(0.1), true)

	EqualBool(t, convert.Bool("0"), false)
	EqualBool(t, convert.Bool(false), false)
	EqualBool(t, convert.Bool(""), false)
	EqualBool(t, convert.Bool("false"), false)
	EqualBool(t, convert.Bool("False"), false)
	EqualBool(t, convert.Bool("f"), false)
	EqualBool(t, convert.Bool(0), false)
	EqualBool(t, convert.Bool(0.0), false)
}

func TestBoolErr(t *testing.T) {
	a, err := convert.BoolErr("1")
	Nil2(t, err)
	EqualBool(t, a, true)
}

func TestBool_ArrayMap(t *testing.T) {
	EqualBool(t, convert.Bool([]interface{}{}), false)
	EqualBool(t, convert.Bool(map[string]interface{}{}), false)
	EqualBool(t, convert.Bool(map[int]interface{}{}), false)

	EqualBool(t, convert.Bool([]string{}), false)
	EqualBool(t, convert.Bool(map[string]string{}), false)
	EqualBool(t, convert.Bool(map[int]string{}), false)

	EqualBool(t, convert.Bool([]int{}), false)
	EqualBool(t, convert.Bool(map[string]int{}), false)
	EqualBool(t, convert.Bool(map[int]int{}), false)

	EqualBool(t, convert.Bool([]interface{}{1}), true)
	EqualBool(t, convert.Bool(map[string]interface{}{"1": 1}), true)
	EqualBool(t, convert.Bool(map[int]interface{}{1: "asdasdasd"}), true)

	EqualBool(t, convert.Bool([]string{"1"}), true)
	EqualBool(t, convert.Bool(map[string]string{"1": "1"}), true)
	EqualBool(t, convert.Bool(map[int]string{1: "1"}), true)

	EqualBool(t, convert.Bool([]int{1}), true)
	EqualBool(t, convert.Bool(map[string]int{"1": 1}), true)
	EqualBool(t, convert.Bool(map[int]int{1: 1}), true)
}
