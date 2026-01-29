package xlsx2json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

func Json(xlsReader io.Reader) (string, error) {
	f, err := excelize.OpenReader(xlsReader)
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

	rows.Next()
	headers, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("error reading headers: %v", err)
	}

	tableMap := []map[string]string{}
	for rows.Next() {
		rowValues, err := rows.Columns()
		if err != nil {
			return "", fmt.Errorf("error reading row: %v", err)
		}

		eachRow := map[string]string{}
		for headersIndex, header := range headers {
			if headersIndex < len(rowValues) {
				eachRow[header] = rowValues[headersIndex]
			} else {
				eachRow[header] = ""
			}
		}
		tableMap = append(tableMap, eachRow)
	}

	jsonSheet, err := json.Marshal(tableMap)
	if err != nil {
		return "", fmt.Errorf("error parsing to json: %v", err)
	}

	return string(jsonSheet), nil
}
