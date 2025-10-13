package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: loadRoutes(),
	}

	log.Print("Starting Server")
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
