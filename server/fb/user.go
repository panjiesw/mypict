package fb

import "github.com/wuman/firebase-server-sdk-go"

type User struct {
	ID       string
	Name     string
	Email    string
	Pict     string
	Verified bool
	Issuer   string
}

func UserFromToken(t *firebase.Token) *User {
	id, ok := t.UID()
	if !ok {
		logger.Warn("UID not found")
	}
	iss, ok := t.Issuer()
	if !ok {
		logger.Warn("Issuer not found")
	}
	name, _ := t.Name()
	email, _ := t.Email()
	pict, _ := t.Picture()
	verified, _ := t.IsEmailVerified()

	return &User{id, name, email, pict, verified, iss}
}
