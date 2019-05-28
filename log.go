package simple_util

import "log"

func logLoadJson(size int, fileName string) {
	log.Printf("load %10d byte from %s\n", size, fileName)
}
