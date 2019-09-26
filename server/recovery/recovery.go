package recovery

import (
	"errors"
	"net/http"
	"runtime/debug"
)

/**
* Recovery methods are how Go handles recovery from panics.
* This handler can be registered together with others to allow recovery in
* the HTTP server.
 */
func HandleRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			debug.PrintStack()
			if r != nil {
				switch t := r.(type) {
				case error:
					err = t
				case string:
					err = errors.New(t)
				default:
					err = errors.New("Vad gick fel här? Vi vet inte.")
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})

}
