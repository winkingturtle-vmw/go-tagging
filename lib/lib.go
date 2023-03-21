package lib

import (
	"github.com/urfave/cli/v2"
)

type Lib struct {
	App *cli.App
}

func New(name string) Lib {
	app := &cli.App{
		Name: name,
	}
	lib := Lib{
		App: app,
	}
	return lib
}

func (l *Lib) WithUsage(usage string) {
	l.App.Usage = usage
}

func (l *Lib) WithUsageText(usageText string) {
	l.App.UsageText = usageText
}

func (l *Lib) WithArgsUsage(args string) {
	l.App.ArgsUsage = args
}
