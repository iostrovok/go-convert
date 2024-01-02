package convert_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func Fatal(t *testing.T, err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Fatal(err)
	}
}

func NotNil2(t *testing.T, err error) {
	if err == nil {
		Fatal(t, errors.Errorf("expected not nil errors"))
	}
}

func Nil2(t *testing.T, a any) {
	if a != nil {
		Fatal(t, errors.Errorf("not nil"))
	}
}

func Equal(t *testing.T, obtained, expected any) {
	if !isType(obtained, expected) {
		Fatal(t, errors.Errorf("wrong types"))
	}

	// Simple comparing. It is enough for our goals.
	if fmt.Sprintf("%v", obtained) != fmt.Sprintf("%v", expected) {
		Fatal(t, errors.Errorf("wrong values: %v != %v", obtained, expected))
	}
}

func EqualBool(t *testing.T, obtained, expected bool) {
	if obtained != expected {
		Fatal(t, errors.Errorf("wrong values: %t != %t", obtained, expected))
	}
}

func isType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func EqualStringsArray(t *testing.T, obtained, check []string) {
	if len(obtained) != len(check) {
		err := errors.Errorf("wrong length in EqualStringsArray")
		Fatal(t, err)
	}

	for i, v := range check {
		Equal(t, obtained[i], v)
	}
}
