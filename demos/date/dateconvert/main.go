package main

import (
	"fmt"
	"time"
)

func main() {
	date, err := DateToDate(time.RFC3339, time.RFC822, "2022-06-07T16:12:17+08:00", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(date)

	date, err = DateToDate(time.RFC3339, "2006-01-02 15:04:05", "2022-06-07T16:12:17+08:00", 7)
	if err != nil {
		panic(err)
	}
	fmt.Println(date)
	date, err = DateToDate(time.RFC3339, "2006-01-02 15:04:05-07", "2022-06-07T16:12:17+08:00", 7)
	if err != nil {
		panic(err)
	}
	fmt.Println(date)

	d, err := time.Parse("2006-01-02 15:04:05.999999-07", "2022-07-28 12:58:20.28+03")
	if err != nil {
		return
	}
	fmt.Println(d.Format(time.RFC3339Nano))
	fmt.Println(d.Unix())
}

func DateToDate(fromLayout string, toLayout string, date string, toZone int) (string, error) {
	parse, err := time.Parse(fromLayout, date)
	if err != nil {
		return "", err
	}
	_, curZone := parse.Zone() // seconds
	curZoneHour := curZone / 3600
	zoneOffset := toZone - curZoneHour
	// fmt.Println(curZoneHour)
	// fmt.Println(zoneOffset)
	zoneTime := parse.Add(time.Duration(zoneOffset) * time.Hour)
	return zoneTime.Format(toLayout), nil
}
