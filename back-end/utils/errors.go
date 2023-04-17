package utils

import "log"

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogIfError(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}
