package multicloser

import (
	"io"

	"github.com/hashicorp/go-multierror"
)

// WithHandler implements a basic MultiCloser with the ability to customise how
// errors are handled in the `Add` method.
type WithHandler struct {
	ErrorHandler func(error)
	closers      []io.Closer
}

// Add implements MultiCloser
// If `e` is not nil, the handler will be called before returning nil.
func (m *WithHandler) Add(c io.Closer, e error) io.Closer {
	if e != nil {
		m.ErrorHandler(e)
		return nil
	}
	m.closers = append(m.closers, c)
	return c
}

// Close implements io.Closer
func (m *WithHandler) Close() (err error) {
	for _, c := range m.closers {
		err = multierror.Append(err, c.Close())
	}
	return
}
