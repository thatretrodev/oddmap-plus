package oddmap

type oMap struct {
	Name   string
	width  int
	height int
	board  [][]rune
}

func (m oMap) set(name string, width int, height int) {
	m.Name = name
	m.width = width
	m.height = height
	//code
}
