package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := &cli.App{
		Name: "submarine",
		UsageText: "submarine {copy|view}",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "copy",
				Aliases: []string{"c"},
				Usage: "copy one solution file to another. can read or write from .json and .solution files.",
				UsageText: "submarine copy [src] [dest]",
				Action: actionCopy,
			},
			{
				Name: "view",
				Aliases: []string{"v"},
				Usage: "view one or more .solution or .json files",
				UsageText: "submarine view [files...]",
				Action: actionView,
			},
		},
	}

	app.Run(os.Args)
}
