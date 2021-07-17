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

	prefix := os.Getenv("PATH_PREFIX")
	log.Printf("Routing with a prefix %s\n", prefix)

	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
