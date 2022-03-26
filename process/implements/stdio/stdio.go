package stdio

import (
	"io"
	"os/exec"

	"github.com/streamwest-1629/rehearsal/process"
)

type ProcessProperty struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	Dir     *string  `json:"dir"`
}

type proc struct {
	cmd      exec.Cmd
	out, err io.Writer
}

func NewProcess(props ProcessProperty) (*process.Process, error) {

	p := proc{
		cmd: *exec.Command(props.Command, props.Args...),
		out: nil,
		err: nil,
	}

	if props.Dir != nil {
		p.cmd.Dir = *props.Dir
	}

	if stdin, err := p.cmd.StdinPipe(); err != nil {
		return nil, err
	} else {
		return process.NewProcess(
			&p,
			process.NewProcIn("stdin", stdin),
		)
	}
}

func (p *proc) Begin() error {
	return p.cmd.Start()
}

func (p *proc) Stop() error {
	return p.cmd.Process.Kill()
}
