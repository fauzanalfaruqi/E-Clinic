package utils

import (
	"strconv"
	"strings"
)

func SanitizeStatusQuery(status string) []string {
	var arrStatus []string

	if status != "" {
		arrStr := strings.Split(status, "#")
		for _, v := range arrStr {
			v = strings.ToUpper(v)
			switch v {
			case "DONE":
				arrStatus = append(arrStatus, v)
				continue
			case "WAITING":
				arrStatus = append(arrStatus, v)
				continue
			case "CANCELED":
				arrStatus = append(arrStatus, v)
				continue
			default:
				continue
			}
		}
	}

	return arrStatus
}

func SanitizeDowQuery(dayOfWeek string) []int {
	var arrDow []int

	if dayOfWeek != "" {
		strarr := strings.Split(dayOfWeek, "#")
		for _, v := range strarr {
			dow, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			arrDow = append(arrDow, dow)
		}
	}

	return arrDow

}

func ValidateStartEndDate(startDate, endDate string) (string, string, error) {
	var err error

	if startDate == "" {
		startDate = GetDefaultStartDate()
	} else {
		startDate, err = FormatDate(startDate)
		if err != nil {
			return "", "", err
		}
	}

	if endDate == "" {
		endDate = GetDefaultEndDate()
	} else {
		endDate, err = FormatDate(endDate)
		if err != nil {
			return "", "", err
		}
	}

	return startDate, endDate, nil
}
