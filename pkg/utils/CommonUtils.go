package utils

import (
	"regexp"
	"time"
)

func IsDateTimeValid(stringDate string) bool {
	_, err := time.Parse("2023-02-19 10:00:00", stringDate)
	return err == nil
}

func IsDateTimeValidRegexp(stringDate string) bool {
	_, err := regexp.MatchString("^(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2}):(\\d{2})", stringDate)
	return err == nil
}

func RegexpString(pattern string, value string) bool {
	isMatched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false
	}
	return isMatched
}
