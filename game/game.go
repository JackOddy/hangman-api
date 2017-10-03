package game

import (
	"github.com/Pallinder/go-randomdata"
	"strings"
)

type Game struct {
	Word           string
	Id             string
	Remaining      []rune
	GuessedLetters map[rune]bool
	Strikes        int
	Complete       bool
	Win            bool
}

type Guess struct {
	Letter string
}

func CreateGame() Game {
	word := strings.ToLower(randomdata.City())
	return Game{
		word,
		strings.ToLower(randomdata.SillyName()),
		[]rune(word),
		map[rune]bool{},
		0,
		false,
		false,
	}
}

func (self *Game) TakeTurn(guess *Guess) (finished bool) {
	letter := []rune(guess.Letter)[0]

	if self.play(letter) {
		self.updateRemaining(letter)
		self.GuessedLetters[letter] = true
	} else {
		self.Strikes++
	}

	if self.finished() {
		finished = true
	}
	return
}

func (self *Game) play(letter rune) (included bool) {
	for _, l := range self.Remaining {
		if letter == l {
			return true
		}
	}
	return
}

func (self *Game) updateRemaining(letter rune) {
	var remaining []rune
	for _, l := range self.Remaining {
		if l != letter {
			remaining = append(remaining, l)
		}
	}
	self.Remaining = remaining
}

func (self *Game) finished() (finished bool) {

	if self.Strikes > 10 {
		finished = true
		self.Complete = true
	}

	if self.Remaining == nil {
		self.Win = true
		finished = true
		self.Complete = true
	}
	return
}
