package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//http handle func convert the function that you are passing as a parameater into a http handler
//http handler is a interface with a single method  https://pkg.go.dev/net/http#Handler
//we are creating a struct which implements the http handler

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//Request is a struct type so we are using * there and ResponseWriter is an interface
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//we have implemented dependency injection so that the NewHello function which takes logger
	//we can replace Logger with anything else

	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//Handle err using standard go http library
		http.Error(w, "Ooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", d)
}
