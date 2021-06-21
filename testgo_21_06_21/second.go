package main

import "fmt"

//func main () {
//	// go 初始化列表初始化
//	source := [...]int{1, 2, 3}
//	sli := source[0:1]
//
//	fmt.Printf("sli value is %v\n", sli)
//	fmt.Printf("sli len is %v\n", len(sli))
//	fmt.Printf("sli cap is %v\n", cap(sli))
//
//	// 对切片内成员进行修改，原有数组也会进行修改
//	// 因为切片作为指向原有数组的引用
//	sli[0] = 4
//	fmt.Printf("sli value is %v\n", sli)
//	fmt.Printf("source value is %v\n", source)
//
//	// 通过 make 函数动态创建切片
//	// 在创建过程中指定切片的长度和容量
//	sli = make([]int, 2, 4)
//	fmt.Printf("sli value is %v\n", sli)
//	fmt.Printf("sli len is %v\n", len(sli))
//	fmt.Printf("sli cap is %v\n", cap(sli))
//
//	// make 函数创建的新切片都被初始化为类型的初始值
//
//	// 切片本质是一个结构体
//	// 包含三部分: address + len + cap
//	// 作为一个引用空间，该空间和元素空间完全就是两个空间，所以切片的首地址和头号元素的首地址完全不同
//}

// 直接声明新的切片，类似于数组的初始化
// 但是不需要指定大小，否则就变成了数组
func main () {
	/*
	ex := []int{1, 2, 3}
	fmt.Printf("ex value is %v\n", ex)
	fmt.Printf("ex len is %v\n", len(ex))
	fmt.Printf("ex cap is %v\n", cap(ex))
	*/

	// golang 提供 append 内置函数用于动态向切片添加成员
	// 它将返回新的切片，如果当前切片容量可以容纳更多的成员
	// 添加的操作将在切片指向的原有数组上进行，将会覆盖掉原有数组的值
	// 如果当前切片的容量不足以容纳更多的成员
	// 那么切片将会进行扩容（扩容过程类似于 C++ Vector）

	/*
	// array: arr1 arr2
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}

	// slice: sli1, sli2
	sli1 := arr1[0:2] // 长度为 2，容量为 4
	sli2 := arr2[2:4] // 长度为 2，容量为 2

	fmt.Printf("sli1 pointer is %p, len is %v, cap is %v, value is %v\n", &sli1, len(sli1), cap(sli1), sli1)
	fmt.Printf("sli2 pointer is %p, len is %v, cap is %v, value is %v\n", &sli2, len(sli2), cap(sli2), sli2)

	newSli1 := append(sli1, 5)
	fmt.Printf("newSli1 pointer is %p, len is %v, cap is %v, value is %v\n", &newSli1, len(newSli1), cap(newSli1), newSli1)
	fmt.Printf("source arr1 become %v\n", arr1)

	newSli2 := append(sli2, 5)
	fmt.Printf("newSli2 pointer is %p, len is %v, cap is %v, value is %v\n", &newSli2, len(newSli2), cap(newSli2), newSli2)
	fmt.Printf("source arr2 become %v\n", arr2)
	*/
	
	arr3 := [...]int{1, 2, 3, 4}
	sli3 := arr3[0:2:2] // 长度为 2，容量为 2

	fmt.Printf("sli3 pointer is %p, len is %v, cap is %v, value is %v\n", &sli3, len(sli3), cap(sli3), sli3)

	newSli3 := append(sli3, 5)
	fmt.Printf("newSli3 pointer is %p, len is %v, cap is %v, value is %v\n", &newSli3, len(newSli3), cap(newSli3), newSli3)
	fmt.Printf("source arr3 become %v\n", arr3)
}