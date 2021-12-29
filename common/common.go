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

func WithResult(err error) bool {
	return err != nil
}

func Combine(a uint, in []uint) bool {
	for _, i := range in {
		if i == a {
			return true
		}
	}

	return false
}
