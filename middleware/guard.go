package middleware

import "net/http"

type Guard struct {
	controller func()
}

func (g Guard) ServeHTTP(http.ResponseWriter, *http.Request) {
}
