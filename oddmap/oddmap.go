package oddmap

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Up    = "UP"
	Down  = "DOWN"
	Right = "RIGHT"
	Left  = "LEFT"
)

type Map struct {
	Name string
	Map  [][]string
}

type position struct {
	X int
	Y int
}

type OddMap struct {
	Name           string
	width          int
	height         int
	Boards         map[string]Map
	Player         position
	initState      bool
	CurrentMap     [][]string
	currentMapName string
}

func (m *OddMap) Set(name string, width int, height int) {
	if m.initState {
		log.Print("ERROR: CANNOT REINITIALIZE BOARD")
	} else {
		m.Name = name
		m.width = width
		m.height = height
		m.initState = true
		//code
	}
}

func (m *OddMap) SetPlayerPosition(x int, y int) {
	m.Player.X = x
	m.Player.Y = y
}

func (m *OddMap) MovePlayer(direction string, force int) {
	if direction == "UP" {
		m.Player.Y -= force
	} else if direction == "DOWN" {
		m.Player.Y += force
	} else if direction == "LEFT" {
		m.Player.X -= force
	} else if direction == "RIGHT" {
		m.Player.X += force
	} else {
		log.Print("ERROR: INVALID DIRECTION")
	}
	m.checkPlayerPosition()
}

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

func (m *OddMap) LoadMap(name string) {
	m.CurrentMap = m.Boards[name].Map
	m.currentMapName = name
}

func (m *OddMap) PlainMap() string {
	m.regenerateBoard()
	return formatToBoard(m.CurrentMap)
}

func (m *OddMap) PlayerMap() string {
	a := *m
	//m.regenerateBoard()
	a.CurrentMap[a.Player.Y][a.Player.X] = "X"
	return formatToBoard(m.CurrentMap)
}

func (m *OddMap) checkPlayerPosition() {
	if m.Player.X < 0 {
		m.Player.X = 0
	} else if m.Player.X >= m.width {
		m.Player.X = m.width
	} else if m.Player.Y < 0 {
		m.Player.Y = 0
	} else if m.Player.Y > m.height {
		m.Player.Y = m.height
	}

}

func (m *OddMap) regenerateBoard() {
	m.LoadMap(m.currentMapName)
}
