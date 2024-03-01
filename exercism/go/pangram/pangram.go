package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	inputCleaned := strings.ToLower(input)
	totalChars := int('z' - 'a' + 1)

	alphaMap := 0b1<<totalChars - 0b1

	bitMap := 0b0

	for _, char := range inputCleaned {
		if 'a' <= char && char <= 'z' {
			bitMap |= 0b1 << (char - 'a')
		}
	}

	return alphaMap&bitMap == alphaMap
}
