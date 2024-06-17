package main

import (
	"fmt"
	"log"
	
	"net/http"
)

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/view", snippetView)
	http.HandleFunc("/snippet/create", snippetCreate)
	log.Println("Starting server")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
