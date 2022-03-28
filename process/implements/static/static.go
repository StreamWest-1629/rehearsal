package static

import (
	"context"
	"errors"
	"io"

	"github.com/streamwest-1629/rehearsal/process"
	"github.com/streamwest-1629/rehearsal/process/util"
)

type ProcessProperty struct {
	Data io.Reader
	Ctx  context.Context
}

type proc struct {
	src        io.Reader
	dst        io.WriteCloser
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewProcess(props ProcessProperty) (*process.Process, error) {

	ctx, cancelFunc := context.WithCancel(props.Ctx)
	p := proc{
		src:        props.Data,
		dst:        nil,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}

	return process.NewProcess(
		&p,
		process.NewProcOut("out", util.SetWriterFunc(func(wc io.WriteCloser) error {

			p.dst = wc
			return nil
		})),
	)
}

func (p *proc) Begin() error {
	if p.dst == nil {
		return errors.New("destination is nil")
	} else {
		go util.Copy(p.dst, p.src, p.ctx)
		return nil
	}
}

func (p *proc) Stop() error {
	p.cancelFunc()
	return nil
}
