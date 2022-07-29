package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func RandomNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rand.Intn(100))
}

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	port := "3000"
	router := mux.NewRouter()
	router.HandleFunc("/", RandomNumber)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("[+] Starting number-service server on port", port)
	log.Fatal(server.ListenAndServe())
}
