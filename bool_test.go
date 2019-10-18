package convert

import (
	. "github.com/iostrovok/check"
)

func (s *testSuite) TestBool(c *C) {

	c.Assert(Bool([]interface{}{}), Equals, true)

	c.Assert(Bool(1), Equals, true)
	c.Assert(Bool("1"), Equals, true)
	c.Assert(Bool("-1"), Equals, true)
	c.Assert(Bool(true), Equals, true)
	c.Assert(Bool("True"), Equals, true)
	c.Assert(Bool("true"), Equals, true)
	c.Assert(Bool("t"), Equals, true)
	c.Assert(Bool(0.1), Equals, true)

	c.Assert(Bool("0"), Equals, false)
	c.Assert(Bool(false), Equals, false)
	c.Assert(Bool(""), Equals, false)
	c.Assert(Bool("false"), Equals, false)
	c.Assert(Bool("False"), Equals, false)
	c.Assert(Bool("f"), Equals, false)
	c.Assert(Bool(0), Equals, false)
	c.Assert(Bool(0.0), Equals, false)
}

func (s *testSuite) TestBoolErr(c *C) {
	a, err := BoolErr([]interface{}{})
	c.Assert(err, IsNil)
	c.Assert(a, Equals, true)
}
