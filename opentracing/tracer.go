package opentracing

import (
	"context"

	"github.com/ExpansiveWorlds/traced-sql"
	"github.com/opentracing/opentracing-go"
)

type tracer struct {}

type span struct {
	parent opentracing.Span
}

// NewTracer returns a tracer that will fetch spans using opentracing's SpanFromContext function
func NewTracer() tracedSQL.Tracer { return tracer{} }

// GetSpan returns a span
func (tracer) GetSpan(ctx context.Context) tracedSQL.Span {
	if ctx == nil {
		return span{parent: nil}
	}

	return span{parent: opentracing.SpanFromContext(ctx)}
}

func (s span) NewChild(name string) tracedSQL.Span {
	return span{parent: opentracing.StartSpan(name, opentracing.ChildOf(s.parent.Context()))}
}

func (s span) SetLabel(k, v string) {
	s.parent.SetTag(k, v)
}

func (s span) Finish() {
	s.parent.Finish()
}
