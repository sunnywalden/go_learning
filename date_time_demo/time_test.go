package date_time_demo

import (
	"testing"
)

// TestTimeParse test time parse
func TestTimeParse(t *testing.T) {
	testTime := "2020-05-28 19:00:00"
	fmtTime, err := StrToDatetime(testTime, nil)
	if err != nil {
		t.Errorf("Error %s\n", err)
	} else {
		t.Logf("Date: %s\n", fmtTime)
	}
}

// TestWrongTimeFormatParse test time parse for wrong format
func TestWrongTimeFormatParse(t *testing.T) {
	testTime := "20200528 19:00:00"
	fmtTime, err := StrToDatetime(testTime, nil)
	if err != nil {
		t.Errorf("Error %s\n", err)
	} else {
		t.Logf("Date: %s\n", fmtTime)
	}
}
