package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var schema = `
CREATE TABLE IF NOT EXISTS book
(
	id INT NOT NULL AUTO_INCREMENT,
	titlte VARCHAR(30) NOT NULL,
	author VARCHAR(30) NOT NULL,
	page_count INT NOT NULL,
	PRIMARY KEY (id)
)
`

type Book struct {
	Id        int
	Title     string
	Author    string
	PageCount int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(schema)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()

	//map to values
	var id int
	var title string
	var author string
	var pageCount int

	rows, _ := db.Query("SELECT * from book")

	for rows.Next() {
		rows.Scan(&id, &title, &author, &pageCount)
		fmt.Printf("id=%v, title=%v, author=%v, pageCount=%v\n",
			id,
			title,
			author,
			pageCount)
	}

	//map to struct
	rows, _ = db.Query("SELECT * from book")
	b := Book{}
	for rows.Next() {
		rows.Scan(&b.Id, &b.Title, &b.Author, &b.PageCount)
		fmt.Printf("book = %v\n", b)
	}

	fmt.Println("Ping to database succesfull !")

	//Insert INTO (prepared statements)
	b = Book{
		Title:     "Le langage C",
		Author:    "R. Stallman",
		PageCount: 310,
	}

	stmt, _ = db.Prepare("INSERT INTO book (title, author, page_count) VALUES(?, ?, ?)")
	_, err = stmt.Exec(b.Title, b.Author, b.PageCount)
	if err != nil {
		log.Fatal(err)
	}
}
