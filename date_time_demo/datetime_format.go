package date_time_demo

import (
	"errors"
	"regexp"
	"time"
)

// StrToDatetime input time format should like "2006-01-02 15:04:05"
func StrToDatetime(dateStr string, loc *time.Location) (dt *time.Time, err error) {
	//reg := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2}) (\d{2}):`)
	match, err := regexp.MatchString(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, dateStr)
	if err != nil {
		return &time.Time{}, err
	}

	if !match {
		return &time.Time{}, errors.New("time format not match")
	}

	var location *time.Location
	if loc != nil {
		location = loc
	} else {
		location = time.Local
	}
	realTime, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, location)

	if err != nil {
		return &time.Time{}, err
	}
	//fmt.Print(realTime)
	return &realTime, nil
}