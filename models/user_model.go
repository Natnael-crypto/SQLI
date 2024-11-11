package models

import (
	"fmt"
	"log"
	"sqli/initializer"
)

type User struct {
	ID       int
	Username string
	Password string
}

func VulnLogin(username, password string) User {
	queryString := fmt.Sprintf("SELECT * FROM credentials WHERE username='%s' AND PASSWORD='%s'", username, password)
	rows, err := initializer.DB.Query(queryString)
	if err != nil {
		log.Fatalf("error occured while trying to login, %v", err)
	}
	defer rows.Close()

	user := User{}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password)
		fmt.Printf("user: %#v\n", user)
	}

	return user
}
