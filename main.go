package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib/middleware"
	"github.com/gastrodon/groudon"

	"log"
	"net/http"
	"os"
)

const (
	readBanRoute        = "^/id/" + groudon.UUID_PATTERN + "/?$"
	readBansOfUserRoute = "^/user/id/" + groudon.UUID_PATTERN + "/?$"
)

var (
	forbidden = map[string]interface{}{"error": "forbidden"}
)

func main() {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))

	groudon.RegisterCatch(403, forbidden)
	groudon.RegisterMiddleware(middleware.MustAuth)
	groudon.RegisterMiddleware(middleware.MustModerator)
	groudon.RegisterMiddlewareRoute([]string{"GET"}, readBansOfUserRoute, middleware.PaginationParams)

	groudon.RegisterHandler("POST", `^/$`, createBan)
	groudon.RegisterHandler("GET", readBanRoute, readBan)
	groudon.RegisterHandler("GET", readBansOfUserRoute, readBansOfUser)

	http.Handle("/", http.HandlerFunc(groudon.Route))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
