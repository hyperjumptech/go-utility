package HyperUtility

import "io"

// NewMultipleWriter creates a new MultipleWriter
func NewMultipleWriter() *MultipleWriter {
	mw := &MultipleWriter{
		writers: make([]io.Writer, 0),
	}
	return mw
}

// MultipleWriter is a writer that can write to multiple writer.
// it contains an array of 'io.Writer'
type MultipleWriter struct {
	writers []io.Writer
}

// AddWriter add a writer into this multip writter
func (mw *MultipleWriter) AddWriter(w io.Writer) {
	if mw.writers == nil {
		panic("no writers")
	}
	mw.writers = append(mw.writers, w)
}

func (mw *MultipleWriter) Write(p []byte) (n int, err error) {
	if mw.writers == nil {
		panic("no writers")
	}
	for _, w := range mw.writers {
		_, _ = w.Write(p)
	}
	return len(p), nil
}
