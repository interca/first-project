package main

import "fmt"

//在go中数组是值传递
func main() {
	//var a = make([]int, 10)第二种方法创建数组
	var a [5]int
	fmt.Println(a)
	change(a)
	fmt.Println(a)
	var b = [...]int{1, 2, 3, 5, 6, 7, 8}
	fmt.Println(b)
}
func change(a [5]int) {
	a[0] = 1000
}
