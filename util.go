package simple_util

import (
	"log"
)

// handle error
func CheckErr(err error) {
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
}

type handle interface {
	Close() error
}

// handle error while defer Close()
func DeferClose(h handle) {
	err := h.Close()
	CheckErr(err)
}
