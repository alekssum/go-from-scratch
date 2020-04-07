package main

import (
	"fmt"
	"reflect"
)

type order struct {
	id    uint8
	date  string
	time  uint64
	goods []string
}

type user struct {
	id                         uint8
	name, middleName, lastName string
}

func main() {

	o := &order{
		id:    100,
		date:  "2020-05-25",
		time:  190230217340,
		goods: []string{"Макароны", "Хлеб", "Молоко"},
	}

	if reflect.TypeOf(*o).Kind() == reflect.Struct {
		value := reflect.ValueOf(*o)
		for i := 0; i < value.NumField(); i++ {
			fmt.Println(value.Field(i))
		}
	}

}
