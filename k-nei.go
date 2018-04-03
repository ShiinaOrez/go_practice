package main

import "fmt"

type like_some struct{
//    var like map [string]int
//    like=make(map[string]int)
    name string;
    age  int;
}

func main(){
    dave:=like_some{dave,39}
    john:=like_some
    fmt.Printf("please input one's name")
    fmt.Scanf("%s",&john.name)
    fmt.Printf("please input %s's age:",john.name)
    fmt.Scanf("%d",&john.age)
    fmt.Println(dave.name,dave.age)
    fmt.Println(john.name,john.age)
    
}
