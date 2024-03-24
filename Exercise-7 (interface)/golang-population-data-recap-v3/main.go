package main

import (
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	for _, d := range data {
		parts := strings.Split(d, ";")

		name := parts[0]
		age, _ := strconv.Atoi(parts[1])
		address := parts[2]
		height := parts[3]
		isMarried := parts[4]

		person := map[string]interface{}{
			"name":    name,
			"age":     age,
			"address": address,
		}

		if len(height) > 0 {
			height, _ := strconv.ParseFloat(height, 64)
			person["height"] = height
		}

		if len(parts) > 4 && len(isMarried) > 0 {
			isMarried, _ := strconv.ParseBool(isMarried)
			person["isMarried"] = isMarried
		}

		result = append(result, person)
	}

	return result
}
