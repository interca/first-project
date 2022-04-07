package main
import"fmt"
func main(){
	var s=sum(1,2)
	fmt.Println(s)
}
func sum(a int,b int)int{
	
	defer fmt.Println("a=",a) 
	defer fmt.Println("b=",b)
	a++
	b++
	var res=a+b
	fmt.Println("res=",a+b)
	return res
}