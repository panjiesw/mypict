package db_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"panjiesw.com/mypict/server/db"
	db2 "panjiesw.com/mypict/server/testutil/db"
)

var d *db.Database

var cwd = flag.String("cwd", "", "set cwd")

func setupDB() []error {
	var wd string
	if *cwd != "" {
		wd = *cwd
	} else {
		wd = "./.."
	}

	if errs := db2.Migrate(wd); errs != nil {
		for _, err := range errs {
			fmt.Printf("Failed to migrate: %s\n", err.Error())
		}
		return errs
	}

	if err := db2.Fixtures(wd); err != nil {
		fmt.Printf("Failed to seed: %s\n", err.Error())
		return []error{err}
	}

	var err error

	d, err = db.New()
	if err != nil {
		fmt.Printf("Failed to create db: %s", err.Error())
		return []error{err}
	}

	return nil
}

func cleanupDB() {
	defer db2.Cleanup()
	d.Close()
}

func TestMain(m *testing.M) {
	if err := setupDB(); err != nil {
		if d == nil {
			db2.Cleanup()
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
