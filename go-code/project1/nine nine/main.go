package main

import "fmt"

func main() {
	var i, j int
	var k int
	for i = 1; i <= 9; i++ {
		k = i
		for j = 1; j <= i; j++ {
			fmt.Print(j, "*", k, "=", k*j, " ")

		}
		fmt.Print("\n")
	}
}
