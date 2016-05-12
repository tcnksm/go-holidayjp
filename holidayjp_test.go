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

func TestHoliday(t *testing.T) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	cases := []struct {
		input  time.Time
		expect string
		error  bool
	}{
		{
			input:  time.Date(2016, time.February, 11, 0, 0, 0, 0, location),
			expect: "建国記念の日",
			error:  false,
		},
		{
			input:  time.Date(2017, time.February, 11, 0, 0, 0, 0, location),
			expect: "建国記念の日",
			error:  false,
		},
		{
			input:  time.Date(2016, time.February, 16, 0, 0, 0, 0, location),
			expect: "",
			error:  true,
		},
	}

	for i, tc := range cases {
		output, err := Holiday(tc.input)
		if tc.error && err == nil {
			t.Fatalf("#%d expects error, but nil returned", i)
		}
		if output != tc.expect {
			t.Fatalf("#%d expects %v to be eq %v", i, output, tc.expect)
		}
	}
}
