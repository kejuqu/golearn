## 目录说明

golearn 下的每一个目录对应 [go docs](https://go.dev/doc/) 中的一个小章节
比如：
greetings 对应 [create a module](https://go.dev/doc/tutorial/create-module)

## Start go

### install go

[install link](https://go.dev/doc/install)

### setup go mod

```
#creating new go.mod: module example/hello
go mod init example/hello
```

### write code

```
// touch hello.go, then paste the below code
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

### run code

```
go run hello.go
```

### import other package

- [search go package](https://pkg.go.dev/)
- import packageName
- go mod tidy; go will add the package module as a requirement, as well as a go.sum file for use in authenticating the module

- 通过运行 go mod tidy 命令，可以在开发过程中保持项目的依赖关系清晰和最新。它可以帮助你删除未使用的依赖项，更新依赖项的版本，并确保项目的 go.mod 文件与实际使用的依赖项保持同步
- go.sum 是一个用于确保模块依赖项的版本安全的文件

#### import local package 需要注意

如果引入的 package 是本地的，需要执行如下命令

```
go mod edit -replace xxx/packageName=../packageName
```

## quick jump for go

- [go-start](https://go.dev/doc/tutorial/getting-started)
- [standard library](https://pkg.go.dev/std)
- [search go package](https://pkg.go.dev/)
- [go command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [Compile and install the application](https://go.dev/doc/tutorial/compile-install)
- [go workspace](https://go.dev/ref/mod#workspaces)

### go 常用指令

- go mod init moduleName // 初始化一个 module
- go run xxx.go // 运行 xxx.go 可执行文件
- go mod tidy // 整理和更新项目的模块依赖关系
- go fmt // 格式化 go 代码
- go build // 编译包及其以来想，但是不会 install, 能运行编译后的可执行文件，需要你指定可执行文件的路径或者在可执行的文件目录中运行它； 如果需要在没有指定可执行文件路径的情况下执行它，则需要使用到用到 install
- go list -f '{{.Target}}' 发现 go 的安装路径
- go install // 编译并且下载 packages
- go help // 查看更多有关 go 的指令

## 学习疑问 QA

> Go 代码被分组到包中，而包被分组到模块中

### go mod tidy 的作用

1. 检查项目的 go.mod 文件中列出的依赖项，并根据实际代码中使用的情况，确定需要保留的依赖项和版本。

2. 移除 go.mod 文件中未被实际代码引用的依赖项。

3. 更新 go.mod 文件中依赖项的版本，以便与实际使用的版本一致。

4. 如果有需要，下载缺失的依赖项。

### 编译和下载应用

1. `go build` // 编写好代码之后, 在当前应用目录下 构建应用，编译包和对应的依赖
2. 在当前应用目录下执行项目 ./projectName
3. `go list -f '{{.Target}}'`(假设结果为 projectPath ) // 在当前项目下运行此命令，发现安装的路径
4. Add the Go install directory to your system's shell path

```
export PATH=$PATH:(projectPath值)
```

5. `go install` 编译和下载应用
6. `$projectName` 运行应用

### go.sum 的作用

> go.sum 是一个用于确保模块依赖项的版本安全的文件,在使用 Go 模块管理工具时，建议将 go.sum 文件纳入版本控制，并确保它与项目的依赖项保持同步。

1. 安全性验证：go.sum 文件中记录了每个模块的预期哈希值。当你构建或下载依赖项时，Go 会验证实际下载的模块文件的哈希值是否与 go.sum 中记录的一致。这样可以确保所使用的依赖项的版本是预期的，并且没有被篡改。

2. 版本锁定：go.sum 文件中还记录了每个依赖项的确切版本。这有助于锁定项目的依赖项，确保团队成员或构建服务器在构建项目时使用相同的依赖项版本，从而提供更一致的构建结果。

3. 加速构建：go.sum 文件中的哈希值可以使 Go 编译器在构建项目时更快地判断依赖项是否已更改。如果依赖项的哈希值与 go.sum 中的记录一致，Go 编译器可以跳过重新下载和验证依赖项，从而提高构建速度。

### log 的作用

1. log.SetPrefix("title: ") 设置日志记录器的前缀：
   通过调用 log.SetPrefix 函数并传递一个字符串参数，可以设置日志记录器输出消息的前缀。
   在这种情况下，设置的前缀为 "title: "，意味着所有使用该日志记录器记录的消息都将以 "title: " 开头。
   log.SetFlags(0) 设置日志记录器的标志：

2. log.SetFlags 函数用于设置日志记录器输出消息的标志。
   传递参数为 0，表示不设置任何标志。这意味着输出的日志消息中不会包含日期、时间或其他额外信息。
   log.Fatal(err) 记录错误消息并终止程序：

3. log.Fatal 函数用于记录错误消息并终止程序的执行。
   传递一个错误对象 err 作为参数，该函数会将错误消息记录到日志记录器中，并调用 os.Exit(1) 终止程序的执行。数字 1 表示程序以非零状态退出。

### 开发是遇到的问题

1. vscode 开发环境没有提示
   运行 `go get -u -v github.com/mdempsky/gocode`

2. 在 multi-module workspaces 下单个 module 无代码提示 & 不能跳转到 go 文件
   解决方法：可擦考 [go multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
   比如我当前的 workspace 的名字叫做 golearn, 如果我要让
