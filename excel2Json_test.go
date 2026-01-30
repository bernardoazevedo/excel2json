package xlsx2json

import (
	"os"
	"testing"
)

var tests = []struct {
	filepath     string
	expectedJson string
}{
	{
		"xlsx/test.xlsx", `[{"hA":"a2","hB":"","hC":""},{"hA":"a3","hB":"b3","hC":""},{"hA":"a4","hB":"b4","hC":"c4"},{"hA":"a5","hB":"b5","hC":"c5"},{"hA":"","hB":"b6","hC":""},{"hA":"","hB":"","hC":"c7"},{"hA":"","hB":"","hC":""}]`,
	},
	{
		"xlsx/numbers.xlsx", `[{"Number":"50.5"},{"Number":"120"},{"Number":"12.12"}]`,
	},
	{
		"xls/test.xls", `[{"hA":"a2","hB":"","hC":""},{"hA":"a3","hB":"b3","hC":""},{"hA":"a4","hB":"b4","hC":"c4"},{"hA":"a5","hB":"b5","hC":"c5"},{"hA":"","hB":"b6","hC":""},{"hA":"","hB":"","hC":"c7"},{"hA":"","hB":"","hC":""}]`,
	},
	{
		"xls/numbers.xls", `[{"Number":"50.5"},{"Number":"120"},{"Number":"12.12"}]`,
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
