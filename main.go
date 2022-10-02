package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lukasz0707/go-microservices/handlers"
)

func main() {
	addr := flag.String("addr", ":9090", "HTTP network address")
	flag.Parse()
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         *addr,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	l.Println("starting server on", *addr)
	err := s.ListenAndServe()
	if err != nil {
		l.Fatalf("error starting server: %s\n", err)
	}

}
