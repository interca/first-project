package main

import "fmt"

func main() {
	//string可以进行切片处理
	var str = "abcdefe"
	fmt.Println(str)
	var a = str[3:]
	fmt.Println(a)
	//修改字符串，把字符串先变成切片
	var b = []byte(str)
	b[3] = 'A'
	str = string(b)
	fmt.Println(str)
	fmt.Println(a)
}
