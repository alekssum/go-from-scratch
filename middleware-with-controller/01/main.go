package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Starting server...")
	u := &UserController{}

	apMux := http.ServeMux{}
	apMux.Handle("/api/user/greet", u.Action(u.Greet))
	handler := LogsMiddleware(apMux)

	log.Println("Listening port 8080")
	http.ListenAndServe(":8080", handler)

}

type UserController struct {
	BasicController
}

func (u *UserController) Greet(w http.ResponseWriter, r *http.Request) error {

	w.Write([]byte("Hi from controller"))

	return nil

}

func LogsMiddleware(next http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Some logs")

		next.ServeHTTP(w, r)
	})
}