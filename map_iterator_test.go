package convert

import (
	. "github.com/iostrovok/check"
)

func (s *testSuite) TestMapIteratorEmpty(c *C) {

	iter, err := MapIterator(map[string]interface{}{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b := iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)

	iter, err = MapIterator(map[string]int32{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b = iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)

	iter, err = MapIterator(map[int][]int32{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b = iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)

	iter, err = MapIterator(map[string]string{})
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, false)
	a, b = iter.Next()
	c.Assert(a, IsNil)
	c.Assert(b, IsNil)
}

func (s *testSuite) TestMapIteratorLength(c *C) {

	_, err := MapIterator(map[string]interface{}{}, true)
	c.Assert(err, NotNil)

	_, err = MapIterator(map[int][]int32{}, true)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestMapIteratorRun_1(c *C) {

	data := map[string]interface{}{
		"a": "a",
		"b": nil,
		"c": 1,
		"d": "b",
	}

	iter, err := MapIterator(data)
	c.Assert(err, IsNil)
	c.Assert(iter.HasNext(), Equals, true)

	for {
		if !iter.HasNext() {
			break
		}

		k, v := iter.Next()

		val, find := data[String(k)]
		c.Assert(find, Equals, true)
		c.Assert(val, Equals, v)
	}
}
