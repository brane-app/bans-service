package main

import (
	"github.com/brane-app/tools-library/middleware"
	"github.com/gastrodon/groudon/v2"

	"os"
)

var (
	prefix = os.Getenv("PATH_PREFIX")

	routeRoot           = "^" + prefix + "/?$"
	readBanRoute        = "^" + prefix + "/id/" + groudon.UUID_PATTERN + "/?$"
	readBansOfUserRoute = "^" + prefix + "/user/id/" + groudon.UUID_PATTERN + "/?$"

	forbidden = map[string]interface{}{"error": "forbidden"}
)

func register_handlers() {
	groudon.AddCodeResponse(403, forbidden)

	groudon.AddMiddleware("GET", ".*", middleware.MustAuth)
	groudon.AddMiddleware("POST", ".*", middleware.MustAuth)

	groudon.AddMiddleware("GET", ".*", middleware.MustModerator)
	groudon.AddMiddleware("POST", ".*", middleware.MustModerator)

	groudon.AddMiddleware("GET", readBansOfUserRoute, middleware.PaginationParams)

	groudon.AddHandler("POST", routeRoot, createBan)
	groudon.AddHandler("GET", readBanRoute, readBan)
	groudon.AddHandler("GET", readBansOfUserRoute, readBansOfUser)
}
