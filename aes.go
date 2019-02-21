package simple_util

import (
	"fmt"
	"github.com/liserjrqlxue/crypto/aes"
	"os"
)

func Encode2file(fileName string, data, codeKey []byte) {
	file, err := os.Create(fileName)
	CheckErr(err)
	defer DeferClose(file)
	c := Encode2File(file, data, codeKey)
	fmt.Printf("write %d byte to %s\n", c, fileName)
}

func Encode2File(file *os.File, data, codeKey []byte) int {
	d, err := AES.Encode(data, codeKey)
	CheckErr(err)
	c, err := file.Write(d)
	CheckErr(err)
	return c
}
