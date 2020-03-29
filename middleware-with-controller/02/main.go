package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var handler http.Handler
	mainPage := &mainPage{}
	apiMux := http.ServeMux{}

	apiMux.Handle("/", mainPage.Action(mainPage.Index))
	apiMux.Handle("/error/", mainPage.Action(mainPage.ErrorPage))

	handler = mwBefore(apiMux)

	http.ListenAndServe(":8080", handler)

}

// Controllers
type mainPage struct {
	basicController
}

func (m *mainPage) Index(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Welcome!"))
	return nil
}

func (m *mainPage) ErrorPage(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Something is wrong") 
}

// Middlewares
func mwBefore(next http.ServeMux) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Println("before")
		fmt.Fprintln(w, "before")
		next.ServeHTTP(w, r)	
	})	
}

// Basic controller
type action func (w http.ResponseWriter, r *http.Request) error
	
type basicController struct {}

func (bc *basicController) Action(a action) http.Handler  {

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {

		if err := a(w, r); err != nil {

			http.Error(w, "500 Internal server error", http.StatusInternalServerError)
			log.Println(err)

		}
		
	})
	
}