package main

import "fmt"

func main() {
	var a = find(20)
	fmt.Println(a)
}
func find(n int) []uint64 {
	var a = make([]uint64, n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a
}
