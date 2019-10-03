package container

import "github.com/jeffrey4l/systemd-docker/common"

func GenCmd(binary string, c *common.Context) []string {
	cmd := []string{binary}
	cmd = append(cmd, "run")
	cmd = append(cmd, "--name", c.Name)
	if c.TTY {
		cmd = append(cmd, "--tty")
	}
	if c.Detach {
		cmd = append(cmd, "--detach")
	}
	if c.Init {
		cmd = append(cmd, "--init")
	}
	if c.Rm {
		cmd = append(cmd, "--rm")
	}
	for _, v := range c.Volumes {
		cmd = append(cmd, "--volume", v)
	}
	// add image name
	cmd = append(cmd, c.Image)
	// at last, append the command
	for _, v := range c.Command {
		cmd = append(cmd, v)
	}
	return cmd
}
