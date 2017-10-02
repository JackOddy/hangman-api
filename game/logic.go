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

func Index(res http.ResponseWriter, req *http.Request) {}
func Read(res http.ResponseWriter, req *http.Request) {

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
	guess := mux.Vars(req)["guess"]
	println(guess)
}
