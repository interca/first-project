package main
import "fmt"
func main(){
    var i,j int;
	for i=1;i<=9;i++{
		for j=1;j<=i;j++{
			fmt.Print(j," * ",i," = ",i*j,"\t")
		}
		fmt.Print("\n")
	}
}