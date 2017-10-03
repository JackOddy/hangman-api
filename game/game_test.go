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
