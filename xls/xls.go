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

	headerFound := false
	var tableMap []map[string]string
	var header []string

	for _, eachRow := range rows {
		cols := eachRow.GetCols()

		if !headerFound {
			// header
			var tempHeader []string
			headerFilled := false

			// searching for the first filled row
			for _, eachCol := range cols {
				eachHeader := eachCol.GetString()
				if len(eachHeader) > 0 {
					headerFilled = true
					tempHeader = append(tempHeader, eachHeader)
				}
			}

			// skipping empty rows
			if headerFilled {
				headerFound = true
				header = tempHeader
			}

		} else {
			// body
			tableRow := map[string]string{}
			for headerIndex, eachHeader := range header {
				// checking for empty cols
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
