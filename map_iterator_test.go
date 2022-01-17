package convert_test

import (
	. "github.com/iostrovok/check"

	"github.com/iostrovok/go-convert"
)

func (s *testSuite) TestMapIteratorEmpty(c *C) {

	iter, err := convert.MapIterator(map[string]interface{}{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b := iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)

	iter, err = convert.MapIterator(map[string]int32{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b = iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)

	iter, err = convert.MapIterator(map[int][]int32{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b = iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)

	iter, err = convert.MapIterator(map[string]string{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b = iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)
}

func (s *testSuite) TestMapIteratorLength(c *C) {

	_, err := convert.MapIterator(map[string]interface{}{}, true)
	c.Assert(err, NotNil)

	_, err = convert.MapIterator(map[int][]int32{}, true)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestMapIteratorRun_1(c *C) {

	data := map[string]interface{}{
		"a": "a",
		"b": nil,
		"c": 1,
		"d": "b",
	}

	iter, err := convert.MapIterator(data)
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, true)

	for {
		if !iter.HasNext() {
			break
		}

		k, v := iter.Next()

		val, find := data[convert.String(k)]
		c.Assert(find, Equals, true)
		c.Assert(val, Equals, v)
	}
}
