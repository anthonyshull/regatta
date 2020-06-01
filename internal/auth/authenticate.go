package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"net/http"
)

var LoginURI = "/login"

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == LoginURI {
			next.ServeHTTP(w, r)
		}

		header := r.Header.Get("authorization")

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	login := gocql.PasswordAuthenticator{}

	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&login)

	fmt.Println(login)
}
