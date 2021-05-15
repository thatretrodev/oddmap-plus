package oddmap

import (
	"log"

	"github.com/ulule/deepcopier"
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

func (m *OddMap) LoadMap(name string) {
	m.CurrentMap = m.Boards[name].Map
	m.currentMapName = name
}

func (m *OddMap) PlainMap() string {
	a := OddMap{}
	//m.regenerateBoard()
	deepcopier.Copy(m).To(a)
	return formatToBoard(a.CurrentMap)
}

func (m *OddMap) PlayerMap() string {
	a := OddMap{}
	deepcopier.Copy(m).To(a)
	//m.regenerateBoard()
	a.CurrentMap[a.Player.Y][a.Player.X] = "X"
	return formatToBoard(m.CurrentMap)
}

func (m *OddMap) checkPlayerPosition() {
	if m.Player.X < 0 {
		m.Player.X = m.width - 1
	} else if m.Player.X >= m.width {
		m.Player.X = 0
	} else if m.Player.Y < 0 {
		m.Player.Y = m.height - 1
	} else if m.Player.Y > m.height {
		m.Player.Y = 0
	}

}

func (m *OddMap) regenerateBoard() {
	m.LoadMap(m.currentMapName)
}
