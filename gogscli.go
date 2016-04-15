package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/andreynering/gogscli/cliactions"
	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
)

func ensureEndsWithSlash(str string) string {
	if len(str) == 0 {
		return str
	}
	if !strings.HasSuffix(str, "/") {
		str = str + "/"
	}
	return str
}

func main() {
	auth := cliactions.Auth{
		URL:      ensureEndsWithSlash(os.Getenv("GOGS_URL")),
		User:     os.Getenv("GOGS_USER"),
		Password: os.Getenv("GOGS_PASSWORD"),
	}
	if len(auth.URL) == 0 {
		fmt.Printf("Missing GOGS_URL enviroment variable\n")
		return
	}
	if len(auth.User) == 0 {
		fmt.Printf("Missing GOGS_USER enviroment variable\n")
		return
	}
	if len(auth.Password) == 0 {
		fmt.Printf("Missing GOGS_PASSWORD enviroment variable\n")
		return
	}

	app := cli.NewApp()
	app.Name = "Gogs Cli"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "repository, r",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "issue",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "title",
						},
						cli.StringFlag{
							Name: "body",
						},
					},
					Action: cliactions.IssueAdd(auth),
				},
			},
		},
	}
	app.Run(os.Args)
}
