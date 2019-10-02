package main

import (
	"ecosia-application/app/api"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	serveRequest(api.GetRoutes())
}

func serveRequest(configuredRoutes http.Handler) {
	log.Print("########## SERVER STARTED ##########")
	error := http.ListenAndServe(
		":8000",
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
		)(configuredRoutes))
	if error != nil {
		fmt.Print("Error during server startup: ", error)
	}
	os.Exit(1)
}
