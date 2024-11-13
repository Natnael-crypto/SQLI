package controllers

import (
	"fmt"
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
			fmt.Fprintf(w, "body: %v&%v\n", user.Username, user.Password)
		}

	}
}
