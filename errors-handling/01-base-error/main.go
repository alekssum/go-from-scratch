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

	url := "http://notexistingsite.org"

	_, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("getRemotePage\nerr:%s\nurl:%s", err, url)
	}

	return nil
}
