package main

import (
	"fmt"
	"log"
	"net/http"
)

type OwnError struct {
	URL string
	Err error
}

func (e *OwnError) Error() string {
	return fmt.Sprintf(
		"Resourse error:\nURL %s\nerr: %s",
		e.URL,
		e.Err,
	)
}

func main() {
	err := getRemotePage()
	if err != nil {
		switch err.(type) {
		case *OwnError:
			// type assertion
			e := err.(*OwnError)

			log.Printf("Resource error: URL %s\n", e.URL)
			log.Printf("%s", e)
		default:
			log.Println("internal error")
		}
		return
	}
	log.Println("Everything is ok")
}

func getRemotePage() error {

	url := "http://notexistingsite.org"

	_, err := http.Get(url)

	if err != nil {
		return &OwnError{URL: url, Err: err}
	}

	return nil
}
