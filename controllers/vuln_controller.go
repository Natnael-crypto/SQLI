package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"sqli/models"
	"sqli/views"
)

func VulnController(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}
	views.LoginRender(&buf)
	fmt.Fprint(w, buf.String())
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := models.VulnLogin(username, password)
		fmt.Printf("user in controller: %#v\n", user)
		fmt.Fprintf(w, "body: %v&%v\n", user.Username, user.Password)

	}
}
