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
	{"test.xlsx", `[{"headerA":"a2","headerB":"b2","headerC":"c2"},{"headerA":"a3","headerB":"b3","headerC":"c3"},{"headerA":"a4","headerB":"b4","headerC":"c4"}]`},
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