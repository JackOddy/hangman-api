package game

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/sync/syncmap"
	"net/http"
)

var (
	Games = syncmap.Map{}
)

func Index(res http.ResponseWriter, req *http.Request) {
	games := []interface{}{}

	Games.Range(func(id, game interface{}) bool {
		games = append(games, game)
		return true
	})

	json.NewEncoder(res).Encode(games)
}

func Read(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	game, ok := Games.Load(id)

	if !ok {
		json.NewEncoder(res).Encode("Could not find game")
	}
	json.NewEncoder(res).Encode(game)
}

func Create(res http.ResponseWriter, req *http.Request) {

	game := CreateGame()

	fmt.Println(game)

	if _, err := Games.LoadOrStore(game.Id, &game); err {
		json.NewEncoder(res).Encode("Error creating Game")
	}
	json.NewEncoder(res).Encode(game)
}
func Update(res http.ResponseWriter, req *http.Request) {
}
