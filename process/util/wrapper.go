package util

import "io"

type SetWriterFunc func(io.WriteCloser) error

func (fn SetWriterFunc) SetWriter(w io.WriteCloser) error {
	return fn(w)
}
