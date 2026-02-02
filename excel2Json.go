package excel2json

import (
	"fmt"
	"os"
	"strings"

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
	case "xls":
		json, err = xls.Json(excelFile.File)
	case "xlsx":
		json, err = xlsx.Json(excelFile.File)
	default:
		return "", fmt.Errorf("file extension don't recognized: %s", excelFile.extension())
	}
	if err != nil {
		return "", fmt.Errorf("error parsing file %v: ", err)
	}

	return json, nil
}

func (excelFile ExcelFile) extension() string {
	filePaths := strings.Split(excelFile.File.Name(), "/")
	filename := filePaths[len(filePaths)-1]
	nameParts := strings.Split(filename, ".")
	extension := nameParts[len(nameParts)-1]
	return extension
}
