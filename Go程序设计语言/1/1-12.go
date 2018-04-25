package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	_"os"
	_"time"
	"net/http"
	"net/url"
	"log"
	"strconv"
)

var palette=[]color.Color{color.White,color.Black}

const (
	whiteIndex=0
	blackIndex=1
)

var m map[string]float64=map[string]float64{
	"cycles" : 5 ,
	"res" : 0.001 ,
	"size" : 100 ,
	"nframes" : 64 ,
	"delay" : 8 ,
}
//m:=make(map[string]int)
 

func check(e error){
	if e!=nil{
		panic(e)
	}
}

func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000",nil))
}

func handler(w http.ResponseWriter,r *http.Request){
	err:=r.ParseForm();
	check(err)
//	for k,v:=range r.Form{
//		fmt.Fprintf(w, "Form[%q]=%q\n",k,v)
//	}
//	fmt.Fprintf(w,"%q",r.URL.RawQuery)
	m,_:=url.ParseQuery(r.URL.RawQuery)
	for k,v:=range m{
		if _,ok:=m[k];ok{
			vv,_:=float64(strconv.Atoi(string(v)));
			m[k]=vv
			fmt.Println("%s",k)
		}
	}
	lissajous(w)
}

func lissajous(out io.Writer){
	freq:=rand.Float64()*3.0
	anim:=gif.GIF{LoopCount: int(m["nframes"])}
	phase:=0.0
	for i:=0;i<int(m["nframes"]);i++{
		rect:=image.Rect(0,0,2*int(m["size"])+1,2*int(m["size"])+1)
		img:=image.NewPaletted(rect,palette)
		for t:=0.0;t<m["cycles"]*2*math.Pi;t+=m["res"]{
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(int(m["size"])+int(x*m["size"]+0.5),int(m["size"])+int(y*m["size"]+0.5),blackIndex)
		}
		phase+=0.1
		anim.Delay=append(anim.Delay,int(m["delay"]))
		anim.Image=append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim)
}
