package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"go-cmd/src/config"
	"go-cmd/src/core"
)

func init() {
	core.RegistCommand(cli.Command{
		Name:  "version",
		Usage: "版本号",
		Action: func(c *cli.Context) {
			fmt.Println(config.Version)
		},
	})
}
