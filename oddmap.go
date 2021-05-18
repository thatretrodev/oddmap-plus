package oddmap

import (
	"log"
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
	Map            Map
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
