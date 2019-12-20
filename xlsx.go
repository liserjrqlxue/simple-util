package simple_util

import "github.com/360EntSecGroup-Skylar/excelize/v2"

func Sheet2MapArray(excelFile, sheetName string) ([]string, []map[string]string) {
	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows, err := xlsxFh.GetRows(sheetName)
	CheckErr(err)

	return Slice2MapArray(rows)
}

func Sheet2MapMap(excelFile, sheetName, key string) ([]string, map[string]map[string]string) {
	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows, err := xlsxFh.GetRows(sheetName)
	CheckErr(err)

	return Slice2MapMap(rows, key)
}

func Sheet2MapMapMerge(excelFile, sheetName, key, sep string) ([]string, map[string]map[string]string) {

	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows, err := xlsxFh.GetRows(sheetName)
	CheckErr(err)

	return Slice2MapMapMerge(rows, key, sep)
}

func Sheet2MapMapMergeTrim(excelFile, sheetName, key, sep string) ([]string, map[string]map[string]string) {

	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows, err := xlsxFh.GetRows(sheetName)
	CheckErr(err)

	return Slice2MapMapMergeTrim(rows, key, sep)
}

func Sheet2MapMapMergeReplace(excelFile, sheetName, key, sep, replace string) ([]string, map[string]map[string]string) {

	xlsxFh, err := excelize.OpenFile(excelFile)
	CheckErr(err)
	rows, err := xlsxFh.GetRows(sheetName)
	CheckErr(err)

	return Slice2MapMapMergeTrim(rows, key, sep)
}
