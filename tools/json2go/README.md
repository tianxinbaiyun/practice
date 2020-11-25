## 概述
json2go主要用于将json文件转换为Golang所使用的结构体，可选为屏幕输出、文件输出两种方式。
## 详细
### 概述

json2go是一个基于Golang开发的轻量json文件解析、转换命令行工具，目前支持转换输出到屏幕、输出到文件两种方式。

### 功能

- json文件读取、解析
- golang结构体输出
- 支持输出方式

- 输出到屏幕终端
- 输出到.go文件
### 安装
```
$ go get -u github.com/usthooz/json2go
$ go build
$ go install
```
### 实现思路

在第三方对接时，经常需要将响应的json文件转换为结构体，网上也有类似的工具进行直接转换，但是作为开发者，相信是更喜欢在命令行工作的，所以开发了这款json转换工具。

### 流程结构

如下图所示为项目实现流程及结构：
![alt](http://www.demodashi.com/contentImages/image/20190218/zeotmZsJV7vLjcGjuuv.png)

代码目录结构

### 目录结构

主要代码

```
//常量及变量定义
const (
    // 主命令
    exec = "json2go"
    // version 当前版本
    version = "v1.0"
)
var (
    // command 命令
    command string
    // workPath current work path
    workPath string
    // jsonFile json文件名称
    jsonFile string
    // outputFile 输出文件名称
    outFile string
    // outType 输出类型
    outType string
)
var (
    // commandsMap 命令集
    commandMap map[string]*Command
)
// Command
type Command struct {
    Name   string
    Detail string
    Func   func(name, detail string)
}
//命令初始化

// initCommands
func initCommands() {
  for i, v := range os.Args {
      switch i {
      case 1:
          command = v
      }
  }
  // 初始化命令列表
  commandMap = map[string]*Command{
      "v": &Command{
          Name:   "v",
          Detail: "查看当前版本号",
          Func:   getVersion,
      },
      "help": &Command{
          Name:   "help",
          Detail: "查看帮助信息",
          Func:   getHelp,
      },
      "gen_types": &Command{
          Name:   "gen_types",
          Detail: "根据json文件自动生成struct",
          Func:   genStruct,
      },
  }
}
```
main方法

在使用时，main作为主要调用方，完成命令衔接。
```
func main() {
    // 获取当前目录
    getWorkDir()
    // 初始化命令
    initCommands()
    if len(os.Args) < 2 {
        getHelp("help", commandMap["help"].Detail)
        return
    }
    flag.CommandLine.Parse(os.Args[2:])
    if !checkArgs() {
        return
    }
    c := commandMap[command]
    if c == nil {
        getHelp("help", commandMap["help"].Detail)
        return
    } else {
        c.Func(c.Name, c.Detail)
    }
}
```
### 使用

屏幕输入json2go或者json2go help查看帮助信息，如下图所示。
![alt](http://www.demodashi.com/contentImages/image/20190218/S4gu2ruKHbSXhQrjI69.png)

运行

默认输出到屏幕终端，如下图所示

![alt](http://www.demodashi.com/contentImages/image/20190218/AhGQ8FX7g7ksPJKuPtS.png)

### 常用命令及方法

- 新建json文件
- 使用命令将json文件转换为Golang结构体，可选择输出到文件以及屏幕
- 使用默认配置
```
json2go gen_types
```
输出到文件
```
json2go gen_types -out_type=file -json_file=json2go.json -out_file=out_types.go
```
输出到屏幕
```
json2go gen_types -out_type=print
```