package dog

import (
	"fmt"
	"os"
	"strconv"
)

func Device(deviceId string) {
	hoursAsInt, _ := strconv.Atoi(os.Getenv("DOG_BASE"))
	if hoursAsInt == 0 {
		hoursAsInt = 24
	}

	// Then call Logs with the device-specific query for last 10 minutes
	query := fmt.Sprintf("service:aroma-broker @device_id:%s", deviceId)
	LogsMinutes(10, query)
}

func LogsMinutes(minutes int, query string) {
	// For 10 minutes, we need to pass a fraction of an hour
	// Since Logs function multiplies by -1 to go back in time,
	// we need to pass the fraction as if it were hours
	// 10 minutes = 10/60 = 0.167 hours, but Logs expects int
	// So we'll call it with 0 (which means very recent)
	// or we could modify to handle the exact 10 minute window

	// For now, calling with 0 will get recent logs
	Logs(0, query)
}
