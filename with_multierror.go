package multicloser

import (
	"io"

	"github.com/hashicorp/go-multierror"
)

// MultiErrorMultiCloser implements a MultiCloser that uses Hashicorp's
// multierror library to collect the errors together in the Close method.
type MultiErrorMultiCloser struct {
	closers []io.Closer
}

// Add implements MultiCloser
func (m *MultiErrorMultiCloser) Add(c io.Closer, e error) io.Closer {
	if e != nil {
		panic(e)
	}
	m.closers = append(m.closers, c)
	return c
}

// Close implements io.Closer
func (m *MultiErrorMultiCloser) Close() (err error) {
	for _, c := range m.closers {
		err = multierror.Append(err, c.Close())
	}
	return
}
