package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sqli/initializers"
)

type User struct {
	ID       int
	Username string
	Password string
}

func VulnLogin(username, password string) (User, error) {
	queryString := fmt.Sprintf("SELECT * FROM credentials WHERE username='%s' AND PASSWORD='%s'", username, password)
	log.Printf("queryString: %v\n", queryString)
	rows, err := initializers.DB.Query(queryString)
	if err != nil {
		log.Printf("error occured while trying to login, %v", err)
		return User{}, err
	}
	defer rows.Close()

	user := User{}
	if rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password)
		log.Printf("user: %#v\n", user)
	} else {
		return User{}, err
	}
	return user, nil
}

func SecureLogin(username, password string) (User, error) {
	var err error
	var stmt *sql.Stmt
	preparedString := "SELECT * FROM credentials where username=? AND password=?"
	stmt, err = initializers.DB.Prepare(preparedString)
	log.Printf("preparedString: %v\n", preparedString)
	if err != nil {
		log.Printf("error occured in prepare while trying to login, %v", err)
		return User{}, errors.New("something went wrong please try again")
	}
	defer stmt.Close()

	row := stmt.QueryRow(username, password)
	user := User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Printf("error occured in query while trying to login, %v", err)
		return User{}, errors.New("invalid credentials")
	} else {
		log.Printf("user: %#v\n", user)
	}
	return user, nil
}

func VulnChangePassword(username, oldPassword, newPassword string) error {
	queryString := fmt.Sprintf("UPDATE credentials SET password='%s' WHERE username='%s' AND password='%s'", newPassword, username, oldPassword)
	log.Printf("queryString: %v\n", queryString)
	result, err := initializers.DB.Exec(queryString)
	if err != nil {
		log.Printf("error occured while trying to change password, %v", err)
		return err
	}

	var rowsAffected int64
	if rowsAffected, err = result.RowsAffected(); err != nil {
		log.Printf("error occured while trying to get rows affected, %v", err)
		return err
	}
	log.Printf("rowsAffected: %v\n", rowsAffected)

	return nil
}
