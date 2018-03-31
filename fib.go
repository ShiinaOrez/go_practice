package main

import "fmt"

func main(){
    var fib=[4] int {0,1,1}
    var n int
    fmt.Scanf("%d",&n);
    if (n<=2) {
        fmt.Println("1")
        return 
    }
    for i:=3;i<=n;i++ {
        fib[i%2+1]=fib[1]+fib[2]
    }
    fmt.Printf("%d\n",fib[n%2+1])
}
