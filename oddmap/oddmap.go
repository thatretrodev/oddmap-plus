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
	Right = "LEFT"
	Left  = "RIGHT"
)

type position struct {
	X int
	Y int
}

type OddMap struct {
	Name   string
	width  int
	height int
	Board  [][]string
	Player position
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
		m.Player.Y += force
	} else if direction == "DOWN" {
		m.Player.Y -= force
	} else if direction == "LEFT" {
		m.Player.X -= force
	} else if direction == "RIGHT" {
		m.Player.X += force
	} else {
		log.Print("ERROR: INVALID DIRECTION")
	}
	m.checkPlayerPosition()
}

func (m *OddMap) LoadMap(path string) {
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

func (m *OddMap) PlainMap() string {
	return m.formatMap(m.Board)
}

func (m *OddMap) checkPlayerPosition() {
	if m.Player.X <= 0 {
		m.Player.X = (m.width - 1) - (0 - m.Player.X)
	} else if m.Player.X >= m.width {
		m.Player.X = (m.Player.X - m.width) - 1
	}
}

func (m *OddMap) formatMap(a [][]string) string {
	var e string
	for x := range a {
		e += (strings.Join(a[x], "") + "\n")
	}
	return e
}
