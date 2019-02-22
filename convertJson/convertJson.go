package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/liserjrqlxue/simple-util"
	"os"
)

var (
	input = flag.String(
		"input",
		"",
		"input file",
	)
	format = flag.String(
		"type",
		"tsv",
		"file type,[tsv,xlsx]",
	)
	sep = flag.String(
		"sep",
		"\t",
		"split column",
	)
	merge = flag.String(
		"merge",
		"\n",
		"merge element")
	key = flag.String(
		"key",
		"",
		"column key as main key",
	)
	strucType = flag.String(
		"struct",
		"MapArray",
		"output data struct,[MapArray|MapMap|Slice|Array]",
	)
	prefix = flag.String(
		"prefix",
		"",
		"prefix of output, suffix with .json",
	)
	sheet = flag.String(
		"sheet",
		"",
		"sheet name for xlsx",
	)
)

func main() {
	flag.Parse()
	if *input == "" {
		flag.Usage()
		fmt.Println("-input is required")
		os.Exit(1)
	}
	if *format == "xlsx" && *sheet == "" {
		flag.Usage()
		fmt.Println("-sheet is required for xlsx")
		os.Exit(1)
	}
	if *sep == "" && *strucType != "Array" {
		flag.Usage()
		fmt.Println("-sep is required for struct not Array")
		os.Exit(1)
	}
	if *strucType == "MapMap" && *key == "" {
		flag.Usage()
		fmt.Println(" need -key is required as main key for MapMap struct")
		os.Exit(1)
	}
	if *prefix == "" {
		*prefix = *input
	}

	var i interface{}
	var d []byte
	var err error
	switch *format {
	case "tsv":
		switch *strucType {
		case "Array":
			i = simple_util.File2Array(*input)
		case "Slice":
			i = simple_util.File2Slice(*input, *sep)
		case "MapArray":
			i, _ = simple_util.File2MapArray(*input, *sep, nil)
		case "MapMap":
			i = simple_util.File2MapMapMerge(*input, *key, *sep, *merge)
		default:
			flag.Usage()
			fmt.Println("-struct error for tsv")
			os.Exit(1)
		}
		d, err = json.Marshal(i)
		simple_util.CheckErr(err)
		simple_util.Json2file(d, *prefix+".json")
	case "xlsx":
		switch *strucType {
		case "MapArray":
			i, _ = simple_util.Sheet2MapArray(*input, *sheet)
		case "MapMap":
			_, i = simple_util.Sheet2MapMapMerge(*input, *sheet, *key, *merge)
		default:
			flag.Usage()
			fmt.Println("-struct error for xlsx")
			os.Exit(1)
		}
		d, err = json.Marshal(i)
		simple_util.CheckErr(err)
		simple_util.Json2file(d, *prefix+"."+*sheet+".json")
	}

}
