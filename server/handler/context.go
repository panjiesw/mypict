package handler

import (
	log "github.com/inconshreveable/log15"
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

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "mypict context value " + k.name
}
