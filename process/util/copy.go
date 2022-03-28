package util

import (
	"context"
	"io"

	"github.com/streamwest-1629/go-utilpkgs/errors"
)

var (
	bufferSize      = 32 * 1024
	ErrInvalidWrite = errors.New("invalid write result")
)

func Copy(dst io.Writer, src io.Reader, ctx context.Context) (written int64, err error) {
	return copyBuffer(dst, src, nil, ctx)
}

func copyBuffer(dst io.Writer, src io.Reader, buf []byte, ctx context.Context) (written int64, err error) {

	// check src/dst argument is nil
	if dst == nil {
		return 0, errors.New("argument is nil", errors.NewKV("arg", "dst"))
	} else if src == nil {
		return 0, errors.New("argument is nil", errors.NewKV("arg", "src"))
	}

	// allocate buffer
	if buf == nil {
		buf = make([]byte, bufferSize)
	}

	for {
		written = 0
		if ctx.Err() != nil {
			return written, nil
		}

		if nread, err := src.Read(buf); nread > 0 {
			if ctx.Err() != nil {
				return written, nil
			}

			nwrite, err := dst.Write(buf)
			if nwrite < 0 || nread < nwrite {
				if err != nil {
					return written, err
				} else {
					return written, ErrInvalidWrite
				}
			}
			written += int64(nwrite)
			if nread != nwrite {
				err = io.ErrShortWrite
			}

		} else if err != nil {
			if err == io.EOF {
				return written, nil
			} else {
				return written, err
			}
		}
	}
}
