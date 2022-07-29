package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

func modifyRequest(req *http.Request, servicePrefix string) {
	// Remove service prefix from URL in proxied request
	updatedUrlString := strings.ReplaceAll(req.URL.String(), servicePrefix, "")
	modifiedURL, err := url.Parse(updatedUrlString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+] Modified URL:", modifiedURL)
	req.URL = modifiedURL
}

func NewProxy(target string, servicePrefix string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req, servicePrefix)
	}
	//proxy.ModifyResponse = modifyResponse()
	//proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func HelloGateway(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, gateway!")
}

func main() {
	port := "8080"

	greetingProxy, err := NewProxy("http://host.docker.internal:8001", "/greeting")
	if err != nil {
		panic(err)
	}

	numberProxy, err := NewProxy("http://host.docker.internal:8002", "/number")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", HelloGateway)
	router.PathPrefix("/greeting/").HandlerFunc(ProxyRequestHandler(greetingProxy))
	router.PathPrefix("/number/").HandlerFunc(ProxyRequestHandler(numberProxy))

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("[+] Starting gateway-service server on port", port)
	log.Fatal(server.ListenAndServe())
}
