package simple_util

func Slice2MapArray(slice [][]string) ([]string, []map[string]string) {
	var data []map[string]string
	var keys []string

	for i, row := range slice {
		if i == 0 {
			keys = row
		} else {
			var dataHash = make(map[string]string)
			for j, cell := range row {
				dataHash[keys[j]] = cell
			}
			data = append(data, dataHash)
		}
	}
	return keys, data
}

func Slice2MapMap(slice [][]string, key string) ([]string, map[string]map[string]string) {
	var data = make(map[string]map[string]string)
	var keys []string
	for i, row := range slice {
		if i == 0 {
			keys = row
		} else {
			var dataHash = make(map[string]string)
			for j, cell := range row {
				dataHash[keys[j]] = cell
			}
			data[dataHash[key]] = dataHash
		}
	}
	return keys, data
}

func Slice2MapMapMerge(slice [][]string, key, sep string) ([]string, map[string]map[string]string) {
	var data = make(map[string]map[string]string)
	var keys []string

	for i, row := range slice {
		if i == 0 {
			keys = row
		} else {
			var dataHash = make(map[string]string)
			for j, cell := range row {
				dataHash[keys[j]] = cell
			}
			mainKey := dataHash[key]
			if data[mainKey] == nil {
				data[mainKey] = dataHash
			} else {
				for _, subKey := range keys {
					data[mainKey][subKey] = data[mainKey][subKey] + sep + dataHash[subKey]
				}
			}
		}
	}

	return keys, data
}
