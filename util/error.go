package util

import "log"

func FailOnErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
