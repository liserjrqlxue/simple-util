package simple_util

import "log"

func CheckErr(err error) {
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
}

type handle interface {
	Close() error
}

func DeferClose(h handle) {
	err := h.Close()
	CheckErr(err)
}
