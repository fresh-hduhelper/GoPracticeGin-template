package main

import "fmt"

/*
	欢迎来到热身章节，从这里开始我们将带您一步步地熟悉 GitHub Classroom 的使用

	通常，每一个任务会有一个输入与输出，输入会由后台给出(给出的方式可以是直接通过
	输入流传入，也可以是执行提前在项目中设置好的脚本文件)

	您的程序需要根据传入的输入进行相应的输出，可直接传到标准输出，也可发送一个
	Web 请求，具体要求视不同任务要求来定

	任务通常不会对您的项目结构提出要求，但有些可能会检查您的项目结构是否合理以及
	您是否在指定区域进行作答，通常这种情况发生的概率很小(因为我懒得实现x)

	在开始任务前，请一定记得在命令行执行 go run build.go <相应章节>，如果您忘了这一步
	可以回到项目根目录的 README.md 中查看具体的命令 注: 这条指令必须要运行哦，而且只用运行一次:)
*/

func main() {
	// 这里只需要实现一个判断输入的数是否小于 3 的简单程序即可
	// 具体的实现已经给出了，您可以修改下面的代码使其错误，然后看看打分的进行，当然也可
	// 直接提交，您会在自己 repo (仓库)的相应位置看到一片绿的打分结果 :)

	var num int
	fmt.Scanln(&num)
	fmt.Println(num < 3) // 打印的输出将由后台判断时候正确
}
