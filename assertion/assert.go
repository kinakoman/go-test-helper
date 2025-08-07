package assertion

import "testing"

type Assert struct {
	t *testing.T
}

func NewAssert(t *testing.T) *Assert {
	return &Assert{t: t}
}
