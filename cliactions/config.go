package cliactions

import (
	"log"

	"github.com/andreynering/gogscli/config"
	"github.com/codegangsta/cli"
)

func Config() cli.ActionFunc {
	return func(c *cli.Context) error {
		var (
			cfg, _ = config.Get()
			remote = c.String("remote")
			token  = c.String("token")
		)
		if len(remote) > 0 {
			cfg.Remote.URL = remote
		}
		if len(token) > 0 {
			cfg.Auth.Token = token
		}

		if err := cfg.Save(); err != nil {
			log.Fatal(err)
		}
		return nil
	}
}
