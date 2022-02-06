package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project2/handler"
)

func main() {
	log.Println("starting main for project2...")

	router := handler.SetRoutes()
	port := "0.0.0.0:8081"
	listenAndServe(router, port)
}

func listenAndServe(router *gin.Engine, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		log.Printf("listening on address: %s", port)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s", err)
		}
	}()
	//listen for the system INTERRUPT signals
	<-ctx.Done()

	//Restore default behavior on the interrupt signal and notify user of shutdown
	stop()
	log.Println("shutting down gracefully...")

	//the context is used to tell server that it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exiting")
}
