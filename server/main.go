package main

import (
	"github.com/mgutz/logxi"
	"panjiesw.com/mypict/server/db"
	"panjiesw.com/mypict/server/handler"
	"panjiesw.com/mypict/server/util/config"
)

var mlog = logxi.New("main")

func init() {
	config.BindEnv("MP")
}

func main() {
	c, err := config.Parse("")
	if err != nil {
		mlog.Fatal("Failed to parse config", "err", err)
	}

	d, err := db.Open(c)
	if err != nil {
		mlog.Fatal("Failed to open db connection", "err", err)
	}

	s := handler.New(c, d)
	s.Start()
}
