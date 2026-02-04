package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func Json(file *os.File) (string, error) {
	reader := csv.NewReader(file)



	rows, err := reader.ReadAll()
	if err != nil {
		return "", fmt.Errorf("error reading %s: %v", file.Name(), err)
	}
	
	var tableMap []map[string]string
	var headers []string
	headerFound := false

	for _, row := range rows {
		if !headerFound {
			// header
			var tempHeader []string
			headerFilled := false

			// searching for the first filled row
			for _, col := range row {
				eachHeader := col
				if len(eachHeader) > 0 {
					headerFilled = true
					tempHeader = append(tempHeader, eachHeader)
				}
			}

			// skipping empty rows
			if headerFilled {
				headerFound = true
				headers = tempHeader
			}
		} else {
			tableMap = append(tableMap, getBodyRow(headers, row))
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
