package main

import (
	"fmt"
	"time"
)

func isActiveBad(now, start, stop int) bool {
	return start <= now && now < stop
}

func isActiveGood(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}

func poolBad(delay int) {
	for {
		fmt.Printf("pool use delay bad")
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func poolGood(delay time.Duration) {
	for {
		fmt.Printf("pool use delay good")
		time.Sleep(delay)
	}
}

func ConvertTimeExternal() {
	/**
	Use time.Duration and time.Time in interactions with external systems when possible. For example:

	Command-line flags: flag supports time.Duration via time.ParseDuration
	JSON: encoding/json supports encoding time.Time as an RFC 3339 string via its UnmarshalJSON method
	SQL: database/sql supports converting DATETIME or TIMESTAMP columns into time.Time and back if the underlying driver supports it
	YAML: gopkg.in/yaml.v2 supports time.Time as an RFC 3339 string, and time.Duration via time.ParseDuration.

	When it is not possible to use time.Duration in these interactions, use int or float64 and include the unit in the name of the field.

	For example, since encoding/json does not support time.Duration, the unit is included in the name of the field.
	*/

	type ConfigBad struct {
		Interval int `json:"interval"`
	}

	type ConfigGood struct {
		IntervalMillis int `json:"intervalMillis"`
	}

}

func main() {
	t := time.Now()
	fmt.Println("t.AddDate add to new day with 00:00")
	newDay := t.AddDate(0 /* years */, 0 /* months */, 1 /* days */)
	fmt.Println("t.Add add to 24 hour to time")
	maybeNewDay := t.Add(24 * time.Hour)
	fmt.Printf("newDay %v, maybeNewDat : %v", newDay, maybeNewDay)

	fmt.Println("use time bad why have many risk when compare int")
	isActiveBad(3000, 2000, 1000)
	fmt.Println("use time good, use time package")
	isActiveGood(time.Now(), t, time.Now())

	fmt.Println("convert time with external service")
	ConvertTimeExternal()
}
