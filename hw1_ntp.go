package main

import (
	"fmt"
	"time"
)

func main() {
	var ntpTime time.Time
	var ntpHost string
	fmt.Println("Enter ntp host: ")
	fmt.Scanln(&ntpHost)

	ntpTime = ntpGetTime(ntpHost)
	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("+++++++++++++++++++++++++++++++")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

}
