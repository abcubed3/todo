package http

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func (srv *Server) Run() {
	srv.registerRoutes()
	log.Println("starting server....")
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	// wait = 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func (srv *Server) registerRoutes() {
	srv.r.HandleFunc("/{id:[0-9a-zA-Z]+}", srv.get).Methods("GET")
	srv.r.HandleFunc("/{id:[0-9a-zA-Z]+}", srv.updateDetail).Methods("PUT")
	srv.r.HandleFunc("/", srv.create).Methods("POST")
}
