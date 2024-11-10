package initializer

import (
	"database/sql"
	"embed"
	"io"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	//go:embed *.sql
	initSQL embed.FS
)

func ConnectDB() {
	cfg := mysql.Config{
		User:            os.Getenv("DBUSER"),
		Passwd:          os.Getenv("DBPASS"),
		Net:             "tcp",
		Addr:            os.Getenv("DBADDR"),
		DBName:          os.Getenv("DBNAME"),
		MultiStatements: true,
	}
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("failed to connected to the database, %v", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to verify the database connection, %v", err)
	}
}

func MigrateDB() {
	var (
		err      error
		contents []byte
	)

	sqlFile, err := initSQL.Open("init.sql")
	if err != nil {
		log.Fatalf("error while opening init file in migration, %v", err)
	}

	contents, err = io.ReadAll(sqlFile)
	if err != nil {
		log.Fatalf("error while reading init file in migration, %v", err)
	}

	_, err = DB.Query(string(contents))
	if err != nil {
		log.Fatalf("error while populating data in migration, %v", err)
	}
}
