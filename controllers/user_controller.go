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
		err = req.ParseForm()
		if err != nil {
			log.Printf("error occurred while parsing form data: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			views.LoginRender(w)
			return
		}

		username := req.FormValue("username")
		password := req.FormValue("password")
		action := req.FormValue("action")

		if action == Vuln {
			err = http.ErrNoCookie
		} else {
			user, err = models.SecureLogin(username, password)
		}
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			views.LoginRender(w, err)
		} else {
			tokenExpiry := time.Now().Add(time.Minute * 30)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"isAdmin":  user.IsAdmin,
				"sub":      user.Username,
				"exp":      tokenExpiry.Unix(),
			})

			tokenString, err = token.SignedString([]byte(os.Getenv("JWTSECRET")))
			if err != nil {
				log.Printf("error occurred in login controller while trying to generate token: %v\n", err)
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

			if user.IsAdmin {
				http.Redirect(w, req, Admin, http.StatusFound)
			} else {
				http.Redirect(w, req, Products, http.StatusFound)
			}
		}
	}
}

func ChangePasswordController(w http.ResponseWriter, req *http.Request) {
	hasErrorMsg := false
	switch req.Method {
	case http.MethodGet:
		views.ChangePasswordRender(w, hasErrorMsg)
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			log.Printf("error occurred while parsing form data: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			views.ChangePasswordRender(w, true)
			return
		}

		username := req.FormValue("username")
		oldPassword := req.FormValue("oldPassword")
		newPassword := req.FormValue("newPassword")
		action := req.FormValue("action")

		if action == Vuln {
			err = http.ErrNoCookie
		} else {
			err = models.SecureChangePassword(username, oldPassword, newPassword)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			hasErrorMsg = true
			views.ChangePasswordRender(w, hasErrorMsg)
		} else {
			http.Redirect(w, req, Login, http.StatusFound)
		}
	}
}

func ForgotPasswordController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		views.ForgotPasswordRender(w)
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			log.Printf("error occurred while parsing form data: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			views.ForgotPasswordRender(w, struct {
				IsSuccess bool
				IsFail    bool
			}{false, true})
			return
		}

		username := req.FormValue("username")
		action := req.FormValue("action")

		data := []struct {
			IsSuccess bool
			IsFail    bool
		}{{true, false}, {false, true}}

		if action == Vuln {
			err = http.ErrNoCookie
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

func LogoutController(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{
		Name:  "Authorization",
		Value: "",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, req, Login, http.StatusFound)
}
