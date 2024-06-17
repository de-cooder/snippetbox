package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	te, err := template.ParseFiles("./ui/html/home.tmpl")
	if err != nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	} 
	err = te.Execute(w,nil)
	if(err != nil){
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	w.Write([]byte("Hello from Snippetbox"))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}
