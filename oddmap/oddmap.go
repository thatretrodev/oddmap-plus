package oddmap

import (
	"bufio"
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

type position struct {
	X int
	Y int
}

type OddMap struct {
	Name     string
	width    int
	height   int
	Board    [][]string
	Player   position
	currentL string
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

func (m *OddMap) LoadMap(path string) {
	m.currentL = path
	file, err := os.Open(path)
	if err != nil {
		//handle error
		return
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	m.Board = make([][]string, 0) //clear loaded board
	for s.Scan() {
		m.Board = append(m.Board, strings.Split(s.Text(), ""))
	}

}

func checkPlayerPosition(m *OddMap) {
	if m.Player.X < 0 {
		m.Player.X = (m.width) + m.Player.X
	} else if m.Player.X >= m.width {
		m.Player.X = (m.Player.X - m.width)
	} else if m.Player.Y < 0 {
		m.Player.Y = (m.height) + m.Player.Y + 1
	} else if m.Player.Y > m.height {
		m.Player.Y = (m.Player.Y - m.height) - 1
	}

}

func (m *OddMap) regenerateBoard() {
	m.LoadMap(m.currentL)
}
