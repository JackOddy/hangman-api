package game

import (
	"reflect"
	"testing"
)

func TestCreateGameGeneration(t *testing.T) {
	testGeneratedStuff := map[string]string{"Word": "TestWord", "Id": "TestId"}

	SillyStuff["Word"] = func() string { return "TestWord" }
	SillyStuff["Id"] = func() string { return "TestId" }

	game := CreateGame()

	for key, word := range testGeneratedStuff {
		value := getField(game, key)

		if value != word {
			t.Errorf("expected a newly created game %s to have a value of %s. Instead got %s",
				key, word, value)
		}
	}

}

func getField(g Game, field string) string {
	r := reflect.ValueOf(g)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func TestTakeTurnSuccess(t *testing.T) {
	SillyStuff["Word"] = func() string { return "ABC" }
	SillyStuff["Id"] = func() string { return "TestId" }

	game := CreateGame()
	guess := Guess{"A"}

	game.TakeTurn(&guess)

	if game.Complete == true {
		t.Error("Expected game to not be completed after first turn")
	}
	if game.Strikes != 0 {
		t.Errorf("Expected 0 strikes on first successful play, instead got %d", game.Strikes)
	}
	if game.GuessedLetters[[]rune(guess.Letter)[0]] != true {
		t.Error("Game did not update guessed letters")
	}
}

func TestTakeTurnFail(t *testing.T) {
	SillyStuff["Word"] = func() string { return "ABC" }
	SillyStuff["Id"] = func() string { return "TestId" }

	game := CreateGame()
	guess := Guess{"D"}

	game.TakeTurn(&guess)

	if game.Complete == true {
		t.Error("Expected game to not be completed after first turn")
	}
	if game.Strikes != 1 {
		t.Errorf("Expected 1 strikes on first successful play, instead got %d", game.Strikes)
	}
	if game.GuessedLetters[[]rune(guess.Letter)[0]] != true {
		t.Error("Game did not update guessed letters")
	}
}

// If I had more time I would like to investigate table driven testing a lot more
// After watching a talk by Mitchel Hashimoto about testing in Golang it is
// So different and refreshing to just write your tests as code. I definitely do
// not have the nuances of it all figured out yet, but it is enjoyable.
// Following advice gained from this talk I have attempted table driven assertions
// in the first test here, as well as replacing external package calls by placing
// them in an exported variable. To further this testing I would plan to cover all
// cases using only the exported API (to confirm that it is useful) and would then
// look to moving onto integration tests, similar to how I tackled this for the
// ipChecker
