package lib

import "github.com/urfave/cli"

type Lib struct {
	App *cli.App
}

func New(name string) Lib {
	app := cli.NewApp()
	app.Name = name
	lib := Lib{
		App: app,
	}
	return lib
}

func (l *Lib) WithUsage(usage string) {
	l.App.Usage = usage
}
