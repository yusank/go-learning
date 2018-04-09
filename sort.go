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
	foo()
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

type pair struct {
	key   string
	value string
}

func foo() {
	var pairs []pair

	a := pair{
		key: "aaa",
	}

	pairs = append(pairs, a)

	for i := 0; i < 10; i++ {
		var p pair
		p.key = string(byte(67 + i))
		p.value = fmt.Sprint(i + 123)

		pairs = append(pairs, p)
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})
	for _, v := range pairs {
		fmt.Printf("%+v \n", v)
	}
}
