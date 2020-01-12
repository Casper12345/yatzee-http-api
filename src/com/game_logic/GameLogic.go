package game_logic

import (
	"fmt"
)

type IncomingGame struct {
	DateTime string
	Played   [][] int
}

func Put(game IncomingGame) {
	fmt.Println("mail logic")
	fmt.Println("date time " + game.DateTime)
	fmt.Println("played " + fmt.Sprint(game.Played))
}

func Get(id int) IncomingGame {
	return IncomingGame{DateTime: "12-12-12", Played: [][] int{{1,2,3},{3,4,5}}}
}
