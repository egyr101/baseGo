package main

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
	return nil
}

func ratingSort(players []Player) []Player {
	return nil
}

func gmSort(players []Player) []Player {
	return nil
}