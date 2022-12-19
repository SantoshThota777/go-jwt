package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartAppilcation() {
	fmt.Println("starting application")
	appURLs()

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Fatal("Something wrong starting the service", err)
	}
}
