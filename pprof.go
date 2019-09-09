package simple_util

import (
	"log"
	"os"
	"runtime/pprof"
)

func MemProfile(memProfile string) {
	f, err := os.Create(memProfile)
	if err != nil {
		log.Fatal(err)
	}
	CheckErr(pprof.WriteHeapProfile(f))
	defer DeferClose(f)
}
