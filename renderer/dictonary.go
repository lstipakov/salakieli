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
	dict['A'] = edge1 | edge2 | edge5 | edge6 | edge7 | edge8 | edge9 | edge10 | edge13 | edge14 | edge16
	dict['B'] = edge2 | edge3 | edge9 | edge11 | edge13 | edge14 | edge15
	dict['C'] = edge3 | edge4 | edge9 | edge10 | edge11 | edge12 | edge13 | edge14 | edge15 | edge16
	dict['D'] = edge1 | edge2 | edge3 | edge4 | edge5 | edge6 | edge10 | edge11 | edge12 | edge13 | edge14 | edge16
	dict['E'] = edge3 | edge4 | edge9 | edge10 | edge11 | edge12 | edge13 | edge14 | edge15
	dict['F'] = edge3 | edge4 | edge5 | edge6 | edge9 | edge10 | edge11 | edge12 | edge13 | edge14 | edge15
	dict['G'] = edge3 | edge9 | edge10 | edge11 | edge13 | edge14 | edge15 | edge16
	dict['H'] = edge1 | edge2 | edge5 | edge6 | edge9 | edge10 | edge11 | edge13 | edge14 | edge15
	dict['I'] = edge3 | edge4 | edge7 | edge8 | edge9 | edge11 | edge12 | edge13 | edge15 | edge16
	dict['J'] = edge3 | edge4 | edge5 | edge7 | edge8 | edge9 | edge11 | edge12 | edge13 | edge15 | edge16
	dict['K'] = edge1 | edge2 | edge3 | edge4 | edge5 | edge6 | edge9 | edge10 | edge12 | edge14 | edge15
	dict['L'] = edge1 | edge2 | edge3 | edge4 | edge9 | edge10 | edge11 | edge12 | edge13 | edge14 | edge15 | edge16
	dict['M'] = edge1 | edge2 | edge5 | edge6 | edge10 | edge12 | edge13 | edge14 | edge15 | edge16
	dict['N'] = edge1 | edge2 | edge5 | edge6 | edge10 | edge11 | edge12 | edge14 | edge15 | edge16
	dict['O'] = edge9 | edge10 | edge11 | edge12 | edge13 | edge14 | edge15 | edge16
	dict['P'] = edge2 | edge3 | edge4 | edge5 | edge6 | edge9 | edge11 | edge12 | edge13 | edge14 | edge15
	dict['Q'] = edge9 | edge10 | edge11 | edge12 | edge14 | edge15 | edge16
	dict['R'] = edge2 | edge3 | edge4 | edge5 | edge6 | edge9 | edge11 | edge12 | edge14 | edge15
	dict['S'] = edge3 | edge7 | edge9 | edge10 | edge11 | edge13 | edge14 | edge15
	dict['T'] = edge3 | edge4 | edge5 | edge6 | edge7 | edge8 | edge9 | edge11 | edge12 | edge13 | edge15 | edge16
	dict['U'] = edge1 | edge2 | edge9 | edge10 | edge11 | edge12 | edge13 | edge14 | edge15 | edge16
	dict['V'] = edge1 | edge2 | edge5 | edge6 | edge7 | edge8 | edge10 | edge11 | edge12 | edge14 | edge15 | edge16
	dict['W'] = edge1 | edge2 | edge5 | edge6 | edge9 | edge10 | edge11 | edge12 | edge14 | edge16
	dict['X'] = edge1 | edge2 | edge3 | edge4 | edge5 | edge6 | edge7 | edge8 | edge10 | edge12 | edge14 | edge16
	dict['Y'] = edge1 | edge2 | edge3 | edge4 | edge5 | edge6 | edge7 | edge8 | edge10 | edge12 | edge13 | edge15 | edge16
	dict['Z'] = edge3 | edge4 | edge7 | edge8 | edge9 | edge10 | edge12 | edge13 | edge14 | edge16
}
