package convert

import (
	. "github.com/iostrovok/check"
)

func (s *testSuite) TestString(c *C) {
	c.Assert(String("3049"), Equals, "3049")
	c.Assert(String(nil), Equals, "")
	c.Assert(String(""), Equals, "")
	c.Assert(String(" "), Equals, " ")
	c.Assert(String(int32(3049)), Equals, "3049")
	c.Assert(String(int64(3049)), Equals, "3049")
	c.Assert(String(3049), Equals, "3049")
	c.Assert(String(3049.24343), Equals, "3049.24343")
	c.Assert(String(float32(3049.2434)), Equals, "3049.2434")
	c.Assert(String([]byte{65, 66, 67, 226, 130, 172}), Equals, "ABC€")

	a := make([]string, 1)
	c.Assert(String(&a), Equals, "")
}

func (s *testSuite) TestListOfStringsPErr(c *C) {

	mList := make([]string, 0)

	a, e := ListOfStringsPErr(mList, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mList)

	_, e = ListOfStringsPErr(mList, true)
	c.Assert(e, NotNil)

	mListFull := []string{"10", "20", "40"}
	a, e = ListOfStringsPErr(mListFull, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	_, e = ListOfStringsPErr(mListEmpty, false)
	c.Assert(e, NotNil)
}

func (s *testSuite) TestListOfStringsErr(c *C) {

	mList := make([]string, 0)

	a, e := ListOfStringsErr(mList, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mList)

	_, e = ListOfStringsErr(mList, true)
	c.Assert(e, NotNil)

	_, e = ListOfStringsErr(nil, true)
	c.Assert(e, NotNil)

	mListFull := []string{"10", "20", "40"}
	a, e = ListOfStringsErr(mListFull, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a, e = ListOfStringsErr(mListEmpty, false)
	c.Assert(e, IsNil)
	c.Assert(a, DeepEquals, mListEmpty)
}

func (s *testSuite) TestListOfStringsP(c *C) {

	mList := make([]string, 0)

	a := ListOfStringsP(nil, true)
	c.Assert(a, DeepEquals, mList)

	a = ListOfStringsP(mList, false)
	c.Assert(a, DeepEquals, mList)

	a = ListOfStringsP(mList, true)
	c.Assert(a, DeepEquals, mList)

	mListFull := []string{"10", "20", "40"}
	a = ListOfStringsP(mListFull, false)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a = ListOfStringsP(mListEmpty, false)
	c.Assert(a, DeepEquals, mListFull)
}

func (s *testSuite) TestListOfStrings(c *C) {

	mList := make([]string, 0)

	a := ListOfStrings(nil, true)
	c.Assert(a, DeepEquals, mList)

	a = ListOfStrings(mList, false)
	c.Assert(a, DeepEquals, mList)

	a = ListOfStrings(mList, true)
	c.Assert(a, DeepEquals, mList)

	mListFull := []string{"10", "20", "40"}
	a = ListOfStrings(mListFull, false)
	c.Assert(a, DeepEquals, mListFull)

	mListEmpty := []string{"10", "20", "40", ""}
	a = ListOfStrings(mListEmpty, false)
	c.Assert(a, DeepEquals, mListEmpty)
}
