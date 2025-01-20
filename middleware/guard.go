package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sqli/initializers"
	"sqli/models"
	"sqli/views"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func Guard(controller func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("Guarding")
		var (
			err    error
			cookie *http.Cookie
			token  *jwt.Token
			user   models.User
		)

		cookie, err = req.Cookie("Authorization")
		if err != nil {
			log.Printf("no Auth cookie found, %v", err)
			views.ErrorRender(w, "401 Unauthorized")
			return
		}
		tokenString := cookie.Value
		log.Printf("tokenString: %v\n", tokenString)

		token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWTSECRET")), nil
		})
		if err != nil {
			log.Printf("An errror occured in guard: %v", err)
			views.ErrorRender(w, "401 Unauthorized")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Printf("sub: %v, exp: %v", claims["sub"], claims["exp"])

			// check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				views.ErrorRender(w, "401 Unauthorized")
				return
			}

			// find the user with token sub
			row := initializers.DB.QueryRow("SELECT * FROM credentials WHERE username = ?", claims["sub"])
			err = row.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin)
			if err != nil {
				log.Printf("An errror occured while trying to get user in guard: %v", err)
				views.ErrorRender(w, "401 Unauthorized")
				return
			}

			// attach to request
			userCookie := http.Cookie{Name: "User", Value: user.Username}
			isAdminCookie := http.Cookie{Name: "isAdmin", Value: strconv.FormatBool(user.IsAdmin)}

			req.AddCookie(&userCookie)
			req.AddCookie(&isAdminCookie)

			// authenticate
			controller(w, req)
		} else {
			log.Printf("An errror occured in guard: %v", err)
			views.ErrorRender(w, "401 Unauthorized")
			return
		}
	}
}
