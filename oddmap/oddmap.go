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
	currentMap     [][]string
	currentMapName string
}

func (m *OddMap) Set(name string, width int, height int) {
	m.Name = name
	m.width = width
	m.height = height
	//code
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
	checkPlayerPosition(m)
}

func (m *OddMap) LoadMaps(path string) {
	//m.currentM = path
	file, err := os.Open(path)
	if err != nil {
		//handle error
		return
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	//m.Board = make([][]string, 0) //clear loaded board
	scannedlines := make([]string, 0)
	for s.Scan() { //scan all lines in files to be loaded to individual maps
		scannedlines = append(scannedlines, s.Text())
	}
	fmt.Println(scannedlines) //works until here
	for x := 0; x < len(scannedlines); x++ {
		y := Map{}
		y.Name = scannedlines[x]
		fmt.Println(x)
		for ; x < x+m.height; x++ {
			y.Map = append(y.Map, strings.Split(scannedlines[x], ""))
		}
		m.Boards[y.Name] = y

	}
	fmt.Println(m.Boards)

}

func (m *OddMap) reloadCurrentMap(name string) {
	m.currentMap = m.Boards[(m.currentMapName)].Map

}

func (m *OddMap) LoadMap(name string) {
	m.reloadCurrentMap(name)
}

func checkPlayerPosition(m *OddMap) {
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
	m.reloadCurrentMap(m.currentMapName)
}
