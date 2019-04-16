package main

import (
	"flag"
	"github.com/liserjrqlxue/simple-util"
	"os"
)

var (
	inputJson = flag.String(
		"input",
		"",
		"input json",
	)
)

func main() {
	flag.Parse()
	if *inputJson == "" {
		flag.Usage()
		os.Exit(0)
	}

	var jsonData = simple_util.JsonFile2Interface(*inputJson)
	jsonByte, err := simple_util.JsonIndent(jsonData, "", "\t")
	simple_util.CheckErr(err)
	simple_util.Json2file(jsonByte, *inputJson+".new.json")
}
