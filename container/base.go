package container

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/jeffrey4l/systemd-docker/common"
	log "github.com/sirupsen/logrus"
)

type Container interface {
	Exists() bool
	Run() error
	GenCmd() *Cmd
	Stop(int) error
	Delete(bool) error
	Name() string
	GetContext() *common.Context
}

type BaseContainer struct {
	Context *common.Context
	Binary  string
}

func (c *BaseContainer) run(cmd []string) ([]byte, error) {
	log.Debugf("Running command: %s", cmd)
	output, err := exec.Command(c.Binary, cmd...).Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Fatalf("Failed to run command, get return code: %d, Get stdout: %s", exitErr.ExitCode(), exitErr.Stderr)
		} else {
			log.Fatalf("Get failed: %v, %s", err, err)
		}
	}
	return output, err

}

func (c *BaseContainer) Name() string {
	return c.Context.Name
}

func (c *BaseContainer) GetContext() *common.Context {
	return c.Context
}

func (c *BaseContainer) Exists() bool {
	filter := fmt.Sprintf("name=%s", c.Context.Name)
	format := "{{.ID}} {{.Names}}"

	cmd := []string{"ps", "--filter", filter, "--format", format}

	result, err := c.run(cmd)
	if err != nil {
		return false
	}
	return len(string(result)) != 0
}
func (c *BaseContainer) GenCmd() *Cmd {
	cmd := GenCmd(c.Binary, c.Context)
	return cmd
}

func (c *BaseContainer) Run() error {
	cmd := GenCmd(c.Binary, c.Context)
	if c.Exists() {
		//TODO(jeffrey4l): check whether the container need be re-created.
		//If not, just start the stopped container.
		if err := c.Delete(true); err != nil {
			log.Warnf("Failed to delete container: %s, err: %s", c.Context.Name, err)
		}
	}
	log.Infof("Running exec command: %v", cmd)
	return syscall.Exec(c.Binary, cmd.Cmd, os.Environ())
}

func (c *BaseContainer) Delete(force bool) error {
	cmd := []string{"container", "rm", "--force", c.Context.Name}
	log.Debugf("Start delete container: %s", c.Context.Name)
	_, err := c.run(cmd)
	return err
}

func (c *BaseContainer) Stop(timeout int) error {
	cmd := []string{"container", "stop", "--time", strconv.Itoa(timeout), c.Context.Name}
	log.Debugf("Start stop container <%s> without timeout: %d", c.Context.Name, timeout)
	_, err := c.run(cmd)
	return err
}

func NewContainerRuntime(ctx *common.Context) Container {
	runtime := ctx.Meta.Runtime
	if runtime == "podman" {
		return NewPodmanContainer(ctx)
	}
	return NewDockerContainer(ctx)
}
