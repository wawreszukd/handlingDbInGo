package dbHandler

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "password"
	Dbname   = "test"
)

type MyDb struct {
	Db *sql.DB
}
type Person struct {
	ID        int
	FirstName string

	Lastname    string
	Gender      string
	DateOfBirth string
}

func InitiateDb() (*MyDb, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return &MyDb{Db: db}, nil
}
func (db MyDb) Select() error {
	sqlStatement := "SELECT * FROM person;"

	rows, err := db.Db.Query(sqlStatement)
	if err != nil {
		return err
	}

	defer rows.Close()

	var person Person

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.FirstName, &person.Lastname, &person.Gender, &person.DateOfBirth)
		if err != nil {
			return err
		}
		log.Println(person)
	}

	return nil
}

func (db MyDb) Insert(person Person) error {
	sqlStatement := "INSERT INTO person (first_name, last_name, gender, date_of_birth) VALUES ($1, $2, $3, $4);"

	_, err := db.Db.Exec(sqlStatement, person.FirstName, person.Lastname, person.Gender, person.DateOfBirth)
	if err != nil {
		return err
	}

	return nil
}
