package main

import "net/http"

type Action func(w http.ResponseWriter, r *http.Request) error

type BaseController struct{}

func (bc *BaseController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			http.Error(w, "Server internal error", http.StatusInternalServerError)
		}
	})
}
