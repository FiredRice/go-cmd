package config

import (
	"go-cmd/src/utils"
	"io/fs"
	"path"
)

var (
	Name        = "go-cmd"
	Version     = "1.0.0"
	Author      = "东北炒饭"
	Usage       = "go 命令行应用框架"
	UsageText   = Name + " [global options] command [command options] [arguments...]"
	Copyright   = ""
	Description = "golang 编写的命令行应用程序简易框架"
)

var (
	HistoryPath     = path.Join("./tmp", "cmd_history.txt")
	configLocalPath = path.Join("./tmp", "config.json")
)

type LocalConfig struct {
	Version string `json:"version"`
}

func init() {
	utils.MkdirAll(path.Dir(HistoryPath), fs.ModePerm)
	// utils.MkdirAll(path.Dir(configLocalPath), fs.ModePerm)
}
