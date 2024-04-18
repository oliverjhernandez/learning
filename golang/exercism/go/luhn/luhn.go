package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	sum := 0
	noSpaces := strings.ReplaceAll(id, " ", "")

	if len(noSpaces) <= 1 {
		return false
	}

	shouldDouble := len(noSpaces)%2 == 0

	for i, r := range noSpaces {
		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')

		if (i%2 == 0) == shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}
