package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func main() {
	err := getRemotePage()
	if err != nil {
		switch err := errors.Cause(err).(type) {
		case *url.Error:
			log.Printf("Resourse %s, err: %s", err.URL, err.Err)
		default:
			log.Println(err)
		}
		return
	}
	log.Println("Everything is ok")
}

func getRemotePage() error {

	url := "http://notexistingsite.org"

	_, err := http.Get(url)

	if err != nil {
		return errors.Wrap(err, "Resource error")
	}

	return nil
}
