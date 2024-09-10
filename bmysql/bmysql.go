package bmysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
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
	DB, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}