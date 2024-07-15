package middleware

import (
	"net/http"
	"errors"
	"github.com/Simon-Busch/go_simple_api/internal/tools"
	"github.com/Simon-Busch/go_simple_api/api"
	log "github.com/sirupsen/logrus"
)

var UnAUthorizedError = errors.New("Unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAUthorizedError)
			api.RequestErrorHandler(w, UnAUthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAUthorizedError)
			api.RequestErrorHandler(w, UnAUthorizedError)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
