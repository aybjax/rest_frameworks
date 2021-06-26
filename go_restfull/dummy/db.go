package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)


type Book struct {
	id int
	name string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)

		return
	}

	log.Printf("%#v", db)
	fmt.Println()

	// create table
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY,
		isbn INTEGER NULL,
		author VARCHAR(64) NULL,
		name VARCHAR(64) NULL
	)`)

	if err != nil {
		log.Println("Error creating table")
		log.Fatal(err)

		return
	} else {
		log.Println("Successfully created table")
	}

	statement.Exec()

	// create
	statement, _ = db.Prepare(`INSERT INTO books (name, author, isbn)
				VALUES (?, ?, ?)`)

	statement.Exec("A Tail of Two Cities", "Charles Dickens", 140430547)

	time.Sleep(2 * time.Second)

	// read
	rows, _ := db.Query("SELECT id, name, author FROM books")

	var tempBook Book
	
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book: %s, Author: %s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	time.Sleep(2 * time.Second)

	// update
	statement, _ = db.Prepare("UPDATE books SET name=? where id=?")

	statement.Exec("Aybjax", 1)

	log.Println("Successfully updated the book in database!")

	time.Sleep(2 * time.Second)

	// delete
	statement, _ = db.Prepare("DELETE FROM books WHERE id=?")

	statement.Exec(1)

	log.Println("Successfully deleted the book in the database!")
}