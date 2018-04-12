package main

import(
	"fmt"
	_"io/ioutil"
	"net/http"
	"io"
	"os"
	"strings"
)

const (
    prefix="https://"
)

func main(){
	for _,url:=range os.Args[1:]{
		if(!strings.HasPrefix(url,prefix)){
			url=prefix+url
		}
		resp,err:=http.Get(url)
		if(err!=nil){
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			os.Exit(1)
		}
		n,err:=io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if(err!=nil){
			fmt.Fprintf(os.Stderr,"fetch:reading %s: %v\n",url,err)
		}
		fmt.Printf("%d\n",n)
	}
}

