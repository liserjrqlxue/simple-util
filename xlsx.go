package simple_util

import "github.com/360EntSecGroup-Skylar/excelize"

func Sheet2MapArray(excelFile, sheetName string) ([]string, []map[string]string) {
	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows := xlsxFh.GetRows(sheetName)

	return ArrayArray2MapArray(rows)
}

func Sheet2MapMap(excelFile, sheetName, key string) ([]string, map[string]map[string]string) {
	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows := xlsxFh.GetRows(sheetName)

	return ArrayArray2MapMap(rows, key)
}

func Sheet2MapMapMerge(excelFile, sheetName, key, sep string) ([]string, map[string]map[string]string) {

	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows := xlsxFh.GetRows(sheetName)

	return ArrayArray2MapMapMerge(rows, key, sep)
}
