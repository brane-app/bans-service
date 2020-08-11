package main

import (
	"github.com/gastrodon/groudon"
	"github.com/imonke/monkebase"
	"github.com/imonke/monketype"

	"net/http"
	"strings"
)

func pathSplit(it rune) (ok bool) {
	ok = it == '/'
	return
}

func createBan(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var body CreateBanBody
	var external error
	if err, external = groudon.SerializeBody(request.Body, &body); err != nil || external != nil {
		code = 400
		return
	}

	var requester string = request.Context().Value("requester").(string)
	var ban map[string]interface{} = monketype.NewBan(
		requester,
		body.Banned,
		body.Reason,
		int64(body.Duration),
		body.Forever,
	).Map()

	err = monkebase.WriteBan(ban)
	code = 200
	r_map = map[string]interface{}{"ban": ban}
	return
}

func readBan(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var parts []string = strings.FieldsFunc(request.URL.Path, pathSplit)

	var ban monketype.Ban
	var exists bool
	if ban, exists, err = monkebase.ReadSingleBan(parts[len(parts)-1]); err != nil {
		return
	}

	if !exists {
		code = 404
		r_map = map[string]interface{}{"error": "no_such_ban"}
		return
	}

	code = 200
	r_map = map[string]interface{}{"ban": ban.Map()}
	return
}
