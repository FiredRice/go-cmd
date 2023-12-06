package cmdliner

import (
	"fmt"
	"github.com/peterh/liner"
	"go-cmd/src/config"
)

type CMDLiner struct {
	State   *liner.State
	History *LineHistory
}

// NewLiner 返回 *CMDLiner, 默认设置允许 Ctrl+C 结束
func NewLiner() *CMDLiner {
	pl := &CMDLiner{}

	line := liner.NewLiner()
	line.SetMultiLineMode(true)
	line.SetCtrlCAborts(true)

	pl.State = line

	var err error
	pl.History, err = NewLineHistory(config.HistoryPath)

	if err != nil {
		fmt.Printf("警告: 读取历史命令文件错误, %s\n", err.Error())
	}

	return pl
}

// Close 关闭服务
func (pl *CMDLiner) Close() (err error) {
	err = pl.State.Close()
	if err != nil {
		return err
	}

	if pl.History != nil && pl.History.historyFile != nil {
		return pl.History.historyFile.Close()
	}

	return nil
}
