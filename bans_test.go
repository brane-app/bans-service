package main

import (
	"github.com/google/uuid"
	"github.com/imonke/monkebase"
	"github.com/imonke/monketype"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	nick  = "bruce"
	email = "banner@imonke.io"
)

var (
	banner monketype.User
)

func mustMarshal(it interface{}) (data []byte) {
	var err error
	if data, err = json.Marshal(it); err != nil {
		panic(err)
	}

	return
}

func banOK(banner monketype.User, ban monketype.Ban) (err error) {
	if banner.ID != ban.Banner {
		err = fmt.Errorf("ID mismatch! have: %s, want: %s", banner.ID, ban.Banner)
		return
	}

	if time.Now().Unix() >= ban.Expires {
		err = fmt.Errorf("Ban expires in the past: %d", ban.Expires)
		return
	}

	return
}

func TestMain(main *testing.M) {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))
	banner = monketype.NewUser(nick, "", email)

	var err error
	if err = monkebase.WriteUser(banner.Map()); err != nil {
		panic(err)
	}

	var result int = main.Run()
	monkebase.DeleteUser(banner.ID)
	os.Exit(result)
}

func Test_createBan(test *testing.T) {
	var set []byte
	var sets [][]byte = [][]byte{
		mustMarshal(map[string]interface{}{
			"banned":  uuid.New().String(),
			"reason":  "They enjoy pokemon sword and / or shield",
			"forever": true,
		}),
		mustMarshal(map[string]interface{}{
			"banned": uuid.New().String(),
			"reason": "They enjoy pokemon sword and / or shield",
		}),
		mustMarshal(map[string]interface{}{
			"banned": uuid.New().String(),
		}),
	}

	var code int
	var r_map map[string]interface{}
	var err error

	var ban monketype.Ban
	var request *http.Request
	var valued context.Context

	for _, set = range sets {
		valued = context.WithValue(context.TODO(), "requester", banner.ID)
		if request, err = http.NewRequestWithContext(valued, "POST", "/", bytes.NewReader(set)); err != nil {
			test.Fatal(err)
		}

		if code, r_map, err = createBan(request); err != nil {
			test.Fatal(err)
		}

		if code != 200 {
			test.Errorf("got code %d", code)
		}

		ban = monketype.Ban{}
		ban.FromMap(r_map["ban"].(map[string]interface{}))
		if err = banOK(banner, ban); err != nil {
			test.Fatal(err)
		}
	}
}

func Test_createBan_wasBanned(test *testing.T) {
	var banned string = uuid.New().String()
	var set []byte = mustMarshal(map[string]interface{}{
		"banned": banned,
	})

	var request *http.Request
	var err error
	if request, err = http.NewRequestWithContext(
		context.WithValue(context.TODO(), "requester", banner.ID),
		"POST", "/", bytes.NewReader(set),
	); err != nil {
		test.Fatal(err)
	}

	var code int
	if code, _, err = createBan(request); err != nil {
		test.Fatal(err)
	}

	if code != 200 {
		test.Errorf("got code %d", code)
	}

	var isBanned bool
	if isBanned, err = monkebase.IsBanned(banned); err != nil {
		test.Fatal(err)
	}

	if !isBanned {
		test.Errorf("%s was not banned!", banned)
	}
}
