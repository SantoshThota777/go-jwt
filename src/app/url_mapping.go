package app

import "github.com/SantoshThota777/go-jwt/src/controller/login"

func appURLs() {
	router.HandleFunc("/getMessage", login.LoginPageHandler).Methods("GET")
}
