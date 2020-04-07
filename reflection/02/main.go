package main

import (
	"fmt"
	"reflect"
)

type user struct {
	id                         uint8
	name, middleName, lastName string
}

func main() {

	u := user{100, "Ivan", "Ivanovich", "Ivanov"}

	if reflect.TypeOf(u).Kind() == reflect.Struct {
		user := reflect.ValueOf(u)
		for i := 0; i < user.NumField(); i++ {
			fmt.Printf("Type: %s\tValue: %v\n",
				user.Field(i).Type(),
				user.Field(i),
			)
		}
	}

}
