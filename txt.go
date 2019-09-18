package simple_util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

type MapDb struct {
	Title []string
	Data  []map[string]string
}

// read file to []string
func File2Array(fileName string) []string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2Array(scanner)
}

// read file to [][]array
func File2Slice(fileName, sep string) [][]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2Slice(scanner, sep)
}

// read file to []map[string]string
func File2MapArray(fileName, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2MapArray(scanner, sep, skip)
}

// read file with long line to []map[string]string
func LongFile2MapArray(fileName, sep string, skip *regexp.Regexp) (mapArray []map[string]string, title []string) {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	reader := bufio.NewReader(file)
	return Reader2MapArray(reader, sep, skip)
}

// read files to []map[string]string
func Files2MapArray(fileNames []string, sep string, skip *regexp.Regexp) (Data []map[string]string, Title []string) {
	for _, fileName := range fileNames {
		data, title := File2MapArray(fileName, sep, skip)
		for _, item := range data {
			Data = append(Data, item)
		}
		if len(Title) == 0 {
			Title = title
		} else {
			if len(Title) != len(title) {
				log.Fatal("titles has different columns")
			} else {
				for i := 0; i < len(Title); i++ {
					if Title[i] != title[i] {
						log.Fatal("titles not equal")
					}
				}
			}
		}
	}
	return
}

// read files to []map[string]string
func LongFiles2MapArray(fileNames []string, sep string, skip *regexp.Regexp) (Data []map[string]string, Title []string) {
	for _, fileName := range fileNames {
		data, title := LongFile2MapArray(fileName, sep, skip)
		for _, item := range data {
			Data = append(Data, item)
		}
		if len(Title) == 0 {
			Title = title
		} else {
			if len(Title) != len(title) {
				log.Fatal("titles has different columns")
			} else {
				for i := 0; i < len(Title); i++ {
					if Title[i] != title[i] {
						log.Fatal("titles not equal")
					}
				}
			}
		}
	}
	return
}

// read file to map[string]map[string]string
func File2MapMap(fileName, key, sep string) map[string]map[string]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2MapMap(scanner, key, sep)
}

// read file to map[string]map[string]string
func File2MapMapMerge(fileName, key, sep, merge string) map[string]map[string]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	_, d := Slice2MapMapMerge(Scanner2Slice(scanner, sep), key, merge)
	return d
}

// read two column file to map[string]string
func File2Map(fileName, sep string, override bool) (db map[string]string, err error) {
	db = make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		return db, err
	}
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	slice := Scanner2Slice(scanner, sep)
	for _, kv := range slice {
		kv = append(kv, "NA", "NA")
		v, ok := db[kv[0]]
		if ok && v != kv[1] && !override {
			err = errors.New("dup key[" + kv[0] + "],different value:[" + v + "," + kv[1] + "]")
		}
		db[kv[0]] = kv[1]
	}
	return
}

// read two column file to map[string]string
func Files2Map(fileNames, sep string, override bool) (db map[string]string, err error) {
	db = make(map[string]string)
	fileList := strings.Split(fileNames, ",")
	for _, fileName := range fileList {
		db1, err1 := File2Map(fileName, sep, override)
		for k, v := range db1 {
			db[k] = v
		}
		if err1 != nil {
			err = err1
		}
	}
	return
}
