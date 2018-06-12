package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
)

type user struct {
	username string
	password string
	//indent   []byte
	ident uuid.UUID
}

var db *sql.DB

//var tpl *template.template

func main() {

	//var id uuid.UUID

	//id = uuid.New()

	//fmt.Printf("%T\n", id)

	fmt.Println("Starting...")

	dbinit()
	//insertUser()
	readUsers()

	err := db.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Ending")
}

func dbinit() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:r5aHVw42123pg@localhost/chat?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func readUsers() {
	fmt.Println("readUsers()")
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//make a slice of user structs
	us := make([]user, 0)

	for rows.Next() {
		//allocate a user struct
		u := user{}

		err := rows.Scan(&u.username, &u.password, &u.indent)

		if err != nil {
			panic(err)
		}

		//Add the user struct to the slice
		us = append(us, u)
	}

	fmt.Printf("%v\n", us)
}

/*
func insertUser() {
	fmt.Println("insertUser()")
	rows, err := db.Query("SELECT username FROM users WHERE username=?", "test")
	if err != nil {
		//panic(err)
		//rows.Close()
		_, err = db.Exec("INSERT INTO users (username, password) VALUES (\"test\", \"1234\")")
		if err != nil {
			panic(err)
		}
		fmt.Println("INSERT")
	}

	defer rows.Close()
}
*/
