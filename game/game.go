package game

type Game struct {
	word, id       string
	guessedLetters []string
	turnsTaken     int
	complete       bool
}
