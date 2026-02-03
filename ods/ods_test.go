package ods

import (
	"os"
	"testing"
)

var tests = []struct {
	filepath     string
	expectedJson string
}{
	{
		"test.ods", `[{"hA":"a2","hB":"","hC":""},{"hA":"a3","hB":"b3","hC":""},{"hA":"a4","hB":"b4","hC":"c4"},{"hA":"a5","hB":"b5","hC":"c5"},{"hA":"","hB":"b6","hC":""},{"hA":"","hB":"","hC":"c7"}]`,
	},
	// {
	// 	"numbers.ods", `[{"Number":"50.5"},{"Number":"120"},{"Number":"12.12"}]`,
	// },
	// {
	// 	"emptyLines.ods", `[{"headerA":"A4","headerB":"B4"}]`,
	// },
}

func TestJson(t *testing.T) {
	for _, test := range tests {
		file, err := os.Open(test.filepath)
		if err != nil {
			t.Errorf("error opening %s: %v", test.filepath, err)
		}
		defer file.Close()

		json, err := Json(file)
		if err != nil {
			t.Errorf("error parsing %s: %v", test.filepath, err)
		}

		if json != test.expectedJson {
			t.Errorf("Json(%s) = %s; want %s", test.filepath, json, test.expectedJson)
		}
	}
}

func BenchmarkJson(b *testing.B) {
	test := tests[0]

	file, err := os.Open(test.filepath)
	if err != nil {
		b.Errorf("error opening %s: %v", test.filepath, err)
	}
	defer file.Close()

	for b.Loop() {
		Json(file)
	}
}
