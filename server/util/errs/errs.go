package errs

import (
	"fmt"
	"net/http"
)

// AError is a custom error to be used conveniently by endpoints
type AError struct {
	Type   string `json:"type"`
	Msg    string `json:"message"`
	Status int    `json:"-"`
}

func (e *AError) Error() string {
	return fmt.Sprintf("%s - %s", e.Type, e.Msg)
}

// New create a new AError
func New(typ, msg string, status int) *AError {
	return &AError{Type: typ, Msg: msg, Status: status}
}

func NewDB(msg string) *AError {
	return &AError{Type: "db", Msg: msg, Status: http.StatusInternalServerError}
}

var (
	// ErrDBUsernameExists is returned when trying to signup using existing username.
	// ErrDBUsernameExists    = &AError{Type: "db", Msg: "username already exists", Status: http.StatusNotFound}
	// ErrDBUsernameNotExists = &AError{Type: "db", Msg: "username doesn't exist", Status: http.StatusNotFound}
	// ErrDBEmailExists is returned when trying to signup using existing email.
	// ErrDBEmailExists    = &AError{Type: "db", Msg: "email already exists", Status: http.StatusBadRequest}
	// ErrDBEmailNotExists = &AError{Type: "db", Msg: "email doesn't exist", Status: http.StatusNotFound}

	// ErrDBIDNotExists is an error when db lookup for id didn't successfully found
	ErrDBIDNotExists = New("db", "id doesn't exist", http.StatusNotFound)

	// ErrDBUnknown is an unknown error happened when talking to database
	ErrDBUnknown = New("db", "internal error", http.StatusInternalServerError)

	// ErrAuthNoToken is an error happened when a request contains no token
	ErrAuthNoToken = New("auth", "no token", http.StatusUnauthorized)

	// ErrAuthInvalidToken is an error happened when a request contains an invalid token
	ErrAuthInvalidToken = New("auth", "invalid token", http.StatusUnauthorized)

	// ErrRequestBadParam is an error happened because request is not made properly
	ErrRequestBadParam = New("request", "bad request", http.StatusBadRequest)

	// ErrRequestInvalidCred is an error for when a login attempt failed because of bad credentials
	ErrRequestInvalidCred = New("request", "invalid credentials", http.StatusBadRequest)

	// ErrUnknown is an error when unexpected things happen
	ErrUnknown = New("server", "internal error", http.StatusInternalServerError)
)
