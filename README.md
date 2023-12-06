# GO-CMD
## 简介
用 go 语言实现的简易命令行工具框架

## 安装
```sh
go install github.com/tc-hib/go-winres@latest
go mod tidy
```

## 编译
```sh
go-winres make
go build -ldflags "-w -s"
```

## 关于图标

首先下载插件
```sh
go install github.com/tc-hib/go-winres@latest
```

初始化，替换icon.png
```sh
go-winres init
```

制作 syso 文件
```sh
go-winres make
```