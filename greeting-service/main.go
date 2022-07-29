package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func GoodbyeWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Goodbye, world!")
}

func main() {
	port := "3000"
	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld)
	router.HandleFunc("/hello", HelloWorld)
	router.HandleFunc("/goodbye", GoodbyeWorld)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("[+] Starting greeting-service server on port", port)
	log.Fatal(server.ListenAndServe())
}
