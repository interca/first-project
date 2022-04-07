package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().Unix())
	var a int
	var count=0
	
	for{
		a=rand.Intn(100)
		count++
		if(a==99){
			break;
		}
	}
	fmt.Println(count)
}