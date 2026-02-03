package ods

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlexJarrah/go-ods"
)

func Json(file *os.File) (string, error) {
	fileStats, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("error getting %s stats: %v", file.Name(), err)
	}

	odsFile, readCloser, err := ods.ReadFrom(file, fileStats.Size())
	if err != nil {
		return "", fmt.Errorf("error reading from %s: %v", file.Name(), err)
	}
	defer readCloser.Close()

	odsFile.Content = ods.Uncompress(odsFile.Content, 20)

	tables := odsFile.Content.Body.Spreadsheet.Table
	table := tables[0]
	rows := table.TableRow

	headerFound := false
	var tableMap []map[string]string
	var headers []string

	for _, eachRow := range rows {
		cells := eachRow.TableCell
		if !headerFound {
			// header
			var tempHeader []string
			headerFilled := false

			// searching for the first filled row
			for _, eachCell := range cells {
				eachHeader := eachCell.P
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
			// body
			tableRow := map[string]string{}
			itemFound := false
			for headerIndex, eachHeader := range headers {
				// checking for empty cells
				if headerIndex < len(cells) {
					cellValue := cells[headerIndex].P
					if len(cellValue) > 0 {
						itemFound = true
						tableRow[eachHeader] = cellValue
					} else {
						tableRow[eachHeader] = ""
					}
				} else {
					tableRow[eachHeader] = ""
				}
			}
			if itemFound {
				tableMap = append(tableMap, tableRow)
			}
		}
	}

	jsonSheet, err := json.Marshal(tableMap)
	if err != nil {
		return "", fmt.Errorf("error parsing to json: %v", err)
	}

	return string(jsonSheet), nil
}
