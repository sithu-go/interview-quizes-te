package main

import "fmt"

func main() {
	// fmt.Println(gradeStudent(0))
	scores := []int{90, 48, 63, 75, 70}
	for _, score := range scores {
		fmt.Println(gradeStudent(score))
	}
}

func gradeStudent(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score <= 89:
		return "B"
	case score >= 70 && score <= 79:
		return "C"
	case score >= 60 && score <= 69:
		return "D"
	case score >= 50 && score <= 59:
		return "E"
	case score >= 0 && score <= 49:
		return "F"
	default:
		return "Invalid Score"
	}
}
