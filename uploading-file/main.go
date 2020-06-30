package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Content-Type:", r.Header.Get("Content-Type"))
		log.Println("Body:", string(body))

		w.Write(body)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
