package main

import (
	"fmt"
	_ "go-cmd/src/cmd"
	"go-cmd/src/config"
	"go-cmd/src/core"
	"os"
)

func main() {
	config.Init()
	defer config.Save()

	fmt.Println("提示: 方向键上下可切换历史命令.")
	fmt.Println("提示: 输入 help 获取帮助.")

	app := core.NewApp()
	defer app.Close()

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println("[Start Error]: " + err.Error())
	}
}
