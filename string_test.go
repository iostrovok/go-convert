package convert_test

import (
	. "github.com/iostrovok/check"

	"github.com/iostrovok/go-convert"
)

func (s *testSuite) TestString(c *C) {
	c.Assert(convert.String("3049"), Equals, "3049")
	c.Assert(convert.String(nil), Equals, "")
	c.Assert(convert.String(""), Equals, "")
	c.Assert(convert.String(" "), Equals, " ")
	c.Assert(convert.String(int32(3049)), Equals, "3049")
	c.Assert(convert.String(int64(3049)), Equals, "3049")
	c.Assert(convert.String(3049), Equals, "3049")
	c.Assert(convert.String(3049.24343), Equals, "3049.24343")
	c.Assert(convert.String(float32(3049.2434)), Equals, "3049.2434")
	c.Assert(convert.String([]byte{65, 66, 67, 226, 130, 172}), Equals, "ABC€")

	a := make([]string, 1)
	c.Assert(convert.String(&a), Equals, "")
}

func (s *testSuite) TestListOfStringsPErr(c *C) {

	mList := make([]string, 0)

	a, e := convert.ListOfStringsPErr(mList, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mList)

	_, e = convert.ListOfStringsPErr(mList, true)
	c.Assert(e, NotNil)

	mListFull := []string{"10", "20", "40"}
	a, e = convert.ListOfStringsPErr(mListFull, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	_, e = convert.ListOfStringsPErr(mListEmpty, false)
	c.Assert(e, NotNil)
}

func (s *testSuite) TestListOfStringsErr(c *C) {

	mList := make([]string, 0)

	a, e := convert.ListOfStringsErr(mList, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mList)

	_, e = convert.ListOfStringsErr(mList, true)
	c.Assert(e, NotNil)

	_, e = convert.ListOfStringsErr(nil, true)
	c.Assert(e, NotNil)

	mListFull := []string{"10", "20", "40"}
	a, e = convert.ListOfStringsErr(mListFull, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a, e = convert.ListOfStringsErr(mListEmpty, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mListEmpty)
}

func (s *testSuite) TestListOfStringsP(c *C) {

	mList := make([]string, 0)

	a := convert.ListOfStringsP(nil, true)
	c.Assert(a, DeepEquals, mList)

	a = convert.ListOfStringsP(mList, false)
	c.Assert(a, DeepEquals, mList)

	a = convert.ListOfStringsP(mList, true)
	c.Assert(a, DeepEquals, mList)

	mListFull := []string{"10", "20", "40"}
	a = convert.ListOfStringsP(mListFull, false)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a = convert.ListOfStringsP(mListEmpty, false)
	c.Assert(a, DeepEquals, mListFull)
}

func (s *testSuite) TestListOfStrings(c *C) {

	mList := make([]string, 0)

	a := convert.ListOfStrings(nil, true)
	c.Assert(a, DeepEquals, mList)

	a = convert.ListOfStrings(mList, false)
	c.Assert(a, DeepEquals, mList)

	a = convert.ListOfStrings(mList, true)
	c.Assert(a, DeepEquals, mList)

	mListFull := []string{"10", "20", "40"}
	a = convert.ListOfStrings(mListFull, false)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a = convert.ListOfStrings(mListEmpty, false)
	c.Assert(a, DeepEquals, mListEmpty)
}
