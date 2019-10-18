package convert

import (
	. "github.com/iostrovok/check"

	"math"
)

func (s *testSuite) TestInt64(c *C) {
	c.Assert(Int64(""), Equals, int64(0))
	c.Assert(Int64("3049"), Equals, int64(3049))
	c.Assert(Int64(int32(3049)), Equals, int64(3049))
	c.Assert(Int64(int64(3049)), Equals, int64(3049))
	c.Assert(Int64(3049), Equals, int64(3049))
	c.Assert(Int64(3049.24343), Equals, int64(3049))
}

func (s *testSuite) TestInt32(c *C) {
	c.Assert(Int32(""), Equals, int32(0))
	c.Assert(Int32("3049"), Equals, int32(3049))
	c.Assert(Int32(int32(3049)), Equals, int32(3049))
	c.Assert(Int32(int64(3049)), Equals, int32(3049))
	c.Assert(Int32(3049), Equals, int32(3049))
	c.Assert(Int32(3049.24343), Equals, int32(3049))
}

func (s *testSuite) TestListOfInt64Err(c *C) {
	res, err := ListOfInt64Err([]interface{}{
		"1", 12, 45, 123412323, -1, 0,
	}, false)

	c.Assert(err, IsNil)
	c.Assert(res, DeepEquals, []int64{1, 12, 45, 123412323, -1, 0})
}

func (s *testSuite) TestListOfInt64Err2(c *C) {
	_, err := ListOfInt64Err([]interface{}{
		"1", 12, 45, nil, -1, 0,
	}, false)
	c.Assert(err, NotNil)

	_, err = ListOfInt64Err([]interface{}{
		"", 12, 45, -1, 0,
	}, false)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestListOfInt64ErrEmpty(c *C) {
	_, err := ListOfInt64Err([]interface{}{}, false)
	c.Assert(err, IsNil)

	_, err = ListOfInt64Err([]interface{}{}, true)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestListOfInt32Err(c *C) {
	res, err := ListOfInt32Err([]interface{}{
		"1", 12, 45, 123412323, -1, 0,
	}, false)

	c.Assert(err, IsNil)
	c.Assert(res, DeepEquals, []int32{1, 12, 45, 123412323, -1, 0})
}

func (s *testSuite) TestListOfInt32Err2(c *C) {
	_, err := ListOfInt32Err([]interface{}{
		"1", 12, 45, nil, -1, 0,
	}, false)
	c.Assert(err, NotNil)

	_, err = ListOfInt32Err([]interface{}{
		"", 12, 45, -1, 0,
	}, false)
	c.Assert(err, NotNil)

	_, err = ListOfInt32Err([]interface{}{
		"1", 12, 45, 123412323, -1, 0, math.MaxInt64,
	}, false)
	c.Assert(err, NotNil)

	_, err = ListOfInt32Err([]interface{}{
		"1", 12, 45, 123412323, -1, 0, math.MinInt64,
	}, false)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestListOfInt32ErrEmpty(c *C) {
	_, err := ListOfInt32Err([]interface{}{}, false)
	c.Assert(err, IsNil)

	_, err = ListOfInt32Err([]interface{}{}, true)
	c.Assert(err, NotNil)
}
