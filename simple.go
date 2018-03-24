package multicloser

import (
	"fmt"
	"io"
)

// SimpleMultiCloser implements a very basic MultiCloser. It handles errors by
// simply appending them together in the Close method.
type SimpleMultiCloser struct {
	closers []io.Closer
}

// Add implements MultiCloser
func (m *SimpleMultiCloser) Add(c io.Closer, e error) io.Closer {
	if e != nil {
		panic(e)
	}
	m.closers = append(m.closers, c)
	return c
}

// Close implements io.Closer
func (m *SimpleMultiCloser) Close() (err error) {
	for _, c := range m.closers {
		if cerr := c.Close(); cerr != nil {
			if err == nil {
				err = cerr
			} else {
				err = fmt.Errorf("%v: %v", err, cerr)
			}
		}
	}
	return
}
