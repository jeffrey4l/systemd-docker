package container

import (
	"os/exec"

	"github.com/jeffrey4l/systemd-docker/common"
	log "github.com/sirupsen/logrus"
)

type DockerContainer struct {
	BaseContainer
}

func NewDockerContainer(ctx *common.Context) *DockerContainer {
	runtime := ctx.Meta.Runtime
	path, err := exec.LookPath(runtime)
	if err != nil {
		log.Fatal("Can not binary for docker.", runtime)
	}
	return &DockerContainer{BaseContainer{Context: ctx, Binary: path}}
}
