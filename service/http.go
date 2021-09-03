package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routes() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api", routeHomePage)
	rtr.HandleFunc("/api/post/create", routeCreatePost)
	rtr.HandleFunc("/api/user/sign-up", loginSignUp).Methods("POST")

	http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(rtr))
}

func routeHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header.Get("Content-type"))
	fmt.Fprintf(w, "{}")
}

func routeCreatePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Creating post")
}

func loginSignUp(w http.ResponseWriter, req *http.Request) {
	var u user
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := isUser(u.Username, u.Password)
	if id != 0 {
		/* Sign in */
		fmt.Println("Signing in", id)
	} else {
		/* Sign up */
		fmt.Println(createUser(u))
	}

	fmt.Fprintf(w, "{}")
}

// func middleWare(h http.Handler) mux.MiddlewareFunc {
// 	return func(h http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			fmt.Println("this is a test")
// 			h.ServeHTTP(w, r)
// 		})
// 	}
// }
