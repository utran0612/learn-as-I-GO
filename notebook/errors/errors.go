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

func ParseTime(stringTime string) bool{
	timeComponents := strings.Split(stringTime,":")
	if len(timeComponents) == 3{
		hour := timeComponents[0]
		minute := timeComponents[1]
		second := timeComponents[2]

		intHour, e1:= strconv.Atoi(hour)
		intMinute, e2 := strconv.Atoi(minute)
		intSecond, e3 := strconv.Atoi(second)
		if (e1 != nil && e2 != nil && e3 != nil){
			fmt.Print(errors.New("cannot parse string to time"))
			return false
		}else{
			time := Time{intHour,intMinute,intSecond}
			fmt.Printf("Time is %v %v %v",time.hour,time.minute,time.second) 
			return true 
		}	
	}else{
		fmt.Print(errors.New("cannot parse string to time"))
		return false
	}
}

func main() {
	st := "12:4:5"
	ParseTime(st)
}