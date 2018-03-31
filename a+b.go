package main

import "fmt"

func main(){
    a,b:= 0,0
    fmt.Println("please input two numbers at the next line:")
    fmt.Scanf("%d %d",&a,&b)
    fmt.Printf("--->%d\n",a+b)
}
