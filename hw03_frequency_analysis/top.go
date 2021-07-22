package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

func Top10(input string) []string {
	wordsRegExp := regexp.MustCompile(`[\s\n,.!?]+`)
	words := wordsRegExp.Split(strings.ToLower(input), -1)
	count := map[string]int{}
	var topWord []string

	for _, word := range words {
		if len(word) == 0 || word == "-" {
			continue
		}

		count[word]++

		if count[word] == 1 {
			topWord = append(topWord, word)
		}
	}

	sort.Slice(topWord, func(i, j int) bool {
		a := topWord[i]
		b := topWord[j]

		if count[a] == count[b] {
			return strings.Compare(a, b) == -1
		}

		return count[a] > count[b]
	})

	resLen := int(math.Min(float64(len(topWord)), 10))

	return topWord[:resLen]
}
