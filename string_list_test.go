package convert_test

import (
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestListOfStringsP(t *testing.T) {
	mList := make([]string, 0)
	a := convert.ListOfStringsP(mList)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStringsP(nil)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStringsP([]string{"10", "20", "40", ""})
	EqualStringsArray(t, a, []string{"10", "20", "40"})
}

func TestListOfStringsPErr_checkLen(t *testing.T) {
	mList := make([]string, 0)
	_, e := convert.ListOfStringsPErr(mList, true)
	NotNil2(t, e)

	_, e = convert.ListOfStringsPErr(nil, true)
	NotNil2(t, e)

	mList2 := []string{"10", "20", "40", ""}
	a, e := convert.ListOfStringsPErr(mList2, true)
	Nil2(t, e)
	EqualStringsArray(t, a, []string{"10", "20", "40"})

	a, e = convert.ListOfStringsPErr(mList2, true)
	Nil2(t, e)
	EqualStringsArray(t, a, []string{"10", "20", "40"})

	_, e = convert.ListOfStringsPErr([]string{"", "", ""}, true)
	NotNil2(t, e)
}

func TestListOfStringsPErr_no_checkLen(t *testing.T) {
	mList := make([]string, 0)

	a, e := convert.ListOfStringsPErr(mList, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mList)

	mListFull := []string{"10", "20", "40"}
	a, e = convert.ListOfStringsPErr(mListFull, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	_, e = convert.ListOfStringsPErr(mListEmpty, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mListFull)
}

func TestListOfStringsErr(t *testing.T) {
	mList := make([]string, 0)

	a, e := convert.ListOfStringsErr(mList, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mList)

	_, e = convert.ListOfStringsErr(mList, true)
	NotNil2(t, e)

	_, e = convert.ListOfStringsErr(nil, true)
	NotNil2(t, e)

	mListFull := []string{"10", "20", "40"}
	a, e = convert.ListOfStringsErr(mListFull, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a, e = convert.ListOfStringsErr(mListEmpty, false)
	Nil2(t, e)
	EqualStringsArray(t, a, mListEmpty)
}

func TestListOfStringsP_Empty(t *testing.T) {
	mList := make([]string, 0)

	a := convert.ListOfStringsP(nil)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStringsP(mList)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStringsP(mList)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStringsP([]string{""})
	EqualStringsArray(t, a, mList)
}

func TestListOfStringsP_Full(t *testing.T) {
	mListFull := []string{"10", "20", "40"}
	a := convert.ListOfStringsP(mListFull)
	EqualStringsArray(t, a, mListFull)

	mListEmpty := []string{"10", "20", "", "40"}
	a = convert.ListOfStringsP(mListEmpty)
	EqualStringsArray(t, a, []string{"10", "20", "40"})
}

func TestListOfStrings(t *testing.T) {
	mList := make([]string, 0)

	a := convert.ListOfStrings(nil)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStrings(mList)
	EqualStringsArray(t, a, mList)

	a = convert.ListOfStrings(mList)
	EqualStringsArray(t, a, mList)

	mListFull := []string{"10", "20", "40"}
	a = convert.ListOfStrings(mListFull)
	EqualStringsArray(t, a, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a = convert.ListOfStrings(mListEmpty)
	EqualStringsArray(t, a, mListEmpty)
}
