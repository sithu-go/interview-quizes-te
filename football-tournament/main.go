package main

import "fmt"

type team struct {
	name        string
	win         uint
	loss        uint
	draw        uint
	goalFor     uint
	goalAgainst uint
}

type result struct {
	goalFor  uint
	goalDiff int
	score    uint
}

func main() {

	teams := []*team{
		{
			name:        "A",
			win:         5,
			loss:        1,
			draw:        2,
			goalFor:     23,
			goalAgainst: 20,
		},
		{
			name:        "B",
			win:         6,
			loss:        1,
			draw:        1,
			goalFor:     24,
			goalAgainst: 20,
		},
		{
			name:        "C",
			win:         7,
			loss:        1,
			draw:        0,
			goalFor:     24,
			goalAgainst: 20,
		},
		{
			name:        "D",
			win:         6,
			loss:        1,
			draw:        1,
			goalFor:     15,
			goalAgainst: 11,
		},
		{
			name:        "E",
			win:         8,
			loss:        0,
			draw:        1,
			goalFor:     23,
			goalAgainst: 12,
		},
	}

	for k, t := range teams {
		r := getResult(t)
		if k == 0 {
			fmt.Print("Team", "Goal For", "Goal Diff", "Score")
		}
		fmt.Printf("%s , %d , %d , %d", t.name, r.goalFor, r.goalDiff, r.score)
	}

}

func getResult(t *team) *result {

	return &result{
		goalFor:  t.goalFor,
		goalDiff: int(t.goalFor - t.goalAgainst),
		score:    (t.win * 3) + (t.draw + 1),
	}
}
