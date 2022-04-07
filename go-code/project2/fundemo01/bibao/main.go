package main
import "fmt"
func main(){
  var f=add()
  fmt.Println(f(1))
  fmt.Println(f(1))
  
}
func add()func (int)int{
  var n int=10
  var str="hello"
  fmt.Println("this:",n)
  return func (x int)int{
       n=n+x
	   str=str+"a"
	   //fmt.Println(str)
	   return n
  }
}