package xlsx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func Json(file *os.File) (string, error) {
	f, err := excelize.OpenReader(file, excelize.Options{
		RawCellValue: true,
	})
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}()

	sheets := f.GetSheetList()
	sheetName := sheets[0]

	rows, err := f.Rows(sheetName)
	if err != nil {
		return "", fmt.Errorf("error getting rows iterator: %v", err)
	}

	headerFound := false
	var tableMap []map[string]string
	var headers []string

	for rows.Next() {
		rowValues, err := rows.Columns()
		if err != nil {
			return "", fmt.Errorf("error reading row: %v", err)
		}

		if len(rowValues) > 0 {
			if !headerFound {
				// header
				headers = rowValues
				headerFound = true
			} else {
				// body
				tableMap = append(tableMap, getBodyRow(headers, rowValues))
			}
		}
	}

	jsonSheet, err := json.Marshal(tableMap)
	if err != nil {
		return "", fmt.Errorf("error parsing to json: %v", err)
	}

	return string(jsonSheet), nil
}

func getBodyRow(headers []string, row []string) map[string]string {
	eachRow := map[string]string{}
	for headersIndex, header := range headers {
		// checking for empty cols
		if headersIndex < len(row) {
			eachRow[header] = row[headersIndex]
		} else {
			eachRow[header] = ""
		}
	}
	return eachRow
}
