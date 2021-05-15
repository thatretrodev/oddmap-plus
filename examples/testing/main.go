package main

import (
	"fmt"

	oddmap "github.com/fisik-yum/oddmap"
)

func main() {
	x := oddmap.OddMap{}
	x.Set("Test", 20, 4)
	x.LoadMaps("test.txt")
	x.LoadMap("HOME")
	fmt.Println(x.Name)
	/*for i := range x.Board {
		fmt.Println(x.Board[i])
	}*/
	fmt.Println(x.Player.X)
	x.SetPlayerPosition(18, 0)
	fmt.Println(x.Player.X, x.Player.Y)
	fmt.Println(x.Player.X, x.Player.Y)
	fmt.Println(x.PlayerMap())
	fmt.Println(x.PlainMap())
	var dir string
	var strength int
	for {

		fmt.Println(x.PlayerMap())
		fmt.Println(x.Player.X, x.Player.Y)
		fmt.Println("Enter direction and force:")
		fmt.Scanln(&dir)
		fmt.Scanln(&strength)
		x.MovePlayer(dir, strength)
	}
}
