package rkutil

import (
	"errors"
	"log"
)

func Unexpected(err error) {
	log.Fatal("unexpected error: " + err.Error())
}

func UnexpectedCodePath() {
	Unexpected(errors.New("hit unexpected code path"))
}

func Ensure(v bool, msg string) {
	if !v {
		Unexpected(errors.New(msg))
	}
}
