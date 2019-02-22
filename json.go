package simple_util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func Json2MapMap(jsonBlob []byte) map[string]map[string]string {
	var data = make(map[string]map[string]string)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func JsonFile2MapMap(fileName string) map[string]map[string]string {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	fmt.Printf("load %d byte from %s\n", len(b), fileName)
	return Json2MapMap(b)
}
