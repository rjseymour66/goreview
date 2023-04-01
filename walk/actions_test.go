package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	type testCase struct {
		file     string
		ext      string
		size     int64
		expected bool
	}

	tt := map[string]testCase{
		"FilterNoExtension": {
			file:     "testdata",
			ext:      "",
			size:     0,
			expected: true,
		},
		"FilterExtensionMatch": {
			file:     "testdata/dir.log",
			ext:      ".log",
			size:     0,
			expected: false,
		},
		"FilterExtensionNoMatch": {
			file:     "testdata/dir.log",
			ext:      ".sh",
			size:     0,
			expected: true,
		},
		"FilterExtensionSizeMatch": {
			file:     "testdata/dir.log",
			ext:      ".log",
			size:     10,
			expected: false,
		},
		"FilterExtensionSizeNoMatch": {
			file:     "testdata",
			ext:      "",
			size:     20,
			expected: true,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// act
			// get file info
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Errorf("cannot get %s FileInfo", tc.file)
			}
			// got := filterOut(path string, ext string, minSize int64, info os.FileInfo) bool
			got := filterOut(tc.file, tc.ext, tc.size, info)

			// assert
			if got != tc.expected {
				t.Errorf("got:  %t\nwant:  %t", got, tc.expected)
			}
		})
	}
}
