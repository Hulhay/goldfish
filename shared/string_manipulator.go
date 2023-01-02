package shared

import "strings"

func CreateValue(name string) string {
	val := strings.ReplaceAll(name, " ", "_")
	return strings.ToUpper(val)
}
