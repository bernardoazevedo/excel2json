package xls

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shakinm/xlsReader/xls"
)

func Json(file *os.File) (string, error) {
	workbook, err := xls.OpenReader(file)
	if err != nil {
		return "", fmt.Errorf("error opening %s: %v", file.Name(), err)
	}

	worksheet, err := workbook.GetSheet(0)
	if err != nil {
		return "", fmt.Errorf("error getting sheet: %v", err)
	}
	
	rows := worksheet.GetRows()
	tableMap := []map[string]string{}
	header := []string{}
	for rowIndex, eachRow := range rows {
		cols := eachRow.GetCols()

		if rowIndex == 0 {
			for _, eachCol := range cols {
				eachHeader := eachCol.GetString()
				if len(eachHeader) > 0 {
					header = append(header, eachCol.GetString())
				}
			}
		} else {
			tableRow := map[string]string{}
			for headerIndex, eachHeader := range header {
				if headerIndex < len(cols) {
					tableRow[eachHeader] = cols[headerIndex].GetString()
				} else {
					tableRow[eachHeader] = ""
				}
			}
			tableMap = append(tableMap, tableRow)
		}
	}

	jsonSheet, err := json.Marshal(tableMap)
	if err != nil {
		return "", fmt.Errorf("error parsing to json: %v", err)
	}

	return string(jsonSheet), nil
}
