package main

import "fmt"

type like_some struct {
    name string;
    age int;
    star map[string]int
}

var slice []like_some
var kind = [5]string {"movie","food","study","excrise"}

func main(){
    var n int
    fmt.Printf("please input the numbers of people:")
    fmt.Scanf("%d",&n)
    one:=like_some{}
    for i:=1;i<=n;i++ {
        fmt.Printf("one's name:")
        fmt.Scanf("%s",&one.name)
        fmt.Printf("%s's age:",one.name)
        fmt.Scanf("%d",&one.age)
        one.star=make(map[string]int)
        for j:=0;j<4;j++ {
            ss:=kind[j]
            var num int
            fmt.Printf("please input how do you like %s (1 to 5):",ss)
            fmt.Scanf("%d",&num)
            one.star[ss]=num
        }
        slice=append(slice,one)
    }
    fmt.Print(slice)
}
