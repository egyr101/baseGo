package main

import (
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func main() {

}

func NewPlayer(name string, goals, misses, assists int) Player {
	return Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
		Rating:  calculateRating(float64(goals), float64(misses), float64(assists)),
	}
}

func calculateRating(goals, misses, assists float64) float64 {
	if misses == 0 {
		return (goals + assists) / 2
	}

	return (goals + assists) / 2 / misses
}

func goalsSort(players []Player) []Player {
	sortedPlayers := slices.Clone(players)

	slices.SortFunc(sortedPlayers, func (a, b Player) int{
		if a.Goals < b.Goals{
			return 1
		} else if a.Goals > b.Goals{
			return -1
		} else{
			return 0
		}
	})

	return sortedPlayers
}

func ratingSort(players []Player) []Player {
	sortedPlayers := slices.Clone(players)

	slices.SortFunc(sortedPlayers, func (a, b Player) int{
		if a.Rating < b.Rating{
			return 1
		} else if a.Rating > b.Rating{
			return -1
		} else{
			return 0
		}
	})

	return sortedPlayers
}

func gmSort(players []Player) []Player {
	sortedPlayers := slices.Clone(players)

	slices.SortFunc(sortedPlayers, func (a, b Player) int{
		if float64(a.Goals) / float64(a.Misses) < float64(b.Goals) / float64(b.Misses){
			return 1
		} else if float64(a.Goals) / float64(a.Misses) > float64(b.Goals) / float64(b.Misses){
			return -1
		} else{
			return 0
		}
	})

	return sortedPlayers
}