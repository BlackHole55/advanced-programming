package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MeirhanSyzdykov/Assignment2/internal/api"
	"github.com/MeirhanSyzdykov/Assignment2/internal/queue"
	"github.com/MeirhanSyzdykov/Assignment2/internal/store"
	"github.com/MeirhanSyzdykov/Assignment2/internal/worker"
)

func main() {
	store := store.NewMemoryStore()
	queue := queue.NewTaskQueue(100)

	worker.NewWorkerPool(2, queue.Channel(), store)

	handler := api.NewHandler(store, queue)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handler.Tasks)
	mux.HandleFunc("/tasks/", handler.GetTask)
	mux.HandleFunc("/stats", handler.Stats)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("Server running on :8080")
		log.Fatal(server.ListenAndServe())
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	queue.Close()
	server.Shutdown(ctx)
}
