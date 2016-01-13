package trace

import "io"

// Tracerはコード内での出来事を記録できるオブジェクトを表すインタフェースです。
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {}
