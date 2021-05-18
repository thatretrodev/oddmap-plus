package oddmap

import "log"

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
	//m.CurrentMap = m.Boards[name].Map
	m.Map = m.Boards[name]
	m.currentMapName = name
}

func (m OddMap) PlainMap() string {
	return formatToBoard(m.Map.Map)
}

func (m *OddMap) checkPlayerPosition() {
	if m.Player.X < 0 {
		m.Player.X = m.width
	} else if m.Player.X >= m.width {
		m.Player.X = 0
	} else if m.Player.Y < 0 {
		m.Player.Y = m.height
	} else if m.Player.Y > m.height {
		m.Player.Y = 0
	}

}

func (m *OddMap) SetPlayerPosition(x int, y int) {
	m.Player.X = x
	m.Player.Y = y
}
