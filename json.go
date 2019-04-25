package simple_util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// warpper of json.MarshalIndent
func JsonIndent(v interface{}, prefix, indent string) (b []byte, err error) {
	b, err = json.MarshalIndent(v, prefix, indent)
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	return
}

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

func Json2File(fileName string, a interface{}) error {
	b, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}
	return Json2file(b, fileName)
}

func Json2MapMap(jsonBlob []byte) map[string]map[string]string {
	var data = make(map[string]map[string]string)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func Json2Map(jsonBlob []byte) map[string]string {
	var data = make(map[string]string)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func Json2MapInt(jsonBlob []byte) map[string]int {
	var data = make(map[string]int)
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

func JsonFile2Map(fileName string) map[string]string {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	fmt.Printf("load %d byte from %s\n", len(b), fileName)
	return Json2Map(b)
}

func JsonFile2MapInt(fileName string) map[string]int {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	fmt.Printf("load %d byte from %s\n", len(b), fileName)
	return Json2MapInt(b)
}

func JsonFile2Interface(fileName string) interface{} {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	fmt.Printf("load %d byte from %s\n", len(b), fileName)
	var data interface{}
	err = json.Unmarshal(b, &data)
	CheckErr(err)
	return data
}
