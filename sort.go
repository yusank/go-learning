package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Games struct {
	Id     int
	GameId string
	Status int
}

func main() {
	var gameSlice []*Games
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		g := new(Games)
		g.Id = i
		g.GameId = string(byte(77 - i))
		g.Status = rand.Intn(5)

		gameSlice = append(gameSlice, g)
	}

	//sort.Slice(gameSlice, func(i, j int) bool {
	//	return gameSlice[i].Status < gameSlice[j].Status
	//})

	sort.Slice(gameSlice, func(i, j int) bool {
		fmt.Println(gameSlice[i].Status, gameSlice[j].Status)
		is, js := gameSlice[i].Status, gameSlice[j].Status
		if is == 4 {
			is = 0
		}
		if js == 4 {
			js = 0
		}

		return is > js
		//if gameSlice[i].Status > gameSlice[j].Status && gameSlice[i].Status == 4 {
		//	return false
		//}
		//return gameSlice[i].Status > gameSlice[j].Status
	})

	for _, v := range gameSlice {
		fmt.Printf("%+v \n", *v)
	}
}
