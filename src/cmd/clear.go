package cmd

import (
	"github.com/urfave/cli"
	"go-cmd/src/core"
	"go-cmd/src/utils"
)

func init() {
	core.RegistCommand(cli.Command{
		Name:  "clear",
		Usage: "清空面板",
		Action: func(c *cli.Context) {
			utils.ClearScreen()
		},
	})
}
