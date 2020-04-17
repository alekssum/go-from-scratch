package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserAgent: %v\n\n", r.UserAgent())
}
