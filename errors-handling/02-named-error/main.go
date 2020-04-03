package main

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrPage = errors.New("page error")
)

func main() {

	err := getRemotePage()
	if err != nil {
		switch err {
		case ErrPage:
			log.Printf("remote sourse error: %s", ErrPage)
		default:
			log.Println(err)
		}
	}

}

func getRemotePage() error {

	url := "http://notexistingsite.org"

	_, err := http.Get(url)

	if err != nil {
		return ErrPage
	}

	return nil
}
