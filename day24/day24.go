package day24

import (
	"bufio"
	"os"
	"strings"
)

func readData(filename string) []instruction {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	instructions := make([]instruction, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		left := line[1][0]
		if line[0] == "inp" {
			instructions = append(instructions, inp{varibale: rune(left)})
		} else {
			right := line[2][0]
			switch line[0] {
			case "mul":
				instructions = append(instructions, mul{left: rune(left), right: rune(right)})
			case "add":
				instructions = append(instructions, add{left: rune(left), right: rune(right)})
			case "mod":
				instructions = append(instructions, mod{left: rune(left), right: rune(right)})
			case "div":
				instructions = append(instructions, div{left: rune(left), right: rune(right)})
			case "eql":
				instructions = append(instructions, eql{left: rune(left), right: rune(right)})
			default:
				panic("unknown operator")
			}
		}
	}
	return instructions
}

func ModelNumber(filename string) int {
	instructions := readData(filename)
	context := make(map[rune]int)
	for a := 9; a > 0; a-- {
		for b := 9; b > 0; b-- {
			for c := 9; c > 0; c-- {
				for d := 9; d > 0; d-- {
					for e := 9; e > 0; e-- {
						for f := 9; f > 0; f-- {
							for g := 9; g > 0; g-- {
								for h := 9; h > 0; h-- {
									for i := 9; i > 0; i-- {
										for j := 9; j > 0; j-- {
											for k := 9; k > 0; k-- {
												for l := 9; l > 0; l-- {
													for m := 9; m > 0; m-- {
														for n := 9; n > 0; n-- {
															inpCounter := 0
															for _, instruction := range instructions {
																if _, ok := instruction.(inp); ok {
																	switch inpCounter {
																	case 0:
																		context['i'] = a
																	case 1:
																		context['i'] = b
																	case 2:
																		context['i'] = c
																	case 3:
																		context['i'] = d
																	case 4:
																		context['i'] = e
																	case 5:
																		context['i'] = f
																	case 6:
																		context['i'] = g
																	case 7:
																		context['i'] = h
																	case 8:
																		context['i'] = i
																	case 9:
																		context['i'] = j
																	case 10:
																		context['i'] = k
																	case 11:
																		context['i'] = l
																	case 12:
																		context['i'] = m
																	case 13:
																		context['i'] = n
																	}
																	inpCounter++
																}
																instruction.execute(context)
															}
															if context['z'] == 0 {
																return 1
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return 0
}
