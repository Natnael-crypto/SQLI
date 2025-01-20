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
	IsAdmin  bool
}

var (
	InvalidCredentialsErr = errors.New("Invalid Credentials")
	SomethingWentWrongErr = errors.New("Something went wrong please try again")
)

func SecureLogin(username, password string) (User, error) {
	var (
		err  error
		stmt *sql.Stmt
	)
	preparedString := "SELECT * FROM credentials where username=? AND password=?"
	log.Printf("preparedString: %v\n", preparedString)
	stmt, err = initializers.DB.Prepare(preparedString)
	if err != nil {
		log.Printf("error occured in prepare while trying to login, %v", err)
		return User{}, SomethingWentWrongErr
	}
	defer stmt.Close()

	row := stmt.QueryRow(username, password)
	user := User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin)
	if err != nil {
		log.Printf("error occured in query while trying to login, %v", err)
		return User{}, InvalidCredentialsErr
	} else {
		log.Printf("user: %#v\n", user)
	}
	return user, nil
}

func SecureChangePassword(username, oldPassword, newPassword string) error {
	var (
		err          error
		stmt         *sql.Stmt
		result       sql.Result
		rowsAffected int64
	)
	preparedString := fmt.Sprintf("UPDATE credentials SET password=? WHERE username=? AND password=?")
	log.Printf("preparedString: %v\n", preparedString)
	stmt, err = initializers.DB.Prepare(preparedString)
	if err != nil {
		log.Printf("error occured in prepare while trying to change password, %v", err)
		return SomethingWentWrongErr
	}
	defer stmt.Close()

	result, err = stmt.Exec(newPassword, username, oldPassword)
	if err != nil {
		log.Printf("error occured in query while trying to change password, %v", err)
		return InvalidCredentialsErr
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Printf("error occured while trying to get rows affected, %v", err)
		return SomethingWentWrongErr
	} else if rowsAffected == 0 {
		return InvalidCredentialsErr
	}
	log.Printf("rowsAffected: %v\n", rowsAffected)

	return nil
}

func SecureForgotPassword(username string) error {
	var (
		err  error
		stmt *sql.Stmt
	)
	preparedString := fmt.Sprintf("SELECT * FROM credentials where username=?")
	log.Printf("preparedString: %v\n", preparedString)

	stmt, err = initializers.DB.Prepare(preparedString)
	if err != nil {
		log.Printf("error occured in prepare while trying to look up user for password reset, %v", err)
		return SomethingWentWrongErr
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	user := User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin)
	if err != nil {
		log.Printf("error occured in Query while trying to look up user for password reset, %v", err)
		return InvalidCredentialsErr
	}

	return nil
}
