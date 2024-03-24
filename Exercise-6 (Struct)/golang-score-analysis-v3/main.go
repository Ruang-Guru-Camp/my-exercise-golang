package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}

	min := s.Grades[0]
	max := s.Grades[0]
	sum := 0

	for _, grade := range s.Grades {
		if grade < min {
			min = grade
		}
		if grade > max {
			max = grade
		}
		sum += grade
	}

	average := float64(sum) / float64(len(s.Grades))
	return average, min, max
}

func main() {
	school := School{
		Name:    "ABC School",
		Address: "123 Main Street",
	}

	school.AddGrade(100, 90, 100, 90, 100, 90)
	avg, min, max := Analysis(school)
	fmt.Println("Average:", avg)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
