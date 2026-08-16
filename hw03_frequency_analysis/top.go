package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const topWordsCount = 10

func Top10(text string) []string {
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[word]++
	}

	words := make([]string, 0, len(counts))
	for word := range counts {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] != counts[words[j]] {
			return counts[words[i]] > counts[words[j]]
		}
		return words[i] < words[j]
	})

	if len(words) > topWordsCount {
		words = words[:topWordsCount]
	}

	return words
}
