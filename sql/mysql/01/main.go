package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type handler struct {
	DB *sql.DB
}

var H handler

func init() {
	var err error
	H.DB, err = sql.Open("mysql", "root:@/employees")
	if err != nil {
		log.Fatal(err)
	}

	if err = H.DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

type department struct {
	deptNo, deptName string
}

func main() {

	defer H.DB.Close()

	rows, err := H.DB.Query("SELECT * FROM departments LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	result := []department{}
	for rows.Next() {
		var d department
		rows.Scan(&d.deptNo, &d.deptName)
		result = append(result, d)
	}

	fmt.Println(result)
}
