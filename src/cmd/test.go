package cmd

import (
	"github.com/urfave/cli"
	"go-cmd/src/config"
	"go-cmd/src/core"
)

func init() {
	core.RegistCommand(cli.Command{
		Name:     "test",
		Usage:    "测试",
		Category: "测试",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "y",
				Usage: "确认更新",
			},
			cli.BoolFlag{
				Name:  "version, v",
				Usage: "版本",
			},
			cli.DurationFlag{
				Name:  "num, n, l",
				Usage: "时长",
			},
		},
		Action: func(c *cli.Context) {
			if c.IsSet("version") {
				println("1.0.1")
				return
			}
			if !c.IsSet("n") {
				println("未设置n")
				return
			}
			n := c.Duration("n")
			y := c.Bool("y")
			println(n, y)
		},
		Subcommands: cli.Commands{
			{
				Name:      "sub",
				Usage:     "子命令",
				UsageText: config.Name + " test sub",
				Action: func(c *cli.Context) {
					println("sub command")
				},
				Subcommands: cli.Commands{
					{
						Name: "child",
						Usage: "子命令的子命令",
						Action: func(c *cli.Context) {
							println("child")
						},
					},
				},
			},
		},
	})
}
