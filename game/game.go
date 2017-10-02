package game

import (
	"github.com/Pallinder/go-randomdata"
)

type Game struct {
	Word           string
	Id             string
	GuessedLetters []string
	TurnsTaken     int
	Complete       bool
	active         bool
}

func CreateGame() *Game {
	return &Game{
		randomdata.City(),
		randomdata.SillyName(),
		[]string{},
		0,
		false,
		false,
	}
}
