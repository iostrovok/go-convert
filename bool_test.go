package convert

import (
	. "github.com/iostrovok/check"
)

func (s *testSuite) TestBool(c *C) {

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
	a, err := BoolErr("1")
	c.Assert(err, IsNil)
	c.Assert(a, Equals, true)
}

func (s *testSuite) TestBool_ArrayMap(c *C) {
	c.Assert(Bool([]interface{}{}), Equals, false)
	c.Assert(Bool(map[string]interface{}{}), Equals, false)
	c.Assert(Bool(map[int]interface{}{}), Equals, false)

	c.Assert(Bool([]string{}), Equals, false)
	c.Assert(Bool(map[string]string{}), Equals, false)
	c.Assert(Bool(map[int]string{}), Equals, false)

	c.Assert(Bool([]int{}), Equals, false)
	c.Assert(Bool(map[string]int{}), Equals, false)
	c.Assert(Bool(map[int]int{}), Equals, false)

	c.Assert(Bool([]interface{}{1}), Equals, true)
	c.Assert(Bool(map[string]interface{}{"1": 1}), Equals, true)
	c.Assert(Bool(map[int]interface{}{1: "asdasdasd"}), Equals, true)

	c.Assert(Bool([]string{"1"}), Equals, true)
	c.Assert(Bool(map[string]string{"1": "1"}), Equals, true)
	c.Assert(Bool(map[int]string{1: "1"}), Equals, true)

	c.Assert(Bool([]int{1}), Equals, true)
	c.Assert(Bool(map[string]int{"1": 1}), Equals, true)
	c.Assert(Bool(map[int]int{1: 1}), Equals, true)
}
