package convert_test

import (
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestMapIteratorEmpty(t *testing.T) {
	iter, err := convert.MapIterator(map[string]any{})
	Nil2(t, err)
	EqualBool(t, iter.HasNext(), false)
	a, b := iter.Next()
	Nil2(t, a)
	Nil2(t, b)

	iter, err = convert.MapIterator(map[string]int32{})
	Nil2(t, err)
	EqualBool(t, iter.HasNext(), false)
	a, b = iter.Next()
	Nil2(t, a)
	Nil2(t, b)

	iter, err = convert.MapIterator(map[int][]int32{})
	Nil2(t, err)
	EqualBool(t, iter.HasNext(), false)
	a, b = iter.Next()
	Nil2(t, a)
	Nil2(t, b)

	iter, err = convert.MapIterator(map[string]string{})
	Nil2(t, err)
	EqualBool(t, iter.HasNext(), false)
	a, b = iter.Next()
	Nil2(t, a)
	Nil2(t, b)
}

func TestMapIteratorLength(t *testing.T) {
	_, err := convert.MapIterator(map[string]any{}, true)
	NotNil2(t, err)

	_, err = convert.MapIterator(map[int][]int32{}, true)
	NotNil2(t, err)
}

func TestMapIteratorRun_1(t *testing.T) {
	data := map[string]any{
		"a": "a",
		"b": nil,
		"c": 1,
		"d": "b",
	}

	iter, err := convert.MapIterator(data)
	Nil2(t, err)
	EqualBool(t, iter.HasNext(), true)

	for {
		if !iter.HasNext() {
			break
		}

		k, v := iter.Next()

		val, find := data[convert.String(k)]
		EqualBool(t, find, true)
		Equal(t, val, v)
	}
}
