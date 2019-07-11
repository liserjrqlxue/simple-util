package simple_util

import (
	"io"
	"log"
	"os"
)

// check if a file exists an is not a directory
func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CopyFile(dst, src string) (err error) {
	r, err := os.Open(src)
	if err != nil {
		return
	}
	defer DeferClose(r)

	w, err := os.Create(dst)
	if err != nil {
		return
	}
	defer DeferClose(w)

	n, err := io.Copy(w, r)
	if err != nil {
		return
	}
	log.Printf("CopyFile %d bytes[%s -> %s]", n, src, dst)
	return w.Sync()
}
