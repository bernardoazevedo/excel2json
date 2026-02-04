package excel2json

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bernardoazevedo/excel2json/csv"
	"github.com/bernardoazevedo/excel2json/ods"
	"github.com/bernardoazevedo/excel2json/xls"
	"github.com/bernardoazevedo/excel2json/xlsx"
)

type ExcelFile struct {
	File *os.File
}

func (excelFile ExcelFile) ToJson() (string, error) {
	var err error
	var json string

	switch excelFile.extension() {
	case ".xls":
		json, err = xls.Json(excelFile.File)
	case ".xlsx":
		json, err = xlsx.Json(excelFile.File)
	case ".csv":
		json, err = csv.Json(excelFile.File)
	case ".ods":
		json, err = ods.Json(excelFile.File)
	default:
		return "", fmt.Errorf("the '%s' file extension is not accepted", excelFile.extension())
	}
	if err != nil {
		return "", fmt.Errorf("error parsing file %v: ", err)
	}

	return json, nil
}

func (excelFile ExcelFile) extension() string {
	return filepath.Ext(excelFile.File.Name())
}
