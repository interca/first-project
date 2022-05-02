package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [26]byte

	for i := 0; i < 26; i++ {
		a[i] = byte('A' + i)
	}
	for _, w := range a {
		fmt.Printf("%c\n", w)
	}
	var num [5]int
	//生成随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		num[i] = rand.Intn(100)
	}
	fmt.Println(num)
	var i = 0
	var j = len(num) - 1
	for i < j {
		var temp = num[i]
		num[i] = num[j]
		num[j] = temp
		i++
		j--
	}
	fmt.Println(num)
	var w [5][5]int
	fmt.Println(w)
}
