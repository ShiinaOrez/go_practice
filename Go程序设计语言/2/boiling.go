package main

import (
	"fmt"
)

const boilingF=212.0

func main(){
	f:=boilingF
	c:=(f-32)*5/9
	fmt.Printf("boiling point = %g °F or %g °C\n",f,c)
}