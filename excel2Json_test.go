package excel2json

import (
	"os"
	"testing"
)

var tests = []struct {
	filepath     string
	expectedJson string
}{
	{
		"xlsx/emptyLines.xlsx", `[{"headerA":"A4","headerB":"B4"}]`,
	},
	{
		"xls/emptyLines.xls", `[{"headerA":"A4","headerB":"B4"}]`,
	},
	{
		"csv/emptyLines.csv", `[{"headerA":"A4","headerB":"B4"}]`,
	},
	{
		"ods/emptyLines.ods", `[{"headerA":"A4","headerB":"B4"}]`,
	},
}

func TestToJson(t *testing.T) {
	for _, test := range tests {
		file, err := os.Open(test.filepath)
		if err != nil {
			t.Errorf("error opening %s: %v", test.filepath, err)
			return
		}
		defer file.Close()

		excelFile := ExcelFile{File: file}
		json, err := excelFile.ToJson()
		if err != nil {
			t.Errorf("error parsing %s: %v", test.filepath, err)
		}

		if json != test.expectedJson {
			t.Errorf("Json(%s) = %s; want %s", test.filepath, json, test.expectedJson)
		}
	}
}

func BenchmarkToJson(b *testing.B) {
	test := tests[0]

	file, err := os.Open(test.filepath)
	if err != nil {
		b.Errorf("error opening %s: %v", test.filepath, err)
		return
	}
	defer file.Close()

	excelFile := ExcelFile{File: file}
	
	for b.Loop() {
		excelFile.ToJson()
	}
}
