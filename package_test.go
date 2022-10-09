package main

import (
	"github.com/brane-app/librane/database"

	"os"
	"testing"
)

func TestMain(main *testing.M) {
	database.Connect(os.Getenv("DATABASE_CONNECTION"))
	database.Create()

	setup(main)

	os.Exit(main.Run())
}
