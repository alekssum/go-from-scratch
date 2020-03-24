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

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintln(w, "Home")

}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")

	fmt.Println("getParam:", p.ByName("getParam"))
	fmt.Println("postParam:", p.ByName("postParam"))
	fmt.Println("formValueGetParam:", r.PostFormValue("getParam"))
	fmt.Println("formValuePostParam:", r.PostFormValue("postParam"))
}
