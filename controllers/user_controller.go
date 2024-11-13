package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sqli/models"
	"sqli/views"
)

const (
	vuln   string = "vuln"
	secure string = "secure"
)

func LoginController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		views.LoginRender(w)
	case http.MethodPost:
		var (
			user models.User
			err  error
		)
		req.ParseForm()
		username := req.FormValue("username")
		password := req.FormValue("password")
		login := req.FormValue("login")
		if login == vuln {
			user, err = models.VulnLogin(username, password)
		} else {
			user, err = models.SecureLogin(username, password)
		}
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, err)
		} else {
			// TODO: Redirect to products page
			fmt.Fprintf(w, "body: %v&%v\n", user.Username, user.Password)
		}

	}
}

func ChangePasswordController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		views.ChangePasswordRender(w)
	case http.MethodPost:
		var err error
		req.ParseForm()
		username := req.FormValue("username")
		oldPassword := req.FormValue("oldPassword")
		newPassword := req.FormValue("newPassword")
		change_password := req.FormValue("change_password")

		log.Printf("username: %v\n", username)
		log.Printf("oldPassword: %v\n", oldPassword)
		log.Printf("newPassword: %v\n", newPassword)

		if change_password == vuln {
			err = models.VulnChangePassword(username, oldPassword, newPassword)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
		}

	}
}
