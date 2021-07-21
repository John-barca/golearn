package main

import (
	"fmt"
)

// defer: 在声明时不会立刻执行
// 而是在函数 return 后去执行
// 主要场景有异常处理、记录日志、清理数据、释放资源等
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
func main() {
	x := 1
	y := 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4
}
*/

/*
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("main")
}
*/
// defer 函数定义的顺序 与 实际执行的顺序是相反的，也就是最先声明的最后才执行

// 闭包test
/*
func main() {
	var a = 1
	var b = 2
	defer fmt.Println(a + b)

	a = 2

	fmt.Println("main")
}
// main
// 3
*/

/*
func main() {
	var a = 1
	var b = 2
	defer func() {
		fmt.Println(a + b)
	}()
	a = 2
	fmt.Println("main")
}
*/
// main
// 4
// 闭包获取变量相当于引用传递，而非值传递

func main() {
	var a = 1
	var b = 2

	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)

	a = 2
	fmt.Println("main")
}
// main 
// 3
// 传参是值拷贝
// defer 调用的函数，参数的值在 defer 定义时就确定了
// 而 defer 函数内部所使用的变量的值需要在这个函数运行时才确定