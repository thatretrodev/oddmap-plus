package main

import (
	"C"
	"TA/oddmap"
	"fmt"
)

func main() {
	x := oddmap.OddMap{}
	x.Set("Test", 20, 4)
	x.LoadMap("test.txt")
	fmt.Println(x.Name)
	for i := range x.Board {
		fmt.Println(x.Board[i])
	}
	fmt.Println(x.Player.X)
	fmt.Println(x.PlainMap())
	fmt.Println(x.PlayerMap())
	x.SetPlayerPosition(18, 0)
	fmt.Println(x.Player.X, x.Player.Y)
	fmt.Println(x.PlayerMap())
	x.MovePlayer(oddmap.Up, 0)
	fmt.Println(x.Player.X, x.Player.Y)
	fmt.Println(x.PlayerMap())
	fmt.Println(x.PlainMap())

}
