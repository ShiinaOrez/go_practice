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
    min := 99999
    from := 1
    to := 1
    for i:=0;i<n;i++ {
        for j:=i+1;j<n;j++ {
            tot:=0
            for k:=0;k<4;k++ {
                tot+=(slice[i].star[kind[k]]-slice[j].star[kind[k]])*(slice[i].star[kind[k]]-slice[j].star[kind[k]])
            }
            tot+=((slice[i].age-slice[j].age)*(slice[i].age-slice[j].age))
//            fmt.Println(tot,i,j)
            if tot<min {
                min=tot
                from = i
                to= j
            }
        }
    }
    fmt.Println("%s and %s will be the good frined!",slice[from].name,slice[to].name)
}
