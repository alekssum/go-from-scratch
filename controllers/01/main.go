package main

import (
	"fmt"
	"log"
	"net/http"
)

type MainPageController struct {
	BaseController
}

func (mc *MainPageController) mainPage(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Main page")
	return nil
}

type TestControllers struct {
	BaseController
}

func (c *TestControllers) ShowGetParams(w http.ResponseWriter, r *http.Request) error {

	r.ParseForm()

	fmt.Fprintln(w, r.Form)

	return nil
}

func main() {
	log.Println("Starting server...")

	main := &MainPageController{}
	test := &TestControllers{}

	http.Handle("/", main.Action(main.mainPage))
	http.Handle("/test/", test.Action(test.ShowGetParams))

	log.Println("Listening port 8080")
	http.ListenAndServe(":8080", nil)

}
