package main

import(
    "fmt"
    "strconv"
)

func check(e error){
    if e!=nil {
        panic(e)
    }
}

func main(){
//    var s string
//    fmt.Scanf("%s",s)
    a,err:=strconv.Atoi("600")
    check(err)
    fmt.Print(a)
}

