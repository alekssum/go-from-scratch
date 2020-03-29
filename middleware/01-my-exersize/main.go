package main

import (
	"log"
	"net/http"
)

func main() {

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminHandler := adminMiddleware(adminMux)

	siteMux := http.NewServeMux()
	siteMux.Handle("/admin/", adminHandler)
	siteMux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api"))
	})

	siteHandler := commonMiddleware(siteMux)

	http.ListenAndServe(":8080", siteHandler)

}

func adminIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin index page"))
}

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("adminMiddleware...")

		next.ServeHTTP(w, r)
	})
}

func commonMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("commonMiddleware...")

		next.ServeHTTP(w, r)
	})

}
