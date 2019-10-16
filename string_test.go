package convert

import (
	. "gopkg.in/check.v1"
)

func (s *testSuite) TestString(c *C) {
	c.Assert(String("3049"), Equals, "3049")
	c.Assert(String(""), Equals, "")
	c.Assert(String(int32(3049)), Equals, "3049")
	c.Assert(String(int64(3049)), Equals, "3049")
	c.Assert(String(3049), Equals, "3049")
	c.Assert(String(3049.24343), Equals, "3049.24343")
	c.Assert(String(float32(3049.2434)), Equals, "3049.2434")

	a := make([]string, 1)
	c.Assert(String(&a), Equals, "")
}
