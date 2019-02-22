package simple_util

import (
	"bufio"
	"compress/gzip"
	"os"
	"regexp"
)

func Gz2Array(fileName string) []string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	gr, err := gzip.NewReader(file)
	CheckErr(err)
	defer DeferClose(gr)

	scanner := bufio.NewScanner(gr)
	return Scanner2Array(scanner)
}

func Gz2Slice(fileName, sep string) [][]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	gr, err := gzip.NewReader(file)
	CheckErr(err)
	defer DeferClose(gr)

	scanner := bufio.NewScanner(gr)
	return Scanner2Slice(scanner, sep)
}

func Gz2MapArray(fileName, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	gr, err := gzip.NewReader(file)
	CheckErr(err)
	defer DeferClose(gr)

	scanner := bufio.NewScanner(gr)
	return Scanner2MapArray(scanner, sep, skip)
}

func Gz2MapMap(fileName, key, sep string) map[string]map[string]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer DeferClose(file)

	gr, err := gzip.NewReader(file)
	CheckErr(err)
	defer DeferClose(gr)

	scanner := bufio.NewScanner(gr)
	return Scanner2MapMap(scanner, key, sep)
}
