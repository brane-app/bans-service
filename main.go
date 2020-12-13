package main

import (
	"github.com/gastrodon/groudon"
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib/middleware"

	"log"
	"net/http"
	"os"
)

var (
	forbidden = map[string]interface{}{"error": "forbidden"}
)

func main() {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))

	groudon.RegisterCatch(403, forbidden)
	groudon.RegisterMiddleware(middleware.MustAuth)
	groudon.RegisterMiddleware(middleware.MustModerator)

	groudon.RegisterHandler("POST", `^/$`, createBan)
	groudon.RegisterHandler("GET", `^/id/`+groudon.UUID_PATTERN+`/?$`, readBan)

	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
