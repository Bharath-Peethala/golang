package entities

import (
	"context"
	"log"
	"time"
	"golang/bmysql"
)

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func InsertAlbum(a Album) (int64, error) {
	query := "INSERT INTO album (title,artist,price) VALUES(?,?,?)"
	rows, err := runQuery(query, a.Title, a.Artist, a.Price)
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}

func DeleteAlbum(id int) int64 {
	query := "DELETE FROM album where id= ?"
	rows, err := runQuery(query, id)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func UpdateAlbum(id int, a Album) (int64,error) {
	query := "UPDATE album SET title = ?, artist = ?, price= ?  where id= ?"
	rows, err := runQuery(query, a.Title, a.Artist, a.Price, id)
	if err != nil {
		log.Fatal(err)
	}
	return rows,err
}

func GetAllAlbums() []Album {
	query := "SELECT title,artist,price FROM album"
	albums, err := getRows(query)
	if err != nil {
		log.Fatal(err)
	}
	return albums
}

func runQuery(query string, args ...interface{}) (int64, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := bmysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error when preparing SQL statement: %s", err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Printf("Error when executing SQL statement: %s", err)
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error when finding rows affected: %s", err)
		return 0, err
	}

	log.Printf("%d Albums affected ", rows)
	return rows, nil
}

func getRows(query string) ([]Album, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := bmysql.DB.PrepareContext(ctx, query)
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