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
		Rating:  calculateRating(),
	}
}

func calculateRating() float64 {
	return 0
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