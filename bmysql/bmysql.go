package bmysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  float64 `json:"price"`
}

func InitializeConnection() {
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
}

func InsertAlbum(a Album) {
	query := "INSERT INTO album (title,artist,price) VALUES(?,?,?)"
	runQuery(query, a.Title, a.Artist, a.Price)
}

func DeleteAlbum(id int) {
	query := "DELETE FROM album where id= ?"
	runQuery(query, id)
}

func UpdateAlbum(id int, a Album) {
	query := "UPDATE album SET title = ?, artist = ?, price= ?  where id= ?"
	runQuery(query, a.Title, a.Artist, a.Price, id)
}

func GetAllAlbums() []Album {
	query := "SELECT title,artist,price FROM album"
	albums, err := getRows(query)
	if err != nil {
		log.Fatal(err)
	}
	return albums
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

func getRows(query string) ([]Album, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error when preparing SQL statement: %s", err)
		return []Album{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("Error when executing SQL statement: %s", err)
		return []Album{}, err
	}
	defer rows.Close()

	var albums = []Album{}
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.Title, &album.Artist, &album.Price); err != nil {
			return []Album{}, err
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return []Album{}, err
	}

	return albums, nil
}