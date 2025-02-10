package main

import (
	"fmt"
	"sort"
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

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

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
	commanders := Team{
		"Washington Commanders",
		[]string{"Jayden Daniels", "Terry McLaurin"},
	}
	nfl := League{
		Name: "National Football League",
		Teams: map[string]Team{
			"Philadelphia Eagles":   eagles,
			"Kansas City Chiefs":    chiefs,
			"Washington Commanders": commanders,
		},
		Wins: map[string]int{
			"Philadelphia Eagles":   0,
			"Kansas City Chiefs":    0,
			"Washington Commanders": 0,
		},
	}

	fmt.Println(nfl)

	nfl.MatchResult("commanders", 1000, "eagles", 0)
	nfl.MatchResult("chiefs", 50, "commanders", 51)
	nfl.MatchResult("chiefs", 22, "eagles", 40)

	fmt.Println(nfl)

	ranking := nfl.Ranking()

	fmt.Println(ranking)
}
