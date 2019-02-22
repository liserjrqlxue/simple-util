package simple_util

import (
	"bufio"
	"strings"
)

func Scanner2Array(scanner *bufio.Scanner) []string {
	var array []string
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	CheckErr(scanner.Err())
	return array
}

func Scanner2Slice(scanner *bufio.Scanner, sep string) [][]string {
	var slice [][]string
	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, sep)
		slice = append(slice, array)
	}
	CheckErr(scanner.Err())
	return slice
}

func Scanner2MapArray(scanner *bufio.Scanner, sep string) ([]map[string]string, []string) {
	var mapArray []map[string]string
	var title []string
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
	return mapArray, title
}

func Scanner2MapMap(scanner *bufio.Scanner, key, sep string) map[string]map[string]string {
	var db = make(map[string]map[string]string)
	var title []string
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
			db[dataHash[key]] = dataHash
		}
		i++
	}
	CheckErr(scanner.Err())
	return db
}
