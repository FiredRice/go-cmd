package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"go-cmd/src/config"
	"go-cmd/src/core"
)

func init() {
	core.RegistCommand(cli.Command{
		Name:  "config",
		Usage: "配置信息",
		Action: func(c *cli.Context) {
			jsonBytes, err := json.Marshal(config.Get())
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			var out bytes.Buffer
			json.Indent(&out, jsonBytes, "", "  ")
			fmt.Println(out.String())
		},
	})
}
