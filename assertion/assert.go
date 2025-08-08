package assertion

import (
	"testing"

	"github.com/kinakoman/go-logger/logger"
)

type Assert struct {
	t      *testing.T
	logger Logger
}

type Logger interface {
	Error(message string)
}

func NewAssert(t *testing.T) *Assert {
	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	return &Assert{t: t, logger: logger}
}
