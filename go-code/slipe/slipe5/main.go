package main

import "fmt"

//冒泡
func main() {
	var a = []int{1, 5, 6, -1, 4, 8, 22, 14, -23, 7}
	for i := 0; i < len(a)-1; i++ {
		var flag = 1
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				flag = 0
				var temp = a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
		if flag == 1 {
			break
		}
	}
	fmt.Println(a)
}
