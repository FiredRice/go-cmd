package core

import (
	"fmt"
	"github.com/urfave/cli"
	"go-cmd/src/args"
	"go-cmd/src/cmdliner"
	"go-cmd/src/config"
	"os"
	"strings"
)

type App struct {
	app  *cli.App
	line *cmdliner.CMDLiner
}

func NewApp() *App {
	var instance App
	line := cmdliner.NewLiner()
	line.ReadHistory()

	app := cli.NewApp()
	app.Name = config.Name
	app.Version = config.Version
	app.Author = config.Author
	app.Usage = config.Usage
	app.UsageText = config.UsageText
	app.Copyright = config.Copyright
	app.Description = config.Description
	app.Commands = commands
	app.HideVersion = true
	app.Action = func(c *cli.Context) {
		if c.NArg() != 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), app.Name)
			return
		}

		for {
			prompt := app.Name + " > "

			commandLine, err := line.State.Prompt(prompt)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if commandLine == "" {
				continue
			}

			line.State.AppendHistory(commandLine)

			cmdArgs := args.Parse(commandLine)
			if len(cmdArgs) == 0 {
				continue
			}

			s := []string{os.Args[0]}
			s = append(s, cmdArgs...)

			c.App.Run(s)
		}
	}

	line.State.SetCompleter(func(line string) (cmds []string) {
		var (
			lineArgs = args.Parse(line)
			numArgs  = len(lineArgs)
		)
		if numArgs == 0 {
			return
		}
		// 补全 flags
		if strings.HasPrefix(lineArgs[0], "-") {
			for _, name := range helpFlags {
				if strings.HasPrefix(name, lineArgs[0]) {
					cmds = append(cmds, name)
				}
			}
			if !app.HideVersion {
				for _, name := range versionFlags {
					if strings.HasPrefix(name, lineArgs[0]) {
						cmds = append(cmds, name)
					}
				}
			}
			for _, flag := range app.Flags {
				for _, name := range strings.Split(flag.GetName(), ",") {
					name = strings.Trim(name, " ")
					if len(name) == 1 {
						name = "-" + name
					} else {
						name = "--" + name
					}
					if strings.HasPrefix(name, lineArgs[0]) {
						cmds = append(cmds, name)
					}
				}
			}
			return
		}
		// 补全子命令
		for _, cmd := range app.Commands {
			for _, name := range cmd.Names() {
				if strings.HasPrefix(name, lineArgs[0]) {
					if numArgs == 1 {
						cmds = append(cmds, name+" ")
					} else {
						subs := subCompleter(lineArgs, cmd, 1)
						for _, subStr := range subs {
							cmds = append(cmds, name+" "+subStr)
						}
					}
				}
			}
		}
		for _, c := range cmds {
			c += " "
		}
		return
	})

	instance.app = app
	instance.line = line

	return &instance
}

func (t *App) Commands() []cli.Command {
	return t.app.Commands
}

// 启动程序
func (t *App) Run(arguments []string) error {
	return t.app.Run(arguments)
}

func (t *App) Close() {
	t.line.DoWriteHistory()
	t.line.Close()
}
