package oddmap

import (
	"fmt"
	"strings"
)

func PlainMap(m *OddMap) string {
	fmt.Println(m.CurrentMap)
	m.regenerateBoard()
	return formatToBoard(m.CurrentMap)
}

func PlayerMap(m *OddMap) string {
	m.regenerateBoard()
	fmt.Println(m.CurrentMap)
	m.CurrentMap[m.Player.Y][m.Player.X] = "X"
	return formatToBoard(m.CurrentMap)
}

func formatToBoard(m [][]string) string {
	var e string
	for x := range m {
		e += (strings.Join(m[x], "") + "\n")
	}
	return e
}
