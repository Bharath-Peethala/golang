package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var table = "album"

type album struct {
	title  string
	artist string
	price  float64
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "2019",
		Net:    "tcp",
		Addr:   "127.0.0.1:3307",
		DBName: "recordings",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	InsertAlbum(album{title: "Indra", artist: "Mani Sharma", price:58.25})
	DeleteAlbum(1)
	UpdateAlbum(2, album{title: "OG", artist: "Thaman", price: 100.25})
}

func InsertAlbum(a album) {
	query := "INSERT INTO album (title,artist,price) VALUES(?,?,?)"
	runQuery(query, table, a.title, a.artist, a.price)
}

func DeleteAlbum(id int) {
	query := "DELETE FROM album where id= ?"
	runQuery(query, id)
}

func UpdateAlbum(id int, a album) {
	query := "UPDATE album SET title = ?, artist = ?, price= ?  where id= ?"
	runQuery(query, a.title, a.artist, a.price, id)
}

func runQuery(query string, args ...interface{}) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error when preparing SQL statement: %s", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Printf("Error when executing SQL statement: %s", err)
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error when finding rows affected: %s", err)
		return
	}

	log.Printf("%d Albums affected ", rows)
}
