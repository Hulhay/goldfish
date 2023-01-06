package shared

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func CreateValue(name string) string {
	val := nonAlphanumericRegex.ReplaceAllString(name, "_")
	return strings.ToUpper(val)
}
