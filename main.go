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
	/*for i := range x.Board {
		fmt.Println(x.Board[i])
	}*/
	fmt.Println(x.Player.X)
	x.SetPlayerPosition(18, 0)
	fmt.Println(x.Player.X, x.Player.Y)
	fmt.Println(x.Player.X, x.Player.Y)
	var dir string
	var strength int
	for {
		fmt.Println(oddmap.PlayerMap(x))
		fmt.Println(x.Player.X, x.Player.Y)
		fmt.Println("Enter direction and force:")
		fmt.Scanln(&dir)
		fmt.Scanln(&strength)
		x.MovePlayer(dir, strength)
	}
}
