package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Print("request")
	io.WriteString(w, "hello, world!\n")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", HelloServer)
	err := http.ListenAndServe(":12345", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
