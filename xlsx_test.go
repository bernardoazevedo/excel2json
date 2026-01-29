package xlsx2json

import (
	"fmt"
	"os"
	"testing"
)

var tests = []struct {
	xlsxFilepath string
	expectedJson string
}{
	{
		"test.xlsx", `[{"hA":"a2","hB":"","hC":""},{"hA":"a3","hB":"b3","hC":""},{"hA":"a4","hB":"b4","hC":"c4"},{"hA":"a5","hB":"b5","hC":"c5"},{"hA":"","hB":"b6","hC":""},{"hA":"","hB":"","hC":"c7"},{"hA":"","hB":"","hC":""}]`,
	},
}

func TestJson(t *testing.T) {
	for _, test := range tests {
		file, err := os.Open(test.xlsxFilepath)
		if err != nil {
			t.Errorf("error opening %s: %v", test.xlsxFilepath, err)
		}

		json, err := Json(file)
		if err != nil {
			t.Errorf("error parsing %s: %v", test.xlsxFilepath, err)
		}

		if json != test.expectedJson {
			t.Errorf("Json(%s) = %s; want %s", test.xlsxFilepath, json, test.expectedJson)
		}
	}
}

func BenchmarkJson(b *testing.B) {
	test := tests[0]

	file, err := os.Open(test.xlsxFilepath)
	if err != nil {
		fmt.Printf("error opening %s: %v", test.xlsxFilepath, err)
	}
	
	for i := 0; i < b.N; i++ {
		Json(file)
	}
}