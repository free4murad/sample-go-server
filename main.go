package main

import (
	"context"
	"go-run/handler"
	"go-run/storage/in_memory"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var (
	port = "8080" // This would be taken from a config file
)

func main() {
	// Load any config eg server post, db credentials etc

	ss := in_memory.NewUserDB() //We can also prove credentials here to

	svr := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: handler.GetHandler(ss),
	}

	var wg sync.WaitGroup
	go func() {
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)
		sig := <-sigquit
		wg.Add(1)
		_ = svr.Shutdown(context.Background())
		wg.Done()
		log.Printf("Shutting down server because of - %+v", sig)
	}()

	if err := svr.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error starting server - %v", err)
	} else {
		log.Println("Closed all listeners and idle connections, will wait for pending requests to complete")
		wg.Wait()
		log.Println("Server closed")
	}
}
