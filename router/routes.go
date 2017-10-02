package router

import (
	"hangman-api/game"
)

var routes = Routes{
	Route{
		"GameIndex",
		"GET",
		"/games",
		game.Index,
	},
	Route{
		"GetGame",
		"GET",
		"/games/{id}",
		game.Read,
	},
	Route{
		"NewGame",
		"POST",
		"/games",
		game.Create,
	},
	Route{
		"TakeTurn",
		"PUT",
		"/games/{id}",
		game.Update,
	},
}
