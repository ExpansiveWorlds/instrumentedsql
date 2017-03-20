package instrumentedsql

type Opt func(*wrappedDriver)

func WithLogger(l Logger) Opt {
	return func(w *wrappedDriver) {
		w.Logger = l
	}
}

func WithTracer(t Tracer) Opt {
	return func(w *wrappedDriver) {
		w.Tracer = t
	}
}