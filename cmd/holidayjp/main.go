package main

import (
	"os"
	"time"

	"github.com/tcnksm/go-holidayjp"
)

const (
	ExitCodeOK int = 0
)

func main() {
	if holidayjp.IsHoliday(time.Now()) {
		os.Exit(0)
	}
	os.Exit(1)
}
