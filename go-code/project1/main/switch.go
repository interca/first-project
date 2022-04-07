package main
import "fmt"
func main(){
	var a string
	fmt.Println("输入一个数")
	fmt.Scanf("%s",&a)
	switch a {
	    case "星期一":
		fmt.Println("你好星期一")
	    case "星期二":
			fmt.Println("你好星期二")
		case "星期三":
			fmt.Println("你好星期三")	
	}
	
}