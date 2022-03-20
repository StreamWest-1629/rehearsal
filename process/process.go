package process

import (
	"io"

	"github.com/streamwest-1629/go-utilpkgs/errors"
)

type (
	ProcIn interface {
		SetWriter(w io.Writer) error
	}

	ProcOut interface {
		io.Writer
	}

	ProcIO struct {
		key    string
		input  ProcIn
		output ProcOut
	}

	ProcCommander interface {
		Begin() error
		Stop() error
	}

	pausable interface {
		Pause() error
		Resume() error
	}

	Process struct {
		commander ProcCommander
		pause     func() error
		resume    func() error
		inputs    map[string]ProcIn
		outputs   map[string]ProcOut
	}
)

func NewProcOut(key string, output ProcOut) *ProcIO {
	return &ProcIO{
		key:    key,
		input:  nil,
		output: output,
	}
}

func NewProcIn(key string, input ProcIn) *ProcIO {
	return &ProcIO{
		key:    key,
		input:  input,
		output: nil,
	}
}

func NewProcess(commander ProcCommander, ios ...*ProcIO) (*Process, error) {
	proc := &Process{
		commander: commander,
		pause:     nil,
		resume:    nil,
		inputs:    make(map[string]ProcIn),
		outputs:   make(map[string]ProcOut),
	}

	// check process is pausable
	if pausable, ok := commander.(pausable); ok {
		proc.pause, proc.resume = pausable.Pause, pausable.Resume
	}

	// register process i/o module
	for _, io := range ios {
		if io.input != nil && io.output == nil {

			// check key duplicated
			if _, exist := proc.inputs[io.key]; exist {
				return nil, errors.New("one of input module's key is duplicated",
					errors.NewKV("duplicated key", io.key))
			}

			proc.inputs[io.key] = io.input

		} else if io.output != nil {

			// check key duplicated
			if _, exist := proc.outputs[io.key]; exist {
				return nil, errors.New("one of output module's key is duplicated",
					errors.NewKV("duplicated key", io.key))
			}

			proc.outputs[io.key] = io.output

		} else {
			// invalid ProcIO module
			return nil, errors.New("invalid ProcIO instance")
		}
	}

	return proc, nil
}
