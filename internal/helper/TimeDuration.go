package helper

import (
	"fmt"
	"time"
)

func FormatDurationID(duration time.Duration) string {
	// Bulatkan durasi ke milidetik
	roundedDuration := duration.Round(time.Millisecond)

	// Tentukan format berdasarkan unit waktu
	if roundedDuration < time.Second {
		return fmt.Sprintf("%dms", roundedDuration.Milliseconds())
	} else if roundedDuration < time.Minute {
		seconds := roundedDuration.Seconds()
		return fmt.Sprintf("%.0f detik", seconds)
	} else if roundedDuration < time.Hour {
		minutes := roundedDuration.Minutes()
		return fmt.Sprintf("%.0f menit", minutes)
	} else {
		hours := roundedDuration.Hours()
		return fmt.Sprintf("%.0f jam", hours)
	}
}
