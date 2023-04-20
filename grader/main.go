package main

import "fmt"

func main() {
	// fmt.Println(gradeStudent(0))
	scores := []int{90, 48, 63, 75, 70}
	for _, score := range scores {
		fmt.Println(gradeStudent(score))
	}

	// with nice format
	scoresV2 := map[string]int{
		"A": 90,
		"B": 48,
		"C": 63,
		"D": 75,
		"E": 70,
	}

	var names []string
	for name := range scoresV2 {
		names = append(names, name)
	}
	fmt.Println(names)

	for k, name := range names {
		if k == 0 {
			fmt.Println("Team, Grade")
		}
		fmt.Printf("%s, %s\n", name, gradeStudent(scoresV2[name]))

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
