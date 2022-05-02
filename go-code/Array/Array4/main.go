package main

import "fmt"

func main() {
	var a = make([][]int, 3)
	fmt.Println(a)
	for i := 0; i < 3; i++ {
		a[i] = make([]int, 2)
	}
	fmt.Println(a)
}
