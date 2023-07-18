package main //定义一个包，声明包名为main，表明当前是一个可执行程序（Go通过包来管理明明空间）

import "fmt" //导入一个外部包fmt

func main() { // main函数，是程序执行的入口函数
	fmt.Println("Go Hello World!") //在终端打印出Go Hello World!
}
