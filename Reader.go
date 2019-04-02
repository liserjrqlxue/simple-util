package simple_util

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

func Reader2MapArray(reader *bufio.Reader, sep string, skip *regexp.Regexp) (mapArray []map[string]string, title []string) {
	var err error
	var i = 0
	for {
		line, err := reader.ReadString('\n')
		if skip != nil && skip.MatchString(line) {
			continue
		}
		line = strings.TrimSuffix(line, "\n")
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
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		CheckErr(err)
	}
	return
}
