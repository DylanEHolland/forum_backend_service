package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("unicorns")

func routes() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api", routeHomePage)
	rtr.HandleFunc("/api/post/create", routeCreatePost)
	rtr.HandleFunc("/api/user/sign-up", loginSignUp).Methods("POST")

	http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(rtr))
}

func routeHomePage(w http.ResponseWriter, req *http.Request) {
	user_id := readJwt(req.Header.Get("Authorization"))
	if user_id > 0 {
		fmt.Println("Found user:", userFromId(user_id).Username)
	}

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

	var response loginResponse

	id := isUser(u.Username, u.Password)
	if id != 0 {
		/* Sign in */
		token := createJwt(id)
		response.Token = token
		fmt.Println("Signing in", token)
	} else {
		/* Sign up */
		id := createUser(u)
		token := createJwt(id)
		response.Token = token
	}

	response_json, err := json.Marshal(response)
	if err == nil {
		fmt.Printf("Error")
	}

	fmt.Fprintf(w, string(response_json))
}

func createJwt(id int) string {
	claims := userAuth{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "Forum",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secureSecretText"))
	if err != nil {
		fmt.Println("Help")
	}

	return signedToken
}

func readJwt(jwt_token string) int {
	if jwt_token == "undefined" {
		return 0
	}

	var id int
	token, err := jwt.ParseWithClaims(
		jwt_token,
		&userAuth{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secureSecretText"), nil
		},
	)

	if err == nil {
		fmt.Println("Error")
	}

	claims, ok := token.Claims.(*userAuth)
	if !ok {
		fmt.Println("Claims Error")
		id = 0
	} else {
		id = claims.Id
	}

	return id
}
