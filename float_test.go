package convert

import (
	. "gopkg.in/check.v1"

	"math"
	"reflect"
)

func (s *testSuite) TestFloat64(c *C) {
	c.Assert(Float64(""), Equals, float64(0))
	c.Assert(Float64(nil), Equals, float64(0))
	c.Assert(Float64("3049"), Equals, float64(3049))
	c.Assert(Float64(int32(3049)), Equals, float64(3049))
	c.Assert(Float64(int64(3049)), Equals, float64(3049))
	c.Assert(Float64(3049), Equals, float64(3049))
	c.Assert(Float64(3049.24343), Equals, 3049.24343)
	c.Assert(Float64(math.MaxFloat32), Equals, math.MaxFloat32)
	c.Assert(Float64(math.MaxInt64), Equals, float64(math.MaxInt64))
	c.Assert(Float64(reflect.Invalid), Equals, float64(0))
}

func (s *testSuite) TestFloat32(c *C) {
	c.Assert(Float32(""), Equals, float32(0))
	c.Assert(Float32(nil), Equals, float32(0))
	c.Assert(Float32("3049"), Equals, float32(3049))
	c.Assert(Float32(int32(3049)), Equals, float32(3049))
	c.Assert(Float32(int64(3049)), Equals, float32(3049))
	c.Assert(Float32(3049), Equals, float32(3049))
	c.Assert(Float32(3049.24343), Equals, float32(3049.24343))
	c.Assert(Float32(math.MaxFloat64), Equals, float32(0))
	c.Assert(Float32(math.MaxFloat32), Equals, float32(math.MaxFloat32))
	c.Assert(Float32(math.MaxInt64), Equals, float32(math.MaxInt64))
	c.Assert(Float32(reflect.Invalid), Equals, float32(0))
}
