// go-holidayjp is a package for detecting holiday in Japan.
package holidayjp

import (
	"fmt"
	"time"

	"github.com/tcnksm/go-holidayjp/asset"
	"gopkg.in/yaml.v2"
)

//go:generate go-bindata -pkg=asset -o asset/asset.go ./asset/holidays.yml

var ErrNotHoliday = fmt.Errorf("not holiday")

var holidayMap map[string]string

func init() {
	holidayMap = holidays()
}

// HolidayName returns Holiday name of the given time.
// If it's not holiday, returns error(ErrNotHoliday).
func HolidayName(target time.Time) (string, error) {
	t := target.Format("2006-01-02")
	if name, ok := holidayMap[t]; ok {
		return name, nil
	}
	return "", ErrNotHoliday
}

// IsHoliday reports whether the time is holiday.
func IsHoliday(target time.Time) bool {
	t := target.Format("2006-01-02")
	if _, ok := holidayMap[t]; ok {
		return true
	}
	return false
}

func equal(t0, t1 time.Time) bool {
	return t0.Year() == t1.Year() && t0.Month() == t1.Month() && t0.Day() == t1.Day()
}

func holidays() map[string]string {
	var s map[string]string
	buf, _ := asset.Asset("asset/holidays.yml")
	yaml.Unmarshal(buf, &s)
	return s
}
