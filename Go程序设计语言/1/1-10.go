package main

import (
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func check(err error){
	if(err!=nil){
		panic(err)
	}
	return 
}

func main(){
	var f *os.File
	var err error
	start:=time.Now()
	ch:=make(chan string)
	for _,url:=range os.Args[1:]{
		go fetch(url,ch)
	}
	for range os.Args[1:]{
		f,err=os.Create("info.txt")
		check(err)
		w:=bufio.NewWriter(f)
		_,err=w.WriteString(<-ch)
		check(err)
		w.Flush()
		f.Close()
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string,ch chan<- string){
	var b []byte
	start:=time.Now()
	resp,err:=http.Get(url)
	if(err!=nil){
		ch <- fmt.Sprint(err)
		return 
	}
	b,err=ioutil.ReadAll(resp.Body)
	check(err)
	ch <- fmt.Sprintf("%s",b)
	nbytes,err:=io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if(err!=nil){
		ch <- fmt.Sprintf("while reading %s: %v",url,err)
		return 
	}
	secs:=time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
}