package instrumentedsql

import "context"

// Logger is the interface needed to be implemented by any logging implementation we use, see also NewFuncLogger
type Logger interface {
	Log(ctx context.Context, msg string, keyvals ...interface{})
}

// LogFunc is a logging function that can be passed to NewFuncLogger
type LogFunc func(ctx context.Context, msg string, keyvals ...interface{})

type nullLogger struct{}

func (nullLogger) Log(ctx context.Context, msg string, keyvals ...interface{}) {}

type funcLogger struct {
	logger LogFunc
}

// NewFuncLogger accepts a logging function and returns a matching Logger for it
func NewFuncLogger(logger func(ctx context.Context, msg string, keyvals ...interface{})) Logger {
	return funcLogger{logger: logger}
}

func (l funcLogger) Log(ctx context.Context, msg string, keyvals ...interface{}) {
	l.logger(ctx, msg, keyvals...)
}
