package main

import (
	"testing"
)

func Test_extractDigits(t *testing.T) {
	tests := []struct {
		line          string
		expectedFirst int
		expectedLast  int
	}{
		{
			line:          "5ijiewofj7",
			expectedFirst: 5,
			expectedLast:  7,
		},
		{
			line:          "zzza8azzz",
			expectedFirst: 8,
			expectedLast:  8,
		},
		{
			line:          "j345837938x",
			expectedFirst: 3,
			expectedLast:  8,
		},
		{
			line:          "fzrpfhbfvj6dbxbtfs7twofksfbshrzkdeightwoqg",
			expectedFirst: 6,
			expectedLast:  7,
		},
		{
			line:          "vnrnkfp6",
			expectedFirst: 6,
			expectedLast:  6,
		},
		{
			line:          "0vnrnkfp6",
			expectedFirst: 0,
			expectedLast:  6,
		},
		{
			line:          "2vnrnkfp61",
			expectedFirst: 2,
			expectedLast:  1,
		},
		{
			line:          "9vnrnkfp64",
			expectedFirst: 9,
			expectedLast:  4,
		},
	}
	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			actualFirst, actualLast, err := extractDigits(test.line)
			if err != nil {
				t.Error(err)
			}

			if actualFirst != test.expectedFirst || actualLast != test.expectedLast {
				t.Errorf("expectedFirst: %v, actualFirst: %v | expectedLast: %v, actualLast: %v",
					test.expectedFirst,
					actualFirst,
					test.expectedLast,
					actualLast)
			}
		})
	}
}
