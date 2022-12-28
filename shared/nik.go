package shared

import "regexp"

var (
	NIKFormat = regexp.MustCompile(`^\d{16}$`)
)

func IsNIKFormat(nik string) bool {
	return NIKFormat.MatchString(nik)
}
