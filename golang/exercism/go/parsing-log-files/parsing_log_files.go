package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	validLogType := regexp.MustCompile(pattern)
	return validLogType.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	pattern := `<[~*-=]*>`
	fields := regexp.MustCompile(pattern)
	split := fields.Split(text, -1)
	return split
}

func CountQuotedPasswords(lines []string) int {
	counter := 0
	for _, s := range lines {
		pattern := `".*(?i)password.*"`
		password := regexp.MustCompile(pattern)
		if password.MatchString(s) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d*`
	eol := regexp.MustCompile(pattern)
	return eol.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\w+)`
	username := regexp.MustCompile(pattern)
	for i, s := range lines {
		user := username.FindStringSubmatch(s)
		if len(user) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", user[1], s)
		}
	}
	return lines
}
