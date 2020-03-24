package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	r.GET("/", HomeHandler)
	r.POST("/posts", PostsCreateHandler)
	r.GET("/hello/:name", Hello)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintln(w, "Home")

}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	r.ParseForm()
	fmt.Fprintln(w, r.FormValue("param"))
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}
