package convert

import (
	. "github.com/iostrovok/check"
)

func (s *testSuite) TestSimple(c *C) {

	iter, err := Iterator([]interface{}{})
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), IsNil)

	iter, err = Iterator([]string{})
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), IsNil)

	iter, err = Iterator([]string{}, false)
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), IsNil)

	iter, err = Iterator([]int{1})
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), DeepEquals, 1)
	c.Assert(iter.NextNotNil(), IsNil)

	_, err = Iterator("")
	c.Assert(err, NotNil)
}

func (s *testSuite) TestLength(c *C) {

	_, err := Iterator([]interface{}{}, true)
	c.Assert(err, NotNil)

	_, err = Iterator([]string{}, true)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestIterator(c *C) {

	data := []interface{}{
		"a",
		nil,
		1,
		"b",
	}

	iter, err := Iterator(data)
	c.Assert(err, IsNil)

	c.Assert(iter.NextNotNil(), Equals, "a")
	c.Assert(iter.NextNotNil(), Equals, 1)
	c.Assert(iter.NextNotNil(), Equals, "b")
	c.Assert(iter.NextNotNil(), IsNil)
}

func (s *testSuite) TestNextNotEmptyString(c *C) {

	data := []interface{}{
		"a",
		nil,
		"",
		"b",
		"",
	}

	iter, err := Iterator(data)
	c.Assert(err, IsNil)

	str := iter.NextNotEmptyString()
	c.Assert(str, Equals, "a")

	str = iter.NextNotEmptyString()
	c.Assert(str, Equals, "b")

	str = iter.NextNotEmptyString()
	c.Assert(str, Equals, "")

}

func (s *testSuite) TestCheckMapStringType(c *C) {

	_, notFind := CheckMapStringType(nil)
	c.Assert(notFind, Equals, false)

	a := map[string]interface{}{
		"string_a": 1,
	}
	check, find := CheckMapStringType(a)

	c.Assert(find, Equals, true)
	c.Assert(check, DeepEquals, a)
}

func (s *testSuite) TestNextNotNilMapString(c *C) {

	a := map[string]interface{}{
		"string_a": 1,
	}

	b := map[string]interface{}{
		"string_b1": 0,
		"string_b2": "super",
	}

	d := map[string]interface{}{
		"string_c1": "super",
		"string_c2": "puper",
	}

	data := []interface{}{
		a,
		b,
		nil,
		d,
	}

	iter, err := Iterator(data)
	c.Assert(err, IsNil)

	c.Assert(iter.NextNotNilMapString(), DeepEquals, a)
	c.Assert(iter.NextNotNilMapString(), DeepEquals, b)
	c.Assert(iter.NextNotNilMapString(), DeepEquals, d)
	c.Assert(iter.NextNotNilMapString(), IsNil)
}
