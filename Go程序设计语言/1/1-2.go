package main

import (
	"fmt"
	"os"
)

func main(){
	i:=1
	for _,arg:=range os.Args[1:]{
		fmt.Println(arg,i)
		i+=1
	}
}