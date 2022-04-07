package main

import (
	"fmt"
	"strings"
)
func main(){
   p:=makeSuffix(".jpg")
   fmt.Println(p("nihao"))
   
}
func makeSuffix(suffix string)func (string)string{
   return  func (name string)string{
	   if(strings.HasSuffix(name,suffix)==false){
		   return  name+suffix;
	   }
	   return name
   }
}
