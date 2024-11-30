package initializers

import (
	"database/sql"
	"embed"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	//go:embed *.sql
	initSQL embed.FS
)

// ConnectDB attempts to connect to the database with retries.
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
	var retries = 5 // Max number of retries
	var delay = 2 * time.Second // Delay between retries

	for i := 0; i < retries; i++ {
		DB, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Printf("Error opening connection (attempt %d/%d): %v", i+1, retries, err)
		} else {
			err = DB.Ping()
			if err == nil {
				log.Println("Successfully connected to the database.")
				return
			} else {
				log.Printf("Ping failed (attempt %d/%d): %v", i+1, retries, err)
			}
		}

		// Retry logic: wait before trying again
		if i < retries-1 {
			log.Printf("Retrying in %v...", delay)
			time.Sleep(delay)
		}
	}

	// If we exhausted all retries
	log.Fatalf("Failed to connect to the database after %d attempts, exiting...", retries)
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
