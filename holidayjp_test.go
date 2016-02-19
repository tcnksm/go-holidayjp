package holidayjp

import (
	"log"
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	cases := []struct {
		input  time.Time
		expect bool
	}{
		{
			input:  time.Date(2016, time.February, 11, 0, 0, 0, 0, location),
			expect: true,
		},

		{
			input:  time.Date(2017, time.February, 11, 0, 0, 0, 0, location),
			expect: true,
		},

		{
			input:  time.Date(2016, time.February, 16, 0, 0, 0, 0, location),
			expect: false,
		},
	}

	for i, tc := range cases {
		output := IsHoliday(tc.input)
		if output != tc.expect {
			t.Fatalf("#%d expects %v to be eq %v", i, output, tc.expect)
		}
	}
}
