package main

import (
	"fmt"
	"/home/shiina/go_practice/Golang/2/tempconv"
	"os"
	"strconv"
)

func check(e error){
	if e!=nil {
		panic(e)
	}
}

func main(){
	for _,arg:=range os.Args[1:]{
		t,err:=strconv.ParseFloat(arg,64)
		check(err)
		f:=tempconv.F(t)
		C:=tempconv.C(t)
//		k:=tempconv.K(t)
		fmt.Printf("%s=%s , %s=%s , %s=%s",f,tempconv.FToC(f),c,tempconv.CTOf(c))
	}
}
