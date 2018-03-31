package main

import "fmt"

var n int

func print_stars(num int){
    for i:=1;i<=num;i++{
        fmt.Print("*")
    }
    fmt.Println("")
    if(num==n){
        return
    }
    print_stars(num+1)
    for i:=1;i<=num;i++{
        fmt.Print("*")
    }
    fmt.Println("")
    return
}

func main(){
    fmt.Scanf("%d",&n)
    print_stars(1)
}
