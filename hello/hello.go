package main // package main 表示该文件属于 main 包。这是一个特殊的包名，用于指定可执行程序的入口

/*
golearn
	- example.com/greetings 是一个模块引用路径，是声明在 greetings下的 go.mod 中的
	- example.com/hello 是一个模块引用路径，是声明在 hello下的 go.mod 中的

hello 引用本地 go 模块（greetings）需要执行：
1. go mod edit -replace example.com/greetings=../greetings
2. go mod tidy
3. go run hello.go 将会成功调用 greetings.Hello 方法

 因为 example.com/greetings 还没有发布，只是在 local 本地的，因此需要在 hello 文件夹路径下执行：

 go mod edit -replace example.com/greetings=../greetings , 执行之后 hello/go.mod 中将会新增一行 replace example.com/greetings => ../greetings
 即将不存在的路径 （example.com/greetings) 重定向到 本地目录（模块所在的地方 ../greetings)

 syntax: [go mod edit --replace notExistPath = existPath]

*/
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
