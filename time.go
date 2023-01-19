package time

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cmarkh/errs"
)

// Yesterday returns yesterday or Friday if today is Monday
func Yesterday() time.Time { //add to a global package "tools"
	t := time.Now()
	if t.Weekday() == 1 {
		t = t.AddDate(0, 0, -3)
	} else {
		t = t.AddDate(0, 0, -1)
	}
	return t
}

// SecondsSinceMidnight returns the number of seconds since midnight local time
func SecondsSinceMidnight(now time.Time) (t float64) {
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local) //midnight in local time
	t = time.Since(midnight).Seconds()
	return
}

// AskDate asks for a date and converts to time.Time
func AskDate() (t time.Time, err error) {
	fmt.Print("Enter Date YYYYMMDD: ")
	reader := bufio.NewReader(os.Stdin)
	d, err := reader.ReadString('\n')
	if err != nil {
		errs.Log(err)
		return
	}
	d = strings.TrimSpace(d)
	t, err = time.Parse("20060102", d)
	if err != nil {
		errs.Log(err)
		return
	}
	return
}

// DateBefore returns if earlier date is before the later date looking only at the date (ie ignores time and time zone)
func DateBefore(earlier, later time.Time) bool {
	if earlier.Year() > later.Year() {
		return false
	}
	if earlier.Month() > later.Month() {
		return false
	}
	if earlier.Month() < later.Month() {
		return true
	}
	if earlier.Month() == later.Month() && earlier.Day() >= later.Day() {
		return false
	}
	return true
}

// DateOnOrBefore returns if earlier date is on or before the later date looking only at the date (ie ignores time and time zone)
func DateOnOrBefore(earlier, later time.Time) bool {
	if earlier.Year() > later.Year() {
		return false
	}
	if earlier.Month() > later.Month() {
		return false
	}
	if earlier.Month() < later.Month() {
		return true
	}
	if earlier.Month() == later.Month() && earlier.Day() > later.Day() {
		return false
	}
	return true
}
