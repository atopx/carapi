package main

import (
	"carapi/core"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app *cli.App

const (
	NAME    = "carapi"
	USAGE   = "Create a scaffold for the gin framework"
	VERSION = "0.1.1"
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
				Name:  "frame",
				Usage: "选择`框架`, 目前支持[gin fiber], 默认gin",
				Value: "gin",
			},
			&cli.StringFlag{
				Name:  "db",
				Usage: "选择`数据库`, 目前支持[pgsql mysql], 默认pgsql",
				Value: "pgsql",
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
				Name:  "docker",
				Usage: "enable docker",
			},
			&cli.BoolFlag{
				Name:  "compose",
				Usage: "use docker compose",
			},
		},
	}
}

func action(c *cli.Context) error {
	return core.Execute(
		c.String("name"),
		c.String("output"),
		c.String("remote"),
		c.String("frame"),
		c.Bool("docker"),
		c.Bool("compose"),
	)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
