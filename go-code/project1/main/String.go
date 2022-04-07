package main
import "fmt"
//import "unsafe"
func main(){
	var a int =20
	var b float64=1.2
	var c int ='a'
	var d bool=true
	var str string 
	str =fmt.Sprintf("%d",a);
	fmt.Println("str=",str);
	str=fmt.Sprintf("%f",b)
	fmt.Println("str=",str)
	str=fmt.Sprintf("%c",c)
	fmt.Println("str=",str)
	str=fmt.Sprintf("%t",d)
	fmt.Println("str=",str)
	fmt.Printf("类型是%T",a)
	
}