package initializer

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "192.168.92.43:3306",
		DBName: "sqlidb",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("failed to connected to the database, %v",err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to verify the database connection, %v",err)
	}
}