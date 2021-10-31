package main

import (
	"regexp"
	"sort"
	"strings"
)

// topWords counts the words in the string s
// and returns a slice of string with n most frequent words.
// All the words in the result will be lowercased.
//
// A word in this case is a set of latin letters.
// Any other charachter will not be included in counting.
//
// If there are any other charachters in s, they will be counted as separators.
//
func topWords(s string, n int) (result []string) {
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[a-z]+`)
	matches := re.FindAllString(s, -1)

	freq := make(map[string]int)
	for _, v := range matches {
		freq[v]++
	}

	keys := make([]string, 0, len(freq))
	for key := range freq {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if freq[keys[i]] == freq[keys[j]] {
			return keys[i] < keys[j]
		}

		return freq[keys[i]] > freq[keys[j]]
	})

	if n > len(keys) {
		n = len(keys)
	}

	return keys[0:n]
}
