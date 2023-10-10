package datetimePg

import "time"

// 1970-01-01 08:00:00 +0800 CST
var zeroTime = time.Unix(0, 0)

// IsZero reports whether t represents the zero time instant
func IsZero(t time.Time) bool {
	return t.IsZero() || zeroTime.Equal(t)
}
