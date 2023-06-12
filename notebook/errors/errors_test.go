package main

import (
	"fmt"
	"testing"
)

func TestErrors (t *testing.T){
	table := []struct{
		time string
		ok bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}

	for _, time := range table {
		if time.ok && ParseTime(time.time){
			fmt.Print("Pass")
		}else{
			fmt.Println("Doesn't pass")
		}
	}


}