package common

import (
	"log"
)

// Must not error
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Silent(err error) {
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}
}
