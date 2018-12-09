package simple_util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// handle error
func CheckErr(err error) {
	if err != nil {
		//panic(err)
		log.Fatal(err)
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

// reads file to []string
func File2Array(path string) []string {
	file, err := os.Open(path)
	CheckErr(err)
	defer DeferClose(file)
	var array []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	CheckErr(scanner.Err())
	return array
}

// reads file to [][]array
func File2Slice(path, sep string) [][]string {
	file, err := os.Open(path)
	CheckErr(err)
	defer DeferClose(file)

	var slice [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, sep)
		slice = append(slice, array)
	}
	CheckErr(scanner.Err())
	return slice
}

// reads file to []map[string]string
func File2MapArray(path, sep string) []map[string]string {
	file, err := os.Open(path)
	CheckErr(err)
	defer DeferClose(file)

	var mapArray []map[string]string
	var title []string
	scanner := bufio.NewScanner(file)
	var i = 0
	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, sep)
		if i == 0 {
			title = array
		} else {
			var dataHash = make(map[string]string)
			for j, k := range array {
				dataHash[title[j]] = k
			}
			mapArray = append(mapArray, dataHash)
		}
		i++
	}
	CheckErr(scanner.Err())
	return mapArray
}
