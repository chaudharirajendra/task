package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Common http statuses
const (
	BadRequest   = 400
	NotFound     = 404
	ServError    = 500
	Unauthorized = 401
	Forbidden    = 403
	Conflict     = 409
)

// send will marshal the data to json and return it to the user
func send(w http.ResponseWriter, r *http.Request, i interface{}) {
	if i == nil {
		return
	}
	b, err := json.Marshal(&i)
	if handleError(w, r, ServError, err, "Could not create json") {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(b))
}

// handleError will evaluate for an error and send appropriate response
// and log the issue
func handleError(w http.ResponseWriter, r *http.Request, status int, err error, msg ...string) bool {
	if err == nil {
		return false
	}
	defer r.Body.Close()
	errMsg := err.Error()
	if len(msg) > 0 {
		errMsg = msg[0]
	}
	http.Error(w, errMsg, status)
	return true
}
