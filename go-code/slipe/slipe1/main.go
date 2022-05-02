package main

import "fmt"

//切片，可以无限扩充，类似于ArrayList
func main() {
	var a [8]int = [...]int{1, 2, 3, 4, 5, 6, 7, 5}
	//定义切片
	//切片引用数组1到3下标，不包括3
	var slide = a[0:3]
	fmt.Println(a)
	fmt.Println(slide)
	//切片长度
	fmt.Println(len(slide))
	//切片容量
	fmt.Println(cap(slide))
	slide[0] = 100
	fmt.Println(a)
}
