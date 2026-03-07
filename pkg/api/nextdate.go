package api

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	dateOnly := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateOnly.After(nowOnly)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(dateLayout, dstart)
	if err != nil {
		return "", errors.New("Invalid date format, expected YYYYMMDD")
	}

	ruleParts := strings.Fields(strings.TrimSpace(repeat))
	if len(ruleParts) == 0 || ruleParts[0] == "" {
		return "", errors.New("No repeat rule specified — task will be deleted")
	}

	switch ruleParts[0] {
	case "d":
		if len(ruleParts) != 2 {
			return "", errors.New("Invalid format for repeat rule d")
		}

		var days int
		_, err := fmt.Sscanf(ruleParts[1], "%d", &days)
		if err != nil {
			return "", errors.New("Invalid day number in repeat rule d")
		}
		if days < 1 || days > 400 {
			return "", errors.New("Day number must be between 1 and 400")
		}

		for {
			date = date.AddDate(0, 0, days)
			if afterNow(date, now) {
				break
			}
		}

	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}

	default:
		return "", errors.New("Unknown repeat rule")
	}

	return date.Format(dateLayout), nil
}
