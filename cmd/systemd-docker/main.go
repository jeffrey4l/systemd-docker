package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jeffrey4l/systemd-docker/common"
	"github.com/jeffrey4l/systemd-docker/container"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func JoinCmd(cmd []string) string {
	var newCmd []string
	for _, v := range cmd {
		if strings.Contains(v, " ") {
			v = fmt.Sprintf(`"%s"`, v)
		}
		newCmd = append(newCmd, v)
	}
	return strings.Join(newCmd, " ")
}

func RunContainerAction(c *cli.Context) error {
	config := c.String("config")
	file, err := os.Open(config)
	if err != nil {
		log.Fatalf("Can not open config path: %s", err)
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	ctx := common.Context{}
	log.Debugf("Get context: %s", ctx)
	if err = yaml.UnmarshalStrict(data, &ctx); err != nil {
		log.Fatal(err)
	}

	if c.String("runtime") != "" {
		ctx.Meta.Runtime = c.String("runtime")
	}

	r := container.NewContainerRuntime(&ctx)
	log.Debugf("Start run container: %s", r.Name())

	if c.Bool("print") {
		fmt.Printf("%s\n", r.GenCmd())
		return nil
	} else {
		return r.Run()
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "systemd-docker"
	app.HelpName = "systemd-docker"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		{
			Name: "run",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Usage: "Load container configuration from `FILE`",
				},
				cli.BoolFlag{
					Name:  "print, p",
					Usage: "Print the command",
				},
				cli.StringFlag{
					Name:  "runtime, r",
					Usage: "Container runtime",
				},
			},
			Action: RunContainerAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
