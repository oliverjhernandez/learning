package logs

import "fmt"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		unicode := fmt.Sprintf("%U", r)
		switch unicode {
		case "U+2757":
			return "recommendation"
		case "U+1F50D":
			return "search"
		case "U+2600":
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for i, r := range runes {
		if r == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runes := []rune(log)
	if len(runes) > limit {
		return false
	}
	return true
}
