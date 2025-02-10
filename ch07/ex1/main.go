package main

import (
	"fmt"
)

type Team struct {
	teamName    string
	playerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func main() {
	chiefs := Team{
		teamName:    "Kansas City Chiefs",
		playerNames: []string{"Travis Kelce", "Patrick Mahomes"},
	}
	eagles := Team{
		teamName:    "Philadelphia Eagles",
		playerNames: []string{"Jalen Hurts", "Saquon Barkley"},
	}
	nfl := League{
		Teams: []Team{eagles, chiefs},
		Wins: map[string]int{
			"chiefs": 0,
			"eagles": 1,
		},
	}

	fmt.Println(chiefs)
	fmt.Println(eagles)
	fmt.Println(nfl)
}
