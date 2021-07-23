/*************************************************************************
 * Trying out Go, this is a simple greeting HTTP microservice.
 * Make a request to "/greeting" and it returns the customary greeting.
 *************************************************************************/
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GreetingDto struct {
	Greeting string `json:"greeting"`
}

func handleGetGreeting(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content","application/json")
	err := json.NewEncoder(writer).Encode(GreetingDto{Greeting: "Hello, world!"})
	if err != nil {
		return
	}
}

func greetingRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Path("/greeting").Methods("GET").HandlerFunc( handleGetGreeting )
	return r
}

func main() {
	r := greetingRoutes()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return 
	}
}
