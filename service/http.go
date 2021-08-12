package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "test page!")
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":5000", n)
}
