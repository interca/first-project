package main

import "fmt"

//for-range遍历
func main() {
	var a [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, value := range a {
		fmt.Println(value)
	}
	fmt.Println("-----------------------------")
	//数组是值传递，除非传指针
	change1(a)
	fmt.Println(a)
	change2(&a)
	fmt.Println(a)
}
func change1(a [8]int) { //传值
	a[0] = 1000
}
func change2(a *[8]int) { //传指针
	(*a)[1] = 1000
}
