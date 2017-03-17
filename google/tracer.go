package google

import (
	"context"

	"github.com/ExpansiveWorlds/traced-sql"
	"cloud.google.com/go/trace"
)

type tracer struct {}

type span struct {
	parent *trace.Span
}

// NewTracer returns a tracer that will fetch spans using google tracing's SpanContext function
func NewTracer() tracedSQL.Tracer { return tracer{} }

// GetSpan fetches a span from the context and wraps it
func (tracer) GetSpan(ctx context.Context) tracedSQL.Span {
	return span{parent: trace.FromContext(ctx)}
}

func (span span) NewChild(name string) tracedSQL.Span {
	return span{parent: span.parent.NewChild(name) }
}

func (span span) SetLabel(k, v string) {
	span.parent.SetLabel(k, v)
}

func (span span) Finish() {
	span.parent.Finish()
}