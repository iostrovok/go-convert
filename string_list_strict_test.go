package convert_test

import (
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestListOfStringsStrictP_Empty_check_length(t *testing.T) {
	_, e := convert.ListOfStringsStrictPErr(nil, true)
	NotNil2(t, e)

	_, e = convert.ListOfStringsStrictPErr([]string{}, true)
	NotNil2(t, e)

	_, e = convert.ListOfStringsStrictPErr([]string{""}, true)
	NotNil2(t, e)
}

func TestListOfStringsStrictP_Empty_no_check_length(t *testing.T) {
	_, e := convert.ListOfStringsStrictPErr(nil, false)
	NotNil2(t, e)

	a, e := convert.ListOfStringsStrictPErr([]string{}, false)
	Nil2(t, e)
	EqualStringsArray(t, a, []string{})

	_, e = convert.ListOfStringsStrictPErr([]string{""}, false)
	NotNil2(t, e)
}

func TestListOfStringsStrictP_Full(t *testing.T) {
	mListFull := []string{"10", "20", "40"}
	a, e := convert.ListOfStringsStrictPErr(mListFull, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mListFull)

	a, e = convert.ListOfStringsStrictPErr([]any{10, "20", "40"}, true)
	Nil2(t, e)
	EqualStringsArray(t, a, mListFull)

	a, e = convert.ListOfStringsStrictPErr(mListFull, true)
	Nil2(t, e)
	EqualStringsArray(t, a, mListFull)

	a, e = convert.ListOfStringsStrictPErr([]string{"10", "20", "", "40"}, false)
	NotNil2(t, e)

	a, e = convert.ListOfStringsStrictPErr([]string{"10", "20", "", "40"}, true)
	NotNil2(t, e)

	a, e = convert.ListOfStringsStrictPErr([]string{"10", "20", "30", ""}, false)
	NotNil2(t, e)

	a, e = convert.ListOfStringsStrictPErr([]string{"10", "20", "30", ""}, true)
	NotNil2(t, e)
}
