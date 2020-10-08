package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func ntpGetTime(ntpHost string) time.Time {

	ntpTime, err := ntp.Time(ntpHost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", ntpTime)
	return ntpTime
}
