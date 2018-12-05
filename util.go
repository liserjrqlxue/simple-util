package simple_util

import "log"

func checkErr(err error) {
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
}

type handle interface {
	Close() error
}

func deferClose(h handle) {
	err := h.Close()
	checkErr(err)
}
