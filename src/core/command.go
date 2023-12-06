package core

import "github.com/urfave/cli"

var (
	commands cli.Commands = cli.Commands{}
)

// 注册命令
func RegistCommand(cmd cli.Command) {
	commands = append(commands, cmd)
}
