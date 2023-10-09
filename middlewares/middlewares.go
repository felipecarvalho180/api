package middlewares

import (
	"devbook-api/helpers"
	"log"
	"net/http"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := helpers.ValidateToken(r); err != nil {
			helpers.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
