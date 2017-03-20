package google

import (
	"context"

	"cloud.google.com/go/trace"

	"github.com/ExpansiveWorlds/instrumentedsql"
)

type tracer struct{}

type span struct {
	parent *trace.Span
}

// NewTracer returns a tracer that will fetch spans using google tracing's SpanContext function
func NewTracer() instrumentedsql.Tracer { return tracer{} }

// GetSpan fetches a span from the context and wraps it
func (tracer) GetSpan(ctx context.Context) instrumentedsql.Span {
	if ctx == nil {
		return span{parent: nil}
	}

	return span{parent: trace.FromContext(ctx)}
}

func (s span) NewChild(name string) instrumentedsql.Span {
	return span{parent: s.parent.NewChild(name)}
}

func (s span) SetLabel(k, v string) {
	s.parent.SetLabel(k, v)
}

func (s span) Finish() {
	s.parent.Finish()
}
