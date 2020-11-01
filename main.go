package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sideroff/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "test ", log.LstdFlags)

	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()

	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)

	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Println("Received an interrupt, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
