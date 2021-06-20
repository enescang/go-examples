package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Birthdate string
	IsActive  bool
}

var DB *sql.DB
var err error

func main() {
	DB, err = sql.Open("mysql", "root:@/go_db")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer DB.Close()

	// allRecords()
	// oneRowWithID(6)
	//insertRow(User{Username: "github", Email: "test.github@test.com", FirstName: "foo", LastName: "bar", Birthdate: "2006", IsActive: true})
	deleteRowWithID(9)
}

func allRecords() {
	rows, err := DB.Query("Select * from users")
	if err != nil {
		panic(err.Error())
	}
	var user User

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Birthdate, &user.IsActive)
		if err != nil {
			log.Fatal("Error at allRecords: ", err)
		}
		fmt.Println(user)
	}
}

func oneRowWithID(id int) {
	var user User
	err := DB.QueryRow("Select * from users where ID = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Birthdate, &user.IsActive)
	if err != nil {
		panic(err.Error())
	}
	//we can use user.Email or any other fields
	fmt.Println(user)
}

func insertRow(u User) {
	res, err := DB.Exec("Insert Into users(Username, Email, Password, FirstName, LastName, Birthdate, IsActive) values(?, ?, ?, ?, ?, ?, ?)", u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.Birthdate, u.IsActive)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := res.LastInsertId()
	fmt.Println("One Row Inserted. ID:", id)
}

func deleteRowWithID(id int) {
	res, err := DB.Exec("Delete From users Where ID=?", id)
	if err != nil {
		log.Fatal(err)
	}

	total, _ := res.RowsAffected()
	fmt.Println(total, " Row Deleted.")
}
