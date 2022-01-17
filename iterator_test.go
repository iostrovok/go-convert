package convert_test

import (
	. "github.com/iostrovok/check"

	"github.com/iostrovok/go-convert"
)

func (s *testSuite) TestSimple(c *C) {
	iter, err := convert.Iterator([]interface{}{})
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), IsNil)

	iter, err = convert.Iterator([]string{})
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), IsNil)

	iter, err = convert.Iterator([]string{}, false)
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), IsNil)

	iter, err = convert.Iterator([]int{1})
	c.Assert(err, IsNil)
	c.Assert(iter.NextNotNil(), DeepEquals, 1)
	c.Assert(iter.NextNotNil(), IsNil)

	_, err = convert.Iterator("")
	c.Assert(err, NotNil)
}

func (s *testSuite) TestLength(c *C) {
	_, err := convert.Iterator([]interface{}{}, true)
	c.Assert(err, NotNil)

	_, err = convert.Iterator([]string{}, true)
	c.Assert(err, NotNil)
}

func (s *testSuite) TestIterator(c *C) {
	data := []interface{}{
		"a",
		nil,
		1,
		"b",
	}

	iter, err := convert.Iterator(data)
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

	iter, err := convert.Iterator(data)
	c.Assert(err, IsNil)

	str := iter.NextNotEmptyString()
	c.Assert(str, Equals, "a")

	str = iter.NextNotEmptyString()
	c.Assert(str, Equals, "b")

	str = iter.NextNotEmptyString()
	c.Assert(str, Equals, "")

}

func (s *testSuite) TestCheckMapStringType(c *C) {
	_, notFind := convert.CheckMapStringType(nil)
	c.Assert(notFind, Equals, false)

	a := map[string]interface{}{
		"string_a": 1,
	}
	check, find := convert.CheckMapStringType(a)

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

	iter, err := convert.Iterator(data)
	c.Assert(err, IsNil)

	c.Assert(iter.NextNotNilMapString(), DeepEquals, a)
	c.Assert(iter.NextNotNilMapString(), DeepEquals, b)
	c.Assert(iter.NextNotNilMapString(), DeepEquals, d)
	c.Assert(iter.NextNotNilMapString(), IsNil)
}
