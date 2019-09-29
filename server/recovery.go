package server

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"
)

func HandleError(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if rec := recover(); rec != nil {
				var err error
				switch t := rec.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				log.Fatal("Exception in call to: ", r.URL.RequestURI(), " : ", err)
				debug.PrintStack()
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
