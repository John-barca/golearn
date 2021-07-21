package main

import (
	"fmt"
	"time"
)

// chan：通道，可以理解为队列，遵循先进先出
// go：后面加一个函数，可以创建一个线程，函数可以为已经写好的函数，也可以匿名

/*
func main () {
	fmt.Println("main start")

	go func() {
		fmt.Println("goroutine")
	}()
	fmt.Println("main end")
}
*/

// 结果没有 goroutine
// main 函数是一个主线程，因为主线程执行太快，子线程没来得及执行，所以看不到输出
// 因此主线程休眠一秒钟

/*
func main() {
	fmt.Println("main start")

	go func() {
		fmt.Println("goroutine")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main end")
}
*/

// 使用chan
// 声明不带缓冲的通道
// ch1 := make(chan string)
// 声明带 10 个缓冲的通道
// ch2 := make(chan string, 10)
// 声明只读通道
// ch3 := make(<-chan string)
// 声明只写通道
// ch4 := make(chan<- string)
// 不带缓冲的通道，进和出都会阻塞
// 带缓冲的通道，进一次长度+1，出一次长度-1，如果长度等于缓冲长度时，再进就会阻塞

// 写入
// ch1 := make(chan string, 10)
// ch1 <- "a"
// 读取
// val, ok := <- ch1
// 或
// val := <- ch1
// 关闭
// close(chan)

// 通道 close 以后不能再写入，写入会出现 panic
// 重复 close 会出现 panic
// 只读的 chan 不能 close
// close 以后还可以读取数据

/*
func main () {
	fmt.Println("main start")
	ch := make(chan string, 1)
	ch <- "a" // 写入通道
	go func() {
		val := <- ch // 读取数据
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
*/

/*
func main () {
	fmt.Println("main start")
	ch := make(chan string)
	go func() {
		ch <- "a"
	}()
	go func() {
		val := <- ch
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
*/

// 带缓冲的通道
// 如果长度等于缓冲长度，再进就会阻塞
/*
func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func main () {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
*/

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func customer(ch chan string) {
	for {
		msg := <- ch
		fmt.Println(msg)
	}
}

func main() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	go customer(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}