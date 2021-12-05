package rkutil

import "log"

func Unexpected(err error) {
	log.Fatal("unexpected error: " + err.Error())
}
