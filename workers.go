package oddmap

// worker.go is for the functions that do the heavylifting, like map loading and other stuff
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (m *OddMap) LoadMaps(path string) {
	//m.currentM = path
	m.Boards = make(map[string]Map)
	file, err := os.Open(path)
	if err != nil {
		//handle error
		return
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	//m.Board = make([][]string, 0) //clear loaded board
	//scannedlines := make([]string, 0)
	for s.Scan() { //scan all lines in files to be loaded to individual maps
		sLine := strings.ReplaceAll(s.Text(), " ", "")
		y := Map{}
		if sLine[:5] == "START" {
			y.Name = sLine[5:]
			state := true
			for s.Scan() && state {
				sLine := strings.ReplaceAll(s.Text(), " ", "")
				if sLine[:3] == "END" {
					state = false
					break
				}
				y.Map = append(y.Map, strings.Split(sLine, ""))

			}
			fmt.Println(formatToBoard(y.Map))
		}
		m.Boards[y.Name] = y
	}
	fmt.Println(m.Boards)

}

func formatToBoard(m [][]string) string {
	var e string
	for x := range m {
		e += (strings.Join(m[x], "") + "\n")
	}
	return e
}
