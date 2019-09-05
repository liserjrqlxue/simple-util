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

func Symlink(source, dest string) error {
	_, err := os.Stat(dest)
	if err == nil {
		readLink, err := os.Readlink(dest)
		if err != nil {
			log.Printf("%v\n", err)
		}
		if readLink != source {
			log.Printf("dest is not symlink of source:[%s]->[%s]vs[%s]\n", dest, readLink, source)
			err = os.Symlink(source, dest)
			if err != nil {
				log.Printf("%v\n", err)
			}
		} else {
			log.Printf("dest is symlink of source:[%s]->[%s]", dest, readLink)
		}
	} else if os.IsNotExist(err) {
		err = os.Symlink(source, dest)
		if err != nil {
			log.Printf("Error: Symlink[%s->%s] err:%v", source, dest, err)
		}
	} else {
		log.Printf("Error: dest[%s] stat err:%v", dest, err)
	}
	return nil
}
