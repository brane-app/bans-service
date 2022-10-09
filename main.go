package main

import (
	"github.com/brane-app/librane/database"
	"github.com/gastrodon/groudon/v2"

	"log"
	"net/http"
	"os"
)

func health_check(_ *http.Request) (int, map[string]interface{}, error) {
	return 204, nil, database.Health()
}

func main() {
	database.Connect(os.Getenv("DATABASE_CONNECTION"))

	groudon.AddHandler("GET", "^/health/?$", health_check)
	register_handlers()

	prefix := os.Getenv("PATH_PREFIX")
	log.Printf("Routing with a prefix %s\n", prefix)

	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
