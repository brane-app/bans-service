package main

import (
	"github.com/brane-app/database-library"
	"github.com/gastrodon/groudon/v2"

	"log"
	"net/http"
	"os"
)

func main() {
	database.Connect(os.Getenv("DATABASE_CONNECTION"))

	register_handlers()

	http.Handle(os.Getenv("PATH_PREFIX")+"/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
