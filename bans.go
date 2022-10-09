package main

import (
	"github.com/brane-app/librane/database"
	"github.com/brane-app/librane/tools"
	"github.com/brane-app/librane/types"
	"github.com/gastrodon/groudon/v2"

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
	var ban map[string]interface{} = types.NewBan(
		requester,
		body.Banned,
		body.Reason,
		int64(body.Duration),
		body.Forever,
	).Map()

	err = database.WriteBan(ban)
	code = 200
	r_map = map[string]interface{}{"ban": ban}
	return
}

func readBan(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var parts []string = strings.FieldsFunc(request.URL.Path, pathSplit)

	var ban types.Ban
	var exists bool
	if ban, exists, err = database.ReadSingleBan(parts[len(parts)-1]); err != nil {
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

func readBansOfUser(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var parts []string = tools.SplitPath(request.URL.Path)
	var query map[string]interface{} = request.Context().Value("query").(map[string]interface{})

	var ID string = parts[len(parts)-1]
	var before string = query["before"].(string)
	var size int = query["size"].(int)

	var bans []types.Ban
	if bans, size, err = database.ReadBansOfUser(ID, before, size); err != nil {
		return
	}

	code = 200
	r_map = map[string]interface{}{
		"bans": bans,
		"size": map[string]int{"bans": size},
	}

	return
}
