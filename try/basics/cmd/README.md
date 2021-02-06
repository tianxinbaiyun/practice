# go 命令行操作

Go 语言自带有一套完整的命令操作工具，你可以通过在命令行中执行 go 来查看它们：

```
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        module-private  module configuration for non-public modules
        packages        package lists and patterns
        testflag        testing flags
        testfunc        testing functions

Use "go help <topic>" for more information about that topic.

```

## go bug

用于GO语言调试




## go build 

这个命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。

- 如果是普通包，当你执行 go build 之后，它不会产生任何文件。如果你需要在 $GOPATH/pkg 下生成相应的文件，那就得执行 go install。

- 如果是 main 包，当你执行 go build 之后，它就会在当前目录下生成一个可执行文件。
  如果你需要在 $GOPATH/bin 下生成相应的文件，需要执行 go install，或者使用 go build -o 路径/a.exe。

- 如果某个项目文件夹下有多个文件，而你只想编译某个文件，就可在 go build 之后加上文件名，
  例如 go build a.go；go build 命令默认会编译当前目录下的所有 go 文件。
 你也可以指定编译输出的文件名。
  例如，我们可以指定 go build -o astaxie.exe，默认情况是你的 package 名 (非 main 包)，或者是第一个源文件的文件名 (main 包)。
  （注：实际上，package 名在 Go 语言规范 中指代码中 “package” 后使用的名称，此名称可以与文件夹名不同。默认生成的可执行文件名是文件夹名。）

- go build 会忽略目录下以 _ 或 . 开头的 go 文件。
如果你的源代码针对不同的操作系统需要不同的处理，那么你可以根据不同的操作系统后缀来命名文件。
  例如有一个读取数组的程序，它对于不同的操作系统可能有如下几个源文件：
```
array_linux.go
array_darwin.go
array_windows.go
array_freebsd.go
```

go build 的时候会选择性地编译以系统名结尾的文件（Linux、Darwin、Windows、Freebsd）。
例如 Linux 系统下面编译只会选择 array_linux.go 文件，其它系统命名后缀文件全部忽略。

参数的介绍

```text
-o 　　　　　　　　　　　　指定输出的文件名，可以带上路径，例如 go build -o a/b/c
 
-i　　　　 　　　　　　　　安装相应的包，编译 + go install
 
-a 　　　　　　　　　　　　更新全部已经是最新的包的，但是对标准包不适用
 
-n 　　　　　　　　　　    把需要执行的编译命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
 
-p n 　　　　　　　　　　　指定可以并行可运行的编译数目，默认是 CPU 数目
 
-race 　　　　　　　　　　 开启编译的时候自动检测数据竞争的情况，目前只支持 64 位的机器
 
-v 　　　　　　　　　　　　 打印出来我们正在编译的包名
 
-work 　　　　   　　　　  打印出来编译时候的临时文件夹名称，并且如果已经存在的话就不要删除
 
-x 　　　　　　　　　　　　 打印出来执行的命令，其实就是和 -n 的结果类似，只是这个会执行
 
-ccflags 'arg list' 　　 传递参数给 5c, 6c, 8c 调用
 
-compiler name 　　　　   指定相应的编译器，gccgo 还是 gc
 
-gccgoflags 'arg list'   传递参数给 gccgo 编译连接调用
 
-gcflags 'arg list' 　　  传递参数给 5g, 6g, 8g 调用
 
-installsuffix suffix    为了和默认的安装包区别开来，采用这个前缀来重新安装那些依赖的包，-race 的时候默认已经是 -installsuffix race，大家可以通过 -n 命令来验证
 
-ldflags 'flag list' 　  传递参数给 5l, 6l, 8l 调用
 
-tags 'tag list' 　　　　 设置在编译的时候可以适配的那些 tag，详细的 tag 限制参考里面的 Build Constraints
```

## go clean 

这个命令是用来移除当前源码包和关联源码包里面编译生成的文件。这些文件包括

```text
_obj/        旧的 object 目录，由 Makefiles 遗留
_test/       旧的 test 目录，由 Makefiles 遗留
_testmain.go 旧的 gotest 文件，由 Makefiles 遗留
test.out     旧的 test 记录，由 Makefiles 遗留
build.out    旧的 test 记录，由 Makefiles 遗留
*.[568ao]    object 文件，由 Makefiles 遗留
 
DIR(.exe)      由 go build 产生
DIR.test(.exe) 由 go test -c 产生
MAINFILE(.exe) 由 go build MAINFILE.go 产生
*.so           由 SWIG 产生

```

