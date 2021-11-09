package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Rahul-D78/micro-go/handler"
)

func main() {

	l := log.New(os.Stdout, "Microservices in GO ", log.LstdFlags)
	hh := handler.NewHello(l)

	//converting the function into a handler type and registering it into DefaultServeMux
	//DefaultServeMux is a server multiplexer which have the logic to call the function based on the path given
	// Read More about it https://pkg.go.dev/net/http#ServeMux

	//create a new ServeMux
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	//modifying for handling blocking connections
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//trap sigterm or interput and gracefully shutdown the server
	//signal.notify is going to brodcast a message on this channel when even an operting system kill command or interrupt is recieved
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	//Block until a signal is recieved
	sig := <-sigChan
	l.Println("Recieved terminate, gracefully shutdown", sig)

	//In case of upgrade a version or some other work which needed to shut the server down can potentially loose all
	//the clients connection but doing it `gracefully` using the GO server
	//Wait until the request that are currently handled by the function have completed it will then shutdown

	//30 secs to attempt to gracefully shuts down if the handlers are still working after those 40 secs forcefully close it
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
