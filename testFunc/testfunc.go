package main

import "fmt"

// golang 中函数名第一位不能是数字
// 同包内，函数名不可重名
// 一个函数如果希望包外代码访问，函数名首字母需要大写
// 如果相邻参数类型相同，可以省略类型
func cal(a, b int) int {
	return a + b;
}

// Golang 函数支持多返回值
// 还支持对返回值进行命名
func div (dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return 
}
// 使用命名返回值函数中，函数结束前需要显式使用 return 语句返回
// 命名返回值和非命名返回值不能混合使用，否则出现编译错误

/*
// 匿名函数可以在声明之后直接调用它
func (name string) {
	fmt.Println("Name is ", name)
}("aaaaaa")
// 声明匿名函数之后，在后加上调用的参数列表
// 可立即对匿名函数调用，还可以将匿名函数赋值给函数类型变量
// 用于多次调用或求值

currentTime := func () {
	fmt.Println(time.Now())
}
// 调用匿名函数
currentTime()
*/

/*
// 匿名作回调函数
func proc (input string, processor func(str string)) {
	// 调用匿名
	processor(input)
}

func main() {
	proc("aaa", func (str string) {
		for _, v := range str {
			fmt.Printf("%c\n", v)
		}
	}) 
}
*/

// 闭包是携带状态的函数，将函数内部和外部连接起来的桥梁
// 通过闭包，可以读取函数内部的变量也可以使用闭包封装私有状态

// 闭包能引用其作用域上部变量进行修改，被捕获到闭包中的变量随着闭包生命周期一直存在
// 函数本身不存储信息，但是闭包的变量使得闭包本身具备存储信息能力
// 计数器

func createCounter (initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	// 引用变量 initial，创建闭包
	return func() int {
		initial++
		// 返回计数
		return initial
	}
}

func main () {
	// 计数器 1
	c1 := createCounter(1)

	fmt.Println(c1())
	fmt.Println(c1())

	// 计数器 2
	c2 := createCounter(10)

	fmt.Println(c2())
	fmt.Println(c1())
}