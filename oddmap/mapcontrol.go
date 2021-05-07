package oddmap

import "strings"

func PlainMap(m OddMap) string {
	m.regenerateBoard()
	return formatToBoard(m.Board)
}

func PlayerMap(m OddMap) string {
	m.regenerateBoard()
	m.Board[m.Player.Y][m.Player.X] = "X"
	return formatToBoard(m.Board)
}

func formatToBoard(m [][]string) string {
	var e string
	for x := range m {
		e += (strings.Join(m[x], "") + "\n")
	}
	return e
}
