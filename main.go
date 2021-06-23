package main

import (
	"ginhelper/core"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app *cli.App

const (
	NAME    = "ginhelper"
	USAGE   = "Create a scaffold for the gin framework"
	VERSION = "0.1.0"
)

func init() {
	app = &cli.App{
		Name:    NAME,
		Usage:   USAGE,
		Version: VERSION,
		Action:  action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "指定`项目名称`",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "指定`项目路径`",
			},
			&cli.StringFlag{
				Name:  "remote",
				Usage: "指定`GIT地址`",
			},
			&cli.BoolFlag{
				Name:        "vendor",
				Usage:       "automatic command 'go mod download'",
				DefaultText: "true",
			},
			&cli.BoolFlag{
				Name:        "docker",
				Usage:       "enable docker",
				DefaultText: "true",
			},
			&cli.BoolFlag{
				Name:        "compose",
				Usage:       "use docker compose file",
				DefaultText: "true",
			},
		},
	}
}

func action(c *cli.Context) error {
	return core.Execute(
		c.String("name"),
		c.String("output"),
		c.String("remote"),
		c.Bool("vendor"),
		c.Bool("docker"),
		c.Bool("compose"),
	)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
