package number

import (
	"regexp"
	"strings"
)

var numbersRegex = regexp.MustCompilePOSIX("[0-9]+")

func ExtractNumber(text string) string {
	nums := numbersRegex.FindAllString(text, -1)
	return strings.Join(nums, "")
}
