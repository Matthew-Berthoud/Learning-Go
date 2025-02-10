package main

import (
	"fmt"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func main() {
	chiefs := Team{
		"Kansas City Chiefs",
		[]string{"Travis Kelce", "Patrick Mahomes"},
	}
	eagles := Team{
		"Philadelphia Eagles",
		[]string{"Jalen Hurts", "Saquon Barkley"},
	}
	nfl := League{
		Name: "National Football League",
		Teams: map[string]Team{
			"chiefs": chiefs,
			"eagles": eagles,
		},
		Wins: map[string]int{
			"chiefs": 0,
			"eagles": 1,
		},
	}

	fmt.Println(nfl)
}
