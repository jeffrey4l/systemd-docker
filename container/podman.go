package container

import (
	"os/exec"

	"github.com/jeffrey4l/systemd-docker/common"
	log "github.com/sirupsen/logrus"
)

type PodmanContainer struct {
	BaseContainer
}

func NewPodmanContainer(ctx *common.Context) *PodmanContainer {
	runtime := ctx.Meta.Runtime
	path, err := exec.LookPath(runtime)
	if err != nil {
		log.Fatal("Can not binary for podman.", runtime)
	}
	return &PodmanContainer{BaseContainer{Context: ctx, Binary: path}}
}

func (p *PodmanContainer) Exists() bool {
	return true
}

func (p *PodmanContainer) Run() error {
	return nil
}
