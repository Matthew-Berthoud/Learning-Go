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

// Add two methods to League

// MatchResult.
// It takes four parameters:
// the name of the first team,
// its score in the game,
// the name of the second team,
// and its score in the game.
// This method should update the Wins field in League

func (l League) MatchResult(aTeam string, aScore int, bTeam string, bScore int) {
	if aScore > bScore {
		l.Wins[aTeam] += 1
	} else if bScore > aScore {
		l.Wins[bTeam] += 1
	} // else it's a tie
}

// Ranking
// returns a slice of the team names in order of wins.

// Build your data structures and call these methods from the main function in your program using some sample data

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
			"eagles": 0,
		},
	}

	fmt.Println(nfl)

	nfl.MatchResult("chiefs", 22, "eagles", 40)

	fmt.Println(nfl)
}
