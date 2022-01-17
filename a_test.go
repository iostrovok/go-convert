package convert_test

import (
	. "github.com/iostrovok/check"

	"testing"
)

type testSuite struct{}

var _ = Suite(&testSuite{})

func TestService(t *testing.T) { TestingT(t) }
