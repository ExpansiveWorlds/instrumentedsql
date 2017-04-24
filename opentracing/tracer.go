package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/ExpansiveWorlds/instrumentedsql"
)

type tracer struct{}

type span struct {
	parent opentracing.Span
}

// NewTracer returns a tracer that will fetch spans using opentracing's SpanFromContext function
func NewTracer() instrumentedsql.Tracer { return tracer{} }

// GetSpan returns a span
func (tracer) GetSpan(ctx context.Context) instrumentedsql.Span {
	if ctx == nil {
		return span{parent: nil}
	}

	return span{parent: opentracing.SpanFromContext(ctx)}
}

func (s span) NewChild(name string) instrumentedsql.Span {
	if s.parent == nil {
		return span{parent: opentracing.StartSpan(name)}
	}
	return span{parent: opentracing.StartSpan(name, opentracing.ChildOf(s.parent.Context()))}
}

func (s span) SetLabel(k, v string) {
	if s.parent == nil {
		return
	}
	s.parent.SetTag(k, v)
}

func (s span) Finish() {
	if s.parent == nil {
		return
	}
	s.parent.Finish()
}
