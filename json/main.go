package main

import (
	"encoding/json"
	"log"
)

var s string = `{"keys":[
		{"text":"Add word", "data":"{some new data}"},
		{"text":"Edit word", "data":"{some edit data}"}
	]}`

type tree struct {
	Keys []struct {
		Text string `json:"text"`
		Data string `json:"data"`
	} `json:"keys"`
}

func main() {

	t := &tree{}

	err := json.Unmarshal([]byte(s), t)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(t)

}
