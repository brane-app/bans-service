package main

import (
	"github.com/gastrodon/groudon"
	"github.com/imonke/monkebase"
	"github.com/imonke/monkelib/middleware"

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
	groudon.RegisterMiddleware(middleware.RangeQueryParams)
	groudon.RegisterMiddleware(MustModerator)

	groudon.RegisterHandler("POST", `^/$`, createBan)
	groudon.RegisterHandler("GET", `^/id/`+groudon.UUID_PATTERN+`/?$`, readBan)

	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
