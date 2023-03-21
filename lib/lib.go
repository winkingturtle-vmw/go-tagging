package lib

import "github.com/urfave/cli"

type Lib struct {
	App *cli.App
}

func New(name string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	return app
}
