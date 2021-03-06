package db_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"panjiesw.com/mypict/server/db"
	"panjiesw.com/mypict/server/util/config"
	"panjiesw.com/mypict/server/util/dbtest"
)

var (
	d *db.Database

	cwd = flag.String("cwd", "", "set cwd")
)

func setupDB() []error {
	var wd string
	if *cwd != "" {
		wd = *cwd
	} else {
		wd = "./.."
	}

	if errs := dbtest.Migrate(wd); errs != nil {
		for _, err := range errs {
			fmt.Printf("Failed to migrate: %s\n", err.Error())
		}
		return errs
	}

	if err := dbtest.Fixtures(wd); err != nil {
		fmt.Printf("Failed to seed: %s\n", err.Error())
		return []error{err}
	}

	var err error

	c, err := config.Parse("")
	if err != nil {
		fmt.Printf("Failed to init config: %s", err.Error())
		return []error{err}
	}

	d, err = db.Open(c)
	if err != nil {
		fmt.Printf("Failed to create db: %s", err.Error())
		return []error{err}
	}

	return nil
}

func cleanupDB() {
	defer dbtest.Cleanup()
	d.Close()
}

func TestMain(m *testing.M) {
	if err := setupDB(); err != nil {
		if d == nil {
			dbtest.Cleanup()
		} else {
			cleanupDB()
		}
		os.Exit(-1)
		return
	}
	result := m.Run()
	cleanupDB()
	os.Exit(result)
}

func init() {
	flag.Parse()
}
