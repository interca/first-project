package main
import "fmt"
func main(){
	var count int =0
	var sum int =0
	var i int=0
	for i=1;i<=100;i++{
		if(i%9==0){
			count++
			sum=sum+i
			fmt.Println(i)
		}
		
	}
    fmt.Println("你好",count,"你好",sum)
}
