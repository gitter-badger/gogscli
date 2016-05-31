package main

import (
	"os"

	"github.com/andreynering/gogscli/cliactions"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Gogs Cli"
	app.Commands = []cli.Command{
		{
			Name: "config",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "remote",
				},
				cli.StringFlag{
					Name: "token",
				},
			},
			Action: cliactions.Config(),
		},
		{
			Name: "issue",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "assignee",
						},
					},
					Action: cliactions.IssueAdd(),
				},
				{
					Name:   "list",
					Action: cliactions.IssueList(),
				},
			},
		},
	}
	app.Run(os.Args)
}
