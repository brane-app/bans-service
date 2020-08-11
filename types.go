package main

import (
	"github.com/gastrodon/groudon"

	"fmt"
)

const (
	DAY = 60 * 60 * 24
)

func ValidNumber(it interface{}) (ok bool, _ error) {
	_, ok = it.(float64)
	panic(fmt.Sprintf("%t", it))
	return
}

func validReason(it interface{}) (ok bool, _ error) {
	var reason string
	if reason, ok = it.(string); !ok {
		return
	}

	ok = len(reason) <= 255
	return
}

type CreateBanBody struct {
	Banned   string `json:"banned"`
	Reason   string `json:"reason"`
	Forever  bool   `json:"forever"`
	Duration int    `json:"duration"`
}

func (_ CreateBanBody) Validators() (values map[string]func(interface{}) (bool, error)) {
	values = map[string]func(interface{}) (bool, error){
		"banned":   groudon.ValidUUID,
		"reason":   validReason,
		"forever":  groudon.ValidBool,
		"duration": groudon.ValidNumber,
	}

	return
}

func (_ CreateBanBody) Defaults() (values map[string]interface{}) {
	values = map[string]interface{}{
		"reason":   "",
		"forever":  false,
		"duration": DAY,
	}

	return
}
