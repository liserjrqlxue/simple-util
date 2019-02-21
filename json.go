package simple_util

import (
	"fmt"
	"os"
)

func Json2file(json []byte, filenName string) error {
	file, err := os.Create(filenName)
	if err != nil {
		return err
	}
	defer DeferClose(file)

	c, err := file.Write(json)
	if err != nil {
		return err
	}
	fmt.Printf("write %d byte to %s\n", c, filenName)

	return nil
}
