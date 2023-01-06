package shared

import "time"

const (
	FormatDate = `2006-01-02`

	FormatDateTime = `2006-01-02 15:04:05`
)

func StartDateString(date string) string {
	startDate, _ := time.Parse(FormatDate, date)
	return startDate.Format(FormatDateTime)
}

func EndDateString(date string) string {
	endDate, _ := time.Parse(FormatDate, date)
	endDate = endDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	return endDate.Format(FormatDateTime)
}
