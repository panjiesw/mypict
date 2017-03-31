package main

import (
	"panjiesw.com/mypict/server/db"
	"panjiesw.com/mypict/server/handler"
	"panjiesw.com/mypict/server/util/config"
	"panjiesw.com/mypict/server/util/logging"
)

func init() {
	config.BindEnv("MP")
}

func main() {
	c, err := config.Parse("")
	if err != nil {
		panic(err)
	}

	zc := logging.NewZapConfig(c.Env)
	z, err := zc.Build()
	if err != nil {
		panic(err)
	}

	zl := z.Sugar().Named("server")

	d, err := db.Open(c, zl.Named("db"))
	if err != nil {
		zl.Fatalf("Failed to open db connection: %v", err)
	}

	s := handler.New(c, d, zl.Named("handler"))
	s.Start()
}
