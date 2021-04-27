package util

import "time"

func ConvertTimeToString(value *time.Time) string {
	return value.Format(time.RFC3339)
}
