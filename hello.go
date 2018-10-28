package main

import "fmt"

func main() {
	weekDays := map[string]string{
		"mo": "Monday",
		"tu": "Tuesday",
		"we": "Wednwsday",
		"th": "Thursday",
		"fr": "Friday",
		"sa": "Saturday",
		"su": "Sunday",
	}
	selectedDays := [3]string{
		"mo",
		"su",
		"te",
	}
	for i := 0; i < 3; i++ {
		if day, ok := weekDays[selectedDays[i]]; ok {
			fmt.Println(i, day)
		} else {
			fmt.Println("There is no such week day as", selectedDays[i])
		}
	}
	fmt.Println("Hello, World!")
}
