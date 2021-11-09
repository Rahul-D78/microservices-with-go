package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//Handle err using standard go http library
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", d)
}
