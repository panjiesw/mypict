package fb

import (
	"errors"

	"net/http"

	"github.com/mgutz/logxi"
	"github.com/spf13/viper"
	"github.com/wuman/firebase-server-sdk-go"
	"panjiesw.com/mypict/server/util/errs"
)

var logger = logxi.New("fb")

func Initialize() (*firebase.App, error) {
	p := viper.GetString("firebase.service_account_path")
	if p == "" {
		return nil, errors.New("Firebase service account path is not defined")
	}
	return firebase.InitializeApp(&firebase.Options{
		ServiceAccountPath: "",
	})
}

func VerifyTokenFromReq(r *http.Request, required bool) (*User, *errs.AError) {
	fts := r.Header.Get("authorization")
	if fts == "" {
		if required {
			return nil, errs.ErrAuthNoToken
		} else {
			return nil, nil
		}
	}
	ts := fts[7:]

	if ts == "" {
		return nil, errs.ErrAuthInvalidToken
	}

	auth, err := firebase.GetAuth()
	if err != nil {
		logger.Error("Failed to get firebase auth instance", "err", err)
		return nil, errs.ErrUnknown
	}

	token, err := auth.VerifyIDToken(ts)
	if err != nil {
		logger.Warn("Failed to verify token", "err", err, "token", ts)
		return nil, errs.ErrAuthInvalidToken
	}
	return UserFromToken(token), nil
}
