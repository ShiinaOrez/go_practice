package main

import "fmt"

func main(){
    const PI float64= 3.1415926535
    var r int
    var S float64
    fmt.Scanf("%d",&r)
    S=PI*float64(r)*float64(r)
    fmt.Printf("%5f\n",S)
}
