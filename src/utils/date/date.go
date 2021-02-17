package date

import "time"

// GetNow return current UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}
