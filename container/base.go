package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/jeffrey4l/systemd-docker/common"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	//Exists() bool
	Run() error
	GenCmd() []string
	//Stop() error
	//Delete() error
	//Kill() error
}

type BaseContainer struct {
	Context *common.Context
	Binary  string
}

func (c *BaseContainer) run(cmd []string) ([]byte, error) {
	return exec.Command(c.Binary, cmd...).Output()

}

func (c *BaseContainer) Exists() bool {
	filter := fmt.Sprintf("name=%s", c.Context)
	format := "{{.ID}} {{.Names}}"

	cmd := []string{c.Binary, "inspect", "--filter", filter, "--format", format}

	result, err := c.run(cmd)
	if err != nil {
		return false
	}
	return len(string(result)) != 0
}
func (c *BaseContainer) GenCmd() []string {
	cmd := GenCmd(c.Binary, c.Context)
	return cmd
}

func (c *BaseContainer) Run() error {
	cmd := GenCmd(c.Binary, c.Context)
	log.Debugf("Running cmd: %s", cmd)
	return syscall.Exec(c.Binary, cmd, os.Environ())
}

func NewContainerRuntime(ctx *common.Context) Container {
	// runtime := ctx.Meta.Runtime
	return NewDockerContainer(ctx)
	/*
		if runtime == "docker" {
			return
		}
		return NewPodmanContainer(ctx)
	*/
}
