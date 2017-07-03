package renderer

const (
	edge1  = 1 << 0
	edge2  = 1 << 1
	edge3  = 1 << 2
	edge4  = 1 << 3
	edge5  = 1 << 4
	edge6  = 1 << 5
	edge7  = 1 << 6
	edge8  = 1 << 7
	edge9  = 1 << 8
	edge10 = 1 << 9
	edge11 = 1 << 10
	edge12 = 1 << 11
	edge13 = 1 << 12
	edge14 = 1 << 13
	edge15 = 1 << 14
	edge16 = 1 << 15
)

// DICT is a map between byte and edges
var dict = make(map[byte]int)

func init() {
	dict['A'] = edge1 | edge2 | edge5 | edge6 | edge7 | edge8 | edge9 | edge13 | edge16
	dict['B'] = edge2 | edge3 | edge9 | edge11 | edge13 | edge14 | edge15
}
