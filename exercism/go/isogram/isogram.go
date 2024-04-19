package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	counter := make(map[rune]int)
	for _, r := range word {
		lower := unicode.ToLower(r)
		counter[lower]++
		if unicode.IsLetter(lower) && counter[lower] > 1 {
			return false
		}
	}
	return true
}
