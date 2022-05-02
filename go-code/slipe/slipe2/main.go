package main

import "fmt"

//切片声明方式
func main() {
	//大小，容量(可选）
	var a []int = make([]int, 4, 5)
	fmt.Println(a)
	var b []int = make([]int, 10)
	fmt.Println(cap(b))
	var c = "黄aaa裕甲"
	for i := 0; i < len(c); i++ {
		fmt.Printf("%c", c[i])
	}
}
