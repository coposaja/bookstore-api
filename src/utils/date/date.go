package date

import "time"

const (
	dateFormat = "2006-01-02T15:04:05Z"
)

// GetNow return current UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return current UTC in formatted string
func GetNowString() string {
	return GetNow().Format(dateFormat)
}
