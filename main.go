package main

import (
	"embed"

	"sqli/initializer"
	// "sqli/controllers"
	// "sqli/views"
)

var (
	//go:embed .env
	envFile embed.FS
)

func init() {

}

func main() {
	// Router()

	// address := "0.0.0.0:5001"
	// http.ListenAndServe(address, nil)
	// log.Printf("Listening on %v\n", address)
	initializer.LoadEnv(envFile)
	initializer.ConnectDB()
	initializer.MigrateDB()
}
