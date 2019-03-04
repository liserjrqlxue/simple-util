package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/liserjrqlxue/simple-util"
	"os"
	"strings"
)

var (
	input = flag.String(
		"input",
		"",
		"input file",
	)
	sep = flag.String(
		"sep",
		"/",
		"sep string to split lines",
	)
	count = flag.Int(
		"count",
		1,
		"first count columns as output filenames",
	)
)

func main() {
	flag.Parse()
	if *input == "" {
		flag.Usage()
		fmt.Println("-input is reuqired!")
		os.Exit(1)
	}

	inFile, err := os.Open(*input)
	simple_util.CheckErr(err)
	defer simple_util.DeferClose(inFile)

	var outList = make(map[string]*os.File)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), *sep)
		outFile := strings.Join(items[0:*count], *sep)
		if outList[outFile] == nil {
			file, err := os.Create(outFile)
			simple_util.CheckErr(err)
			outList[outFile] = file
		}
		_, err := outList[outFile].WriteString(scanner.Text() + "\n")
		simple_util.CheckErr(err)
	}
	for k, f := range outList {
		f.Close()
		fmt.Printf("close %s\n", k)
	}
}
