package handler

import (
	log "github.com/inconshreveable/log15"
	"gopkg.in/nullbio/null.v6"
	"panjiesw.com/mypict/server/fb"
)

var RootCtxKey = &contextKey{"RootContext"}

type RootCtx struct {
	log   log.Logger
	reqID string
	user  *fb.User
}

func (r *RootCtx) IsLoggedIn() bool {
	return r.user != nil
}

func (r *RootCtx) UID() null.String {
	if r.IsLoggedIn() {
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
