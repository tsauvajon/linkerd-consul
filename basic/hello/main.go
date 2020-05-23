package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)

	fmt.Println("Example app listening on port 8100!")
	log.Fatalln(http.ListenAndServe(":8100", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am healthy!")
}
