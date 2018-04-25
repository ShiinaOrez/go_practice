package main

import (
    "fmt"
)

func main(){
    var i,j int 
    fmt.Scanf("%d %d",&i,&j)
    i,j=j,i
    fmt.Printf("%d %d\n",i,j)
}
