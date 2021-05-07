package oddmap

import "strings"

func formatToBoard(m [][]string) string {
	var e string
	for x := range m {
		e += (strings.Join(m[x], "") + "\n")
	}
	return e
}
