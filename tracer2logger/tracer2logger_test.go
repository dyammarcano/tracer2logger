package tracer2logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLoggerTracer(t *testing.T) {
	cfg := &Config{
		logDir:       "/tmp",
		serviceName:  "testService",
		maxFileSize:  100,
		maxAge:       1,
		maxBackups:   1,
		localTime:    true,
		compress:     true,
		stdout:       true,
		rotateByDate: true,
		tracing:      true,
		structured:   true,
		filename:     "test.log",
		instance:     "testInstance",
	}

	ctx := context.Background()
	loggerTracer, err := NewLoggerTracer(ctx, cfg)

	assert.NoError(t, err)
	assert.NotNil(t, loggerTracer)
}
