package simple_util

import (
	"bufio"
	"log"
	"os"
	"regexp"
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
