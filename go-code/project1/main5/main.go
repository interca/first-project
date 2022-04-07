package main
import "fmt"
func main(){
	var i,n int
	var j int
	fmt.Scanln(&n)
	for i=0;i<=n;i++{
		j=n-i
		fmt.Println(i,"+",j,"=",n)
	}
}