package simple_util

import (
	"bufio"
	"os"
	"regexp"
)

type MapDb struct {
	Title []string
	Data  []map[string]string
}

// reads file to []string
func File2Array(fileName string) []string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2Array(scanner)
}

// reads file to [][]array
func File2Slice(fileName, sep string) [][]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2Slice(scanner, sep)
}

// reads file to []map[string]string
func File2MapArray(fileName, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2MapArray(scanner, sep, skip)
}

// reads file to map[string]map[string]string
func File2MapMap(fileName, key, sep string) map[string]map[string]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	scanner := bufio.NewScanner(file)
	return Scanner2MapMap(scanner, key, sep)
}
