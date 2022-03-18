package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func main() {
	league_id, _ := strconv.Atoi(os.Getenv("F1_LEAGUE"))
	user := os.Getenv("F1_USER")
	password := os.Getenv("F1_PASSWORD")
	api, err := f1fantasy.NewAuthenticatedApi(user, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nPlayers")
	players, err := api.GetPlayers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", players)

	fmt.Println("\nLeaderboard")
	leaderboard, err := api.GetLeagueLeaderboard(league_id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", leaderboard)
}
