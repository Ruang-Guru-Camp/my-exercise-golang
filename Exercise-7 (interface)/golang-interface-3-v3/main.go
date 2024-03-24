package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	var hour, minute int
	switch time.(type) {
	case string:
		val := time.(string)
		parts := strings.Split(val, ":")
		if len(parts) != 2 {
			return "Invalid input"
		}
		hour, _ = strconv.Atoi(parts[0])
		if parts[0] == "" {
			return "Invalid input"
		}
		if parts[1] == "" {
			return "Invalid input"
		}
		minute, _ = strconv.Atoi(parts[1])

	case []int:
		parts := time.([]int)
		if len(parts) != 2 {
			return "Invalid input"
		}
		hour = parts[0]
		minute = parts[1]

	case map[string]int:
		parts := time.(map[string]int)

		var ok bool
		hour, ok = parts["hour"]
		if !ok {
			return "Invalid input"
		}
		minute, ok = parts["minute"]
		if !ok {
			return "Invalid input"
		}
	case Time:
		hour = time.(Time).Hour
		minute = time.(Time).Minute
	}
	amPm := "AM"
	if hour >= 12 {
		amPm = "PM"
	}
	if hour > 12 {
		hour -= 12
	}
	fmt.Println(hour)
	fmt.Println(minute)
	return fmt.Sprintf("%02d:%02d %s", hour, minute, amPm) // TODO: replace this
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
