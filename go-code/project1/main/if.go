package main
import "fmt"
func main(){
	var i int
	fmt.Println("输入一个数")
	fmt.Scanf("%d",&i)
	if(i>10){
		fmt.Println("大于十")
	}else{
		fmt.Println("小于十")
	}
}