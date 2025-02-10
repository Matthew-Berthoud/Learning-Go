package main

import (
	"io"
	"os"
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

func (l League) MatchResult(aTeam string, aScore int, bTeam string, bScore int) {
	if aScore > bScore {
		l.Wins[aTeam] += 1
	} else if bScore > aScore {
		l.Wins[bTeam] += 1
	} // else it's a tie
}

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

// Define an interface called Ranker that has a
// single method called Ranking that returns a slice of strings.

type Ranker interface {
	Ranking() []string
}

// Write a function called RankPrinter with two parameters,
// the first of type Ranker and
// the second of type io.Writer.
// Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline separating each result.
// Call this function from main

func RankPrinter(ranker Ranker, writer io.Writer) {
	for _, v := range ranker.Ranking() {
		io.WriteString(writer, v+"\n")
	}
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

	nfl.MatchResult("commanders", 1000, "eagles", 0)
	nfl.MatchResult("chiefs", 50, "commanders", 51)
	nfl.MatchResult("chiefs", 22, "eagles", 40)

	RankPrinter(nfl, os.Stdout)
}
