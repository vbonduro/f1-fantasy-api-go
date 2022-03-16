package main

import (
	"fmt"

	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func main() {
	// user := os.Getenv("F1_USER")
	// password := os.Getenv("F1_PASSWORD")
	// api, err := f1fantasy.NewAuthenticatedApi(user, password)
	// if err != nil {
	// 	panic(err)
	// }

	api := f1fantasy.NewApi()

	players, err := api.GetPlayers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", players)
}
