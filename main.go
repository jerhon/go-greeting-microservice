package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GreetingDto struct {
	Greeting string `json:"greeting"`
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(GreetingDto{Greeting: "Hello, world!"})
	if err != nil {
		return
	}
}

func getRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Path("/greeting").Methods("GET").HandlerFunc(handleHello)
	return r
}

func main() {
	r := getRoutes()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
