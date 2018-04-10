package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	counts:=make(map[string]int)
	files:=os.Args[1:]
	if len(files)==0 {
		fmt.Println("Can't find any files!")
	}else{
		for _,arg:=range files{
			f,err:=os.Open(arg)
			if err!=nil {
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			tag:=countLines(f,counts)
			if tag==1 {
				fmt.Println(arg)
			}
			f.Close()
		}
	}
}

func countLines(f*os.File,counts map[string]int) int{
	tag:=0
	input:=bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
		if counts[input.Text()]>1 {
			tag=1
		}
	}
	return tag
}