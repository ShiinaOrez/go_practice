package main

import (
	"time"
	"strings"
	"fmt"
	"os"
)

func main(){
	s1_1:=time.Now()
	s,sep:="",""
	for _,arg:=range os.Args[1:]{
		s+=sep+arg  
		sep=" "
	}
	s1:=time.Since(s1_1).Seconds()
	s2_1:=time.Now()
	ss:=strings.Join(os.Args[1:]," ")
	s2:=time.Since(s2_1).Seconds()
	fmt.Println(s1," ",s2)
	fmt.Println(ss)
}