package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func init() {
	http.HandleFunc("/", index)
	go http.ListenAndServe(":8888", nil)
}
func main() {
	fmt.Println("hello,world")
	select {}
}
