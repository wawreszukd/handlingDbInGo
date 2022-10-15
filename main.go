package main

import (
	"databaseTesting/dbHandler"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := dbHandler.InitiateDb()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Db.Close()

	err = db.Insert(dbHandler.Person{
		FirstName:   "John",
		Lastname:    "Jobs",
		Gender:      "MAN",
		DateOfBirth: "1999-01-01",
	})
	if err != nil {
		fmt.Println(err)
	}

	err = db.Select()
	if err != nil {
		fmt.Println(err)
	}

}
