package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	//http.HandleFunc("/form", formhandler)// TO DO : زمانی که فرانت کامل شد
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/helloMehdi" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is nit suported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}
