package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go http.ListenAndServe(":8888", nil)
}
func main() {
	fmt.Println("hello,world")
	select {}
}
