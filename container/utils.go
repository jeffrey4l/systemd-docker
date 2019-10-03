package container

import (
	"fmt"
	"strings"

	"github.com/jeffrey4l/systemd-docker/common"
)

type Cmd struct {
	Binary string
	Cmd    []string
}

func NewCmd(binary string) *Cmd {
	return &Cmd{Binary: binary, Cmd: []string{binary}}
}

func (c *Cmd) AddSimple(args ...string) *Cmd {
	c.Cmd = append(c.Cmd, args...)
	return c
}

func (c *Cmd) AddPair(name string, value ...string) *Cmd {
	for _, v := range value {
		c.Cmd = append(c.Cmd, name, v)
	}
	return c
}

func (c Cmd) String() string {
	var wrapped []string
	for _, v := range c.Cmd {
		if strings.Contains(v, " ") {
			v = fmt.Sprintf(`"%s"`, v)
		}
		wrapped = append(wrapped, v)
	}
	return strings.Join(wrapped, " ")
}

func JoinArry(arg string, values []string) []string {
	cmd := []string{}
	for _, v := range values {
		cmd = append(cmd, arg, v)
	}
	return cmd
}

func GenCmd(binary string, c *common.Context) *Cmd {
	cmd := NewCmd(binary)
	cmd.AddSimple("run", "--name", c.Name)
	if c.TTY {
		cmd.AddSimple("--tty")
	}
	if c.Detach {
		cmd.AddSimple("--detach")
	}
	if c.Init {
		cmd.AddSimple("--init")
	}
	if c.Rm {
		cmd.AddSimple("--rm")
	}
	cmd.AddPair("--volume", c.Volume...)
	cmd.AddPair("--env", c.Env...)

	cmd.AddSimple(c.Image)
	cmd.AddSimple(c.Command...)
	return cmd
}
