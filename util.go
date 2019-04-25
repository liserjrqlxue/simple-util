package simple_util

import (
	"log"
)

// handle error
func CheckErr(err error, msg ...string) {
	if err != nil {
		//panic(err)
		log.Fatal(err, msg)
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
