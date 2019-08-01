package simple_util

import (
	"log"
	"strconv"
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

func CheckAFAllLowThen(item map[string]string, AFList []string, threshold float64, includeEqual bool) bool {
	for _, key := range AFList {
		af := item[key]
		if af == "" || af == "." || af == "0" {
			continue
		}
		AF, err := strconv.ParseFloat(af, 64)
		CheckErr(err)
		if includeEqual {
			if AF > threshold {
				return false
			}
		} else {
			if AF >= threshold {
				return false
			}
		}
	}
	return true
}
