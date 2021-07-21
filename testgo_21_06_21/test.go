package test
// go 文件必须属于一个包
import "fmt"

// go 语言指针的使用特点
func testPtr(num *int) {
	*num = 20
	fmt.Println(*num)
}

// 天然并发
// 从语言层面支持并发，实现简单
// goroutine，轻量级线程，可实现大并发处理，高效利用多核
// 基于 CPS 并发模型实现
// 吸收了管道通信机制，通过管道channel，可以实现不同goroutine之间的相互通信
// 函数可以返回多个值
// 新的创新，比如切片slice，延时执行 defer 等

// 一个函数，同时返回求和和差
// go 函数可以支持返回多个值
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2 // go 语句后面不需要带分号
	sub := n1 - n2
	return sum, sub
}

func main() {
	testPtr();
}
