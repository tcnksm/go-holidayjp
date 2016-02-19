package main

import (
	"os"
	"time"

	"github.com/tcnksm/go-holidayjp"
)

func main() {
	if holidayjp.IsHoliday(time.Now()) {
		os.Exit(0)
	}
	os.Exit(1)
}
