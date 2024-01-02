package convert_test

import (
	"testing"

	"github.com/iostrovok/go-convert"
)

func TestString(t *testing.T) {
	Equal(t, convert.String("3049"), "3049")
	Equal(t, convert.String(nil), "")
	Equal(t, convert.String(""), "")
	Equal(t, convert.String(" "), " ")
	Equal(t, convert.String(int32(3049)), "3049")
	Equal(t, convert.String(int64(3049)), "3049")
	Equal(t, convert.String(3049), "3049")
	Equal(t, convert.String(3049.24343), "3049.24343")
	Equal(t, convert.String(float32(3049.2434)), "3049.2434")
	Equal(t, convert.String([]byte{65, 66, 67, 226, 130, 172}), "ABC€")

	a := make([]string, 1)
	Equal(t, convert.String(a), "")
	Equal(t, convert.String(&a), "")
}
