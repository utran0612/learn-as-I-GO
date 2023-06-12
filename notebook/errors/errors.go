//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

func ParseTime(stringTime string) (Time, error){
	timeComponents := strings.Split(stringTime,":")
	if len(timeComponents) == 3{
		hour := timeComponents[0]
		minute := timeComponents[1]
		second := timeComponents[2]

		intHour, e1:= strconv.Atoi(hour)
		if e1 != nil {
			return Time{},fmt.Errorf("invalid hour component")
		}
		intMinute, e2 := strconv.Atoi(minute)
		if e2 != nil {
			return Time{},fmt.Errorf("invalid minute component")
		}
		intSecond, e3 := strconv.Atoi(second)
		if (e3 != nil){
			return Time{},fmt.Errorf("invalid second components")
		}else{
			time := Time{intHour,intMinute,intSecond}
			return time,nil
		}	
	}else{
		fmt.Print(errors.New("cannot parse string to time"))
		return Time{},fmt.Errorf("invalid number of time components")
	}
}

func main() {
	st := "whoops"
	ParseTime(st)
}