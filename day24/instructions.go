package day24

type instruction interface {
	execute(context map[rune]int)
}

type inp struct {
	varibale rune
}

func (m inp) execute(context map[rune]int) {
	context[m.varibale] = context['i']
}

type mul struct {
	left  rune
	right rune
}

func (m mul) execute(context map[rune]int) {
	factor1 := context[m.left]
	factor2 := int(m.right - '0')
	if factor2 < 0 || factor2 > 9 {
		factor2 = context[m.right]
	}
	context[m.left] = factor1 * factor2
}

type add struct {
	left  rune
	right rune
}

func (a add) execute(context map[rune]int) {
	summand1 := context[a.left]
	summand2 := int(a.right - '0')
	if summand2 < 0 || summand2 > 9 {
		summand2 = context[a.right]
	}
	context[a.left] = summand1 + summand2
}

type mod struct {
	left  rune
	right rune
}

func (m mod) execute(context map[rune]int) {
	value1 := context[m.left]
	value2 := int(m.right - '0')
	if value2 < 0 || value2 > 9 {
		value2 = context[m.right]
	}
	context[m.left] = value1 % value2
}

type div struct {
	left  rune
	right rune
}

func (d div) execute(context map[rune]int) {
	value1 := context[d.left]
	value2 := int(d.right - '0')
	if value2 < 0 || value2 > 9 {
		value2 = context[d.right]
	}
	context[d.left] = value1 / value2
}

type eql struct {
	left  rune
	right rune
}

func (d eql) execute(context map[rune]int) {
	value1 := context[d.left]
	value2 := int(d.right - '0')
	if value2 < 0 || value2 > 9 {
		value2 = context[d.right]
	}
	if value1 == value2 {
		context[d.left] = 1
	} else {
		context[d.left] = 0
	}
}
