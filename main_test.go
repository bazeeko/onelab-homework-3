package main

import (
	"reflect"
	"testing"
)

func TestTopWords(t *testing.T) {
	testCases := []struct {
		str    string
		n      int
		output []string
	}{
		{"hi hi alo how are you", 3, []string{"hi", "alo", "are"}},
		{"i am trump. i am politician. i love america. i hate biden", 3, []string{"i", "am", "america"}},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.", 5, []string{"ut", "ad", "adipiscing", "aliqua", "aliquip"}},
	}

	for _, tc := range testCases {
		if result := topWords(tc.str, tc.n); !reflect.DeepEqual(result, tc.output) {
			t.Errorf("TestTopWords => got: %v, want %v", result, tc.output)
		}
	}
}
