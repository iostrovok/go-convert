package convert_test

import (
	. "github.com/iostrovok/check"

	"github.com/iostrovok/go-convert"
)

func (s *testSuite) TestBool(c *C) {

	c.Assert(convert.Bool(1), Equals, true)
	c.Assert(convert.Bool("1"), Equals, true)
	c.Assert(convert.Bool([]byte("1")), Equals, true)
	c.Assert(convert.Bool("-1"), Equals, true)
	c.Assert(convert.Bool([]byte("-1")), Equals, true)
	c.Assert(convert.Bool(true), Equals, true)
	c.Assert(convert.Bool("True"), Equals, true)
	c.Assert(convert.Bool([]byte("True")), Equals, true)
	c.Assert(convert.Bool([]byte("true")), Equals, true)
	c.Assert(convert.Bool("true"), Equals, true)
	c.Assert(convert.Bool("t"), Equals, true)
	c.Assert(convert.Bool(0.1), Equals, true)

	c.Assert(convert.Bool("0"), Equals, false)
	c.Assert(convert.Bool([]byte("0")), Equals, false)
	c.Assert(convert.Bool(false), Equals, false)
	c.Assert(convert.Bool(""), Equals, false)
	c.Assert(convert.Bool([]byte("")), Equals, false)
	c.Assert(convert.Bool("false"), Equals, false)
	c.Assert(convert.Bool([]byte("false")), Equals, false)
	c.Assert(convert.Bool([]byte("False")), Equals, false)
	c.Assert(convert.Bool("False"), Equals, false)
	c.Assert(convert.Bool("f"), Equals, false)
	c.Assert(convert.Bool(0), Equals, false)
	c.Assert(convert.Bool(0.0), Equals, false)
}

func (s *testSuite) TestBoolErr(c *C) {
	a, err := convert.BoolErr("1")
	c.Assert(err, IsNil)
	c.Assert(a, Equals, true)
}

func (s *testSuite) TestBool_ArrayMap(c *C) {
	c.Assert(convert.Bool([]interface{}{}), Equals, false)
	c.Assert(convert.Bool(map[string]interface{}{}), Equals, false)
	c.Assert(convert.Bool(map[int]interface{}{}), Equals, false)

	c.Assert(convert.Bool([]string{}), Equals, false)
	c.Assert(convert.Bool(map[string]string{}), Equals, false)
	c.Assert(convert.Bool(map[int]string{}), Equals, false)

	c.Assert(convert.Bool([]int{}), Equals, false)
	c.Assert(convert.Bool(map[string]int{}), Equals, false)
	c.Assert(convert.Bool(map[int]int{}), Equals, false)

	c.Assert(convert.Bool([]interface{}{1}), Equals, true)
	c.Assert(convert.Bool(map[string]interface{}{"1": 1}), Equals, true)
	c.Assert(convert.Bool(map[int]interface{}{1: "asdasdasd"}), Equals, true)

	c.Assert(convert.Bool([]string{"1"}), Equals, true)
	c.Assert(convert.Bool(map[string]string{"1": "1"}), Equals, true)
	c.Assert(convert.Bool(map[int]string{1: "1"}), Equals, true)

	c.Assert(convert.Bool([]int{1}), Equals, true)
	c.Assert(convert.Bool(map[string]int{"1": 1}), Equals, true)
	c.Assert(convert.Bool(map[int]int{1: 1}), Equals, true)
}
