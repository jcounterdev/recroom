// This will be the main file for the recroom server program.

package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	port := "1337"

	fmt.Println("hello world")
	
	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}