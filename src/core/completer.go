package core

import (
	"github.com/urfave/cli"
	"strings"
)

var (
	helpFlags    = []string{"--help", "-h"}
	versionFlags = []string{"--version", "-v"}
)

// 子命令递归补全
func subCompleter(args []string, cmd cli.Command, index int) (c []string) {
	numArgs := len(args)
	// 尝试补全flag
	if strings.HasPrefix(args[index], "-") {
		return flagsCompleter(args, cmd, index)
	}
	// 尝试补全子命令
	for _, sub := range cmd.Subcommands {
		for _, name := range sub.Names() {
			if strings.HasPrefix(name, args[index]) {
				if index == numArgs-1 {
					c = append(c, name+" ")
				} else {
					subs := subCompleter(args, sub, index+1)
					for _, subStr := range subs {
						c = append(c, name+" "+subStr)
					}
				}
			}
		}
	}
	// 若未能补全且还有后续，则向下递归
	if len(c) == 0 && index < numArgs-1 {
		subs := subCompleter(args, cmd, index+1)
		for _, sub := range subs {
			c = append(c, args[index]+" "+sub)
		}
		return
	}
	return
}

// flag 补全
func flagsCompleter(args []string, cmd cli.Command, index int) (c []string) {
	if index != len(args)-1 {
		subs := subCompleter(args, cmd, index+1)
		for _, sub := range subs {
			c = append(c, args[index]+" "+sub)
		}
		return
	}
	if !cmd.HideHelp {
		for _, name := range helpFlags {
			if strings.HasPrefix(name, args[index]) {
				c = append(c, name)
			}
		}
	}
	for _, flag := range cmd.Flags {
		for _, name := range strings.Split(flag.GetName(), ",") {
			name = strings.Trim(name, " ")
			if len(name) == 1 {
				name = "-" + name
			} else {
				name = "--" + name
			}
			if strings.HasPrefix(name, args[index]) {
				c = append(c, name)
			}
		}
	}
	return
}
