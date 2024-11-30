package controllers

import (
	"log"
	"net/http"
	"os"
	"sqli/models"
	"sqli/views"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	Vuln     string = "vuln"
	Secure   string = "secure"
	Products string = "/products"
	Login    string = "/login"
	Admin    string = "/admin"
)

func LoginController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		views.LoginRender(w)
	case http.MethodPost:
		var (
			user        models.User
			err         error
			tokenString string
		)
		req.ParseForm()
		username := req.FormValue("username")
		password := req.FormValue("password")
		action := req.FormValue("action")
		if action == Vuln {
			user, err = models.VulnLogin(username, password)
		} else {
			user, err = models.SecureLogin(username, password)
		}
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			views.LoginRender(w, err)
		} else {
			tokenExpiry := time.Now().Add(time.Minute * 5)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"isAdmin":  user.IsAdmin,
				"sub":      user.Username,
				"exp":      tokenExpiry.Unix(),
			})

			tokenString, err = token.SignedString([]byte(os.Getenv("JWTSECRET")))
			if err != nil {
				log.Printf("error occured in login controller while trying to login: %v\n", err)
				return
			}

			cookie := http.Cookie{
				Name:     "Authorization",
				Value:    tokenString,
				Expires:  tokenExpiry,
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
			}

			http.SetCookie(w, &cookie)

			// http.Redirect(w, req, Products, http.StatusFound)
			if user.IsAdmin {
				http.Redirect(w, req, Admin, http.StatusFound)
			} else {
				http.Redirect(w, req, Products, http.StatusFound)
			}
			log.Printf("tokenString: %v\n", tokenString)
			log.Printf("valid credentials: %v&%v\n", user.Username, user.Password)
		}

	}
}

func ChangePasswordController(w http.ResponseWriter, req *http.Request) {
	hasErrorMsg := false
	switch req.Method {
	case http.MethodGet:
		views.ChangePasswordRender(w, hasErrorMsg)
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

		if action == Vuln {
			err = models.VulnChangePassword(username, oldPassword, newPassword)
		} else {
			err = models.SecureChangePassword(username, oldPassword, newPassword)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			hasErrorMsg = true
			views.ChangePasswordRender(w, hasErrorMsg)
		} else {
			http.Redirect(w, req, Login, http.StatusFound)
			views.ChangePasswordRender(w, hasErrorMsg)
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

		if action == Vuln {
			err = models.VulnForgotPassword(username)
		} else {
			err = models.SecureForgotPassword(username)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			views.ForgotPasswordRender(w, data[1])
		} else {
			views.ForgotPasswordRender(w, data[0])
		}
	}
}
