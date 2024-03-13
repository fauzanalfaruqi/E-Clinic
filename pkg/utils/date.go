package utils

import (
	"avengers-clinic/pkg/constants"
	"fmt"
	"time"
)

func FormatDate(date string) (string, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", fmt.Errorf(constants.ErrDateFormat)
	}

	return d.Format("2006-01-02"), nil
}

func GetDefaultStartDate() string {
	return GetNow()
}

func GetDefaultEndDate() string {
	return time.Now().AddDate(0, 1, 0).Format("2006-01-02")
}

func GetNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
