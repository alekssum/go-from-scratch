package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := getRemotePage()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Everything is ok")
}

func getRemotePage() error {

	_, err := http.Get("http://notexistingsite.org")

	if err != nil {
		return fmt.Errorf("getRemotePage: %s", err)
	}

	return nil
}
