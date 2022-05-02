package main

import "fmt"

//append切片扩容
func main() {
	var a = []int{1, 2, 3, 4}
	fmt.Println(a)
	//用append追加
	a = append(a, 100, 200, 300)
	fmt.Println(a)
	//追加自己
	a = append(a, a...)
	fmt.Println(a)
	var b = make([]int, 15)
	copy(b, a)
	fmt.Println(b)

}
