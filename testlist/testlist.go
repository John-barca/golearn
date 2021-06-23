package main

import (
	"fmt"
)

// go 列表通过双向链表实现
// 可高效进行插入删除操作

/*
func main () {
	tmplist := list.New()

	for i := 1; i <= 10; i++ {
		tmplist.PushBack(i)
	}

	first := tmplist.PushFront(0)
	tmplist.Remove(first)

	for i := tmplist.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ")
	}

	// Golang 提供映射的关系容器为 map，内部通过散列表实现
	// map 使用 make 函数初始化
	classMates1 := make(map[int]string)
	classMates1[0] = "john"
	classMates1[1] = "mike"
	classMates1[2] = "johnny"

	fmt.Printf("id %v is %v\n", 1, classMates1[1])

	// 声明时初始化数据
	classMates2 := map[int]string {
		0 : "john",
		1 : "mike",
		2 : "johnny",
	}

	fmt.Printf("id %v is %v\n", 3, classMates2[3])
}
*/

// go 中的 range 关键字用于 for 中迭代 array、slice、channel 和 map 元素
// 遍历对于很多 golang 内置容器来说，形式基本一致
func main () {
	// array 
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range nums {
		// k: index, v: value
		fmt.Println(k, v, " ")
	}

	fmt.Println()

	// slice 
	slis := []int{1, 2, 3, 4, 5}
	for k, v := range slis {
		fmt.Println(k, v, " ")
	}

	fmt.Println()

	// map 
	tmpMap := map[int]string {
		0 : "john",
		1 : "mike", 
		2 : "johnson",
	}

	for k, v := range tmpMap {
		fmt.Println(k, v, " ")
	}

	// for-range 遍历中
	// key 和 value 都是通过拷贝的方式赋值
	// 对它们进行修改不会影响到容器内成员变化
}