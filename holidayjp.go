// go-holidayjp is a package for detecting holiday in Japan.
package holidayjp

import (
	"time"

	"github.com/tcnksm/go-holidayjp/asset"
	"gopkg.in/yaml.v2"
)

//go:generate go-bindata -pkg=asset -o asset/asset.go ./asset/holidays.yml

// Holiday returns Holiday name of the given time.
// If it's not holiday, returns error.
func holiday(t time.Time) (string, error) {
	// TODO
	return "", nil
}

// IsHoliday reports whether the time is holiday.
func IsHoliday(target time.Time) bool {
	for dateStr, _ := range holidays() {
		date, _ := time.Parse("2006-01-02", dateStr)
		if equal(target, date) {
			return true
		}
	}
	return false
}

func equal(t0, t1 time.Time) bool {
	return t0.Year() == t1.Year() && t0.Month() == t1.Month() && t0.Day() == t1.Day()
}

func holidays() map[string]string {
	var s map[string]string
	buf, _ := asset.Asset("holidays.yml")
	yaml.Unmarshal(buf, &s)
	return s
}
