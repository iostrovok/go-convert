package convert

import (
	. "gopkg.in/check.v1"

	"math"
	"reflect"
)

func (s *testSuite) TestUint64(c *C) {
	c.Assert(Uint64(""), Equals, uint64(0))
	c.Assert(Uint64(nil), Equals, uint64(0))
	c.Assert(Uint64("3049"), Equals, uint64(3049))
	c.Assert(Uint64(int32(3049)), Equals, uint64(3049))
	c.Assert(Uint64(int64(3049)), Equals, uint64(3049))
	c.Assert(Uint64(3049), Equals, uint64(3049))
	c.Assert(Uint64(3049.24343), Equals, uint64(3049))
	c.Assert(Uint64(math.MaxUint32), Equals, uint64(math.MaxUint32))
	c.Assert(Uint64(math.MaxInt64), Equals, uint64(math.MaxInt64))
	c.Assert(Uint64(reflect.Invalid), Equals, uint64(0))
}

func (s *testSuite) TestUint32(c *C) {
	c.Assert(Uint32(""), Equals, uint32(0))
	c.Assert(Uint32(nil), Equals, uint32(0))
	c.Assert(Uint32("3049"), Equals, uint32(3049))
	c.Assert(Uint32(int32(3049)), Equals, uint32(3049))
	c.Assert(Uint32(int64(3049)), Equals, uint32(3049))
	c.Assert(Uint32(3049), Equals, uint32(3049))
	c.Assert(Uint32(3049.24343), Equals, uint32(3049))
	c.Assert(Uint32(3032432434323449.24343), Equals, uint32(0))
	c.Assert(Uint32(math.MaxUint32), Equals, uint32(math.MaxUint32))
	c.Assert(Uint32(math.MaxInt64), Equals, uint32(0))
	c.Assert(Uint32(reflect.Invalid), Equals, uint32(0))
}
