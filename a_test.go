package convert

import (
	. "gopkg.in/check.v1"

	"testing"
)

type testSuite struct{}

var _ = Suite(&testSuite{})

func TestService(t *testing.T) { TestingT(t) }
