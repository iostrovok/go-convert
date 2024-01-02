package convert_test

import (
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestSimple(t *testing.T) {
	iter, err := convert.Iterator([]interface{}{})
	Nil2(t, err)
	Nil2(t, iter.NextNotNil())

	iter, err = convert.Iterator([]string{})
	Nil2(t, err)
	Nil2(t, iter.NextNotNil())

	iter, err = convert.Iterator([]string{}, false)
	Nil2(t, err)
	Nil2(t, iter.NextNotNil())

	iter, err = convert.Iterator([]int{1})
	Nil2(t, err)
	Equal(t, iter.NextNotNil(), 1)
	Nil2(t, iter.NextNotNil())

	_, err = convert.Iterator("")
	NotNil2(t, err)
}

func TestLength(t *testing.T) {
	_, err := convert.Iterator([]interface{}{}, true)
	NotNil2(t, err)

	_, err = convert.Iterator([]string{}, true)
	NotNil2(t, err)
}

func TestIterator(t *testing.T) {
	data := []interface{}{
		"a",
		nil,
		1,
		"b",
	}

	iter, err := convert.Iterator(data)
	Nil2(t, err)

	Equal(t, iter.NextNotNil(), "a")
	Equal(t, iter.NextNotNil(), 1)
	Equal(t, iter.NextNotNil(), "b")
	Nil2(t, iter.NextNotNil())
}

func TestNextNotEmptyString(t *testing.T) {
	data := []interface{}{
		"a",
		nil,
		"",
		"b",
		"",
	}

	iter, err := convert.Iterator(data)
	Nil2(t, err)

	str := iter.NextNotEmptyString()
	Equal(t, str, "a")

	str = iter.NextNotEmptyString()
	Equal(t, str, "b")

	str = iter.NextNotEmptyString()
	Equal(t, str, "")

}

func TestCheckMapStringType(t *testing.T) {
	_, notFind := convert.CheckMapStringType(nil)
	EqualBool(t, notFind, false)

	a := map[string]interface{}{
		"string_a": 1,
	}
	check, find := convert.CheckMapStringType(a)

	EqualBool(t, find, true)
	Equal(t, len(check), len(a))
	for k, v := range a {
		Equal(t, check[k], v)
	}
}

func TestNextNotNilMapString(t *testing.T) {
	a := map[string]interface{}{
		"string_a": 1,
	}

	b := map[string]interface{}{
		"string_b1": 0,
		"string_b2": "super",
	}

	d := map[string]interface{}{
		"string_c1": "super",
		"string_c2": "puper",
	}

	data := []interface{}{
		a,
		b,
		nil,
		d,
	}

	iter, err := convert.Iterator(data)
	Nil2(t, err)

	a1, find := iter.NextNotNilMapString()
	EqualBool(t, find, true)
	for k, v := range a {
		Equal(t, a1[k], v)
	}

	b1, find := iter.NextNotNilMapString()
	EqualBool(t, find, true)
	for k, v := range b {
		Equal(t, b1[k], v)
	}

	d1, find := iter.NextNotNilMapString()
	EqualBool(t, find, true)
	Equal(t, len(d1), len(d))
	for k, v := range d {
		Equal(t, d1[k], v)
	}

	_, find = iter.NextNotNilMapString()
	EqualBool(t, find, false)
}
