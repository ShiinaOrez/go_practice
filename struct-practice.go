package main

import "fmt"

type like_some struct{
    star map[string]int
//    like=make(map[string]int)
    name string;
    age  int;
}

var slice []like_some

func main(){
//    dave:=like_some{["fff"]=23,"dave",39}
    var ss string
    var num int
    john:=like_some{}
    fmt.Printf("please input one's name:")
    fmt.Scanf("%s",&john.name)
    fmt.Printf("please input %s's age:",john.name)
    fmt.Scanf("%d",&john.age)
    john.star=make(map[string]int)
    fmt.Println("?=?")
    fmt.Scanf("%s %d",&ss,&num)
    john.star[ss]=num
//    fmt.Println(dave.name,dave.age)
    slice=append(slice,john)
    fmt.Println(john.name,john.age,john.star)
    
}
