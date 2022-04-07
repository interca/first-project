package main
import "fmt"
func main(){
	var i,j int
     var k int
	for i=1;i<10;i++{
		for k=0;k<10-i;k++{
			fmt.Print(" ")
		}
		for j=0;j<2*i-1;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	for i=1;i<10;i++{
		for k=0;k<10-i;k++{
			fmt.Print(" ")
		}
		for j=0;j<2*i-1;j++{
			if((j==0||j==2*i-2)||i==9){
				fmt.Print("*")
			}else if(i!=9){
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	
}