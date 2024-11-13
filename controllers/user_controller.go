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
		action := req.FormValue("action")
		if action == vuln {
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
		action := req.FormValue("action")

		log.Printf("username: %v\n", username)
		log.Printf("oldPassword: %v\n", oldPassword)
		log.Printf("newPassword: %v\n", newPassword)

		if action == vuln {
			err = models.VulnChangePassword(username, oldPassword, newPassword)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
		}

	}
}

func ForgotPasswordController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		views.ForgotPasswordRender(w)
	case http.MethodPost:
		var err error
		req.ParseForm()
		username := req.FormValue("username")
		action := req.FormValue("action")

		// Used in render to decide wheteher or not to show error msgs.
		data := []struct {
			IsSuccess bool
			IsFail    bool
		}{{true, false}, {false, true}}

		if action == vuln {
			err = models.VulnForgotPassword(username)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			views.ForgotPasswordRender(w, data[1])
		} else {
			views.ForgotPasswordRender(w, data[0])
		}
	}
}