参数
```text
-i 清除关联的安装的包和可运行文件，也就是通过 go install 安装的文件
-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
-r 循环的清除在 import 中引入的包
-x 打印出来执行的详细命令，其实就是 -n 打印的执行版本
```

## go doc
文档注释相关，可以搭建本地GO文档服务器，包含自己的项目注释


## go env

用于打印GO语言的环境信息，如 GOPATH 是工作区目录，GOROOT 是GO语言安装目录，
GOBIN 是通过 go install 命令生成可执行文件的存放目录（默认是当前工作区的 bin 目录下），
GOEXE 为生成可执行文件的后缀

## fix

简单的说，这是一个当GO语言版本升级之后，把代码包中旧的语法更新成新版本语法的自动化工具。
它是 go tool fix 的简单封装，它作用于代码包。
当需要升级自己的项目或者升级下载的第三方代码包，可以使用此方法。
（下载并升级代码包可以使用 go get -fix 命令 ）

## go fmt

go 工具集中提供了一个 go fmt 命令它可以帮你格式化你写好的代码文件，
使你写代码的时候不需要关心格式，你只需要在写完之后执行 go fmt <文件名>.go，
你的代码就被修改成了标准格式，但是我平常很少用到这个命令，
因为开发工具里面一般都带了保存时候自动格式化功能，这个功能其实在底层就是调用了 go fmt。

使用 go fmt 命令，其实是调用了 gofmt，而且需要参数 -w，否则格式化结果不会写入文件。
gofmt -w -l src，可以格式化整个项目。

所以 go fmt 是 gofmt 的上层一个包装的命令，我们想要更多的个性化的格式化可以参考 gofmt。

```text
-l 　　　　   显示那些需要格式化的文件
-w 　　　　   把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。
-r 　　　　   添加形如 “a [b:len (a)] -> a [b:]” 的重写规则，方便我们做批量替换
-s 　　　　　 简化文件中的代码
-d 　　　　　 显示格式化前后的 diff 而不是写入文件，默认是 false
-e 　　　　　 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印不同行的前 10 个错误。
-cpuprofile 支持调试模式，写入相应的 cpufile 到指定的文件
```

## go generate

通过扫描Go源码中的特殊注释来识别要运行的常规命令。了解go generate不是go build的一部分很重要。
它不包含依赖关系分析，必须在运行go build之前显式运行。
它旨在由Go package的作者使用，而不是其客户端。

## go get 

这个命令是用来动态获取远程代码包的，目前支持的有 BitBucket、GitHub、Google Code 和 Launchpad。

这个命令在内部实际上分成了两步操作：第一步是下载源码包，第二步是执行 go install。

下载源码包的 go 工具会自动根据不同的域名调用不同的源码工具，对应关系如下：

```text
BitBucket (Mercurial Git)
GitHub (Git)
Google Code Project Hosting (Git, Mercurial, Subversion)
Launchpad (Bazaar)
```
所以为了 go get 能正常工作，你必须确保安装了合适的源码管理工具，并同时把这些命令加入你的 PATH 中。
其实 go get 支持自定义域名的功能，具体参见 go help remote。

```text
-d 　　只下载不安装
-f 　　只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
-fix  在获取源码之后先运行 fix，然后再去做其他的事情
-t 　　同时也下载需要为运行测试所需要的包
-u 　　强制使用网络去更新包和它的依赖包
-v 　　显示执行的命令
```

## go install

这个命令在内部实际上分成了两步操作：第一步是生成结果文件 (可执行文件或者 .a 包)，第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。

参数支持 go build 的编译参数。大家只要记住一个参数 -v 就好了，这个随时随地的可以查看底层的执行信息。

## go list

不加任何标记直接使用，是显示指定包的导入路径，如 go list net/http 就显示 net/http。

该命令加上 -json 标记可以显示完整信息

## go run

编译并执行，只能作用于命令源码文件，一般用于开发中快速测试。

## go test

执行这个命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件。

## go tool

go工具，go tool pprof性能检查工具,   go tool cgo跟C语言和GO语言有关的命令

## go version

打印go版本

## go vet

静态检查工具，一般项目快完成时进行进行优化时需要
