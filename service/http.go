package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func routes() {
	mux := http.NewServeMux()

	// mw := jwtmiddleware.New(jwtmiddleware.Options{
	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("secret"), nil
	// 	},
	// 	SigningMethod: jwt.SigningMethodHS256,
	// })

	mux.HandleFunc("/api", routeHomePage)
	mux.HandleFunc("/api/post/create", routeCreatePost)

	r := negroni.Classic() // Includes some default middlewares
	r.UseHandler(mux)
	//r := router

	http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
}

func routeHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header.Get("Content-type"))
	fmt.Fprintf(w, "{}")
}

func routeCreatePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Creating post")
}

func middleWare(h http.Handler) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("this is a test")
			h.ServeHTTP(w, r)
		})
	}
}
