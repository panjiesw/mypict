package handler

import (
	"net/http"

	"github.com/mgutz/logxi"
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/util/fb"
)

var RootCtxKey = &contextKey{"RootContext"}

func GetContext(r *http.Request) *RootCtx {
	return r.Context().Value(RootCtxKey).(*RootCtx)
}

type RootCtx struct {
	log   logxi.Logger
	reqID string
	user  *fb.User
}

func (r *RootCtx) IsLoggedIn() bool {
	return r.user != nil
}

func (r *RootCtx) UID() null.String {
	if !r.IsLoggedIn() {
		return null.NewString("", false)
	}
	if r.user.ID == "" {
		return null.NewString("", false)
	}
	return null.StringFrom(r.user.ID)
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "mypict context value " + k.name
}
