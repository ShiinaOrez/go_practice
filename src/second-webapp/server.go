package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"hello!")
}

func world(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"world!")
}

func index(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"this is the page of shiina'go practice!\nhello!-->/hello/\nworld!-->/world/")
}

func main(){
    server:=http.Server{
        Addr: "127.0.0.1:8080",
    }
    http.HandleFunc("/",index)
    http.HandleFunc("/hello/",hello)
    http.HandleFunc("/world/",world)
    
    server.ListenAndServe()
}
