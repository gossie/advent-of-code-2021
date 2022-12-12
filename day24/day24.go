package day24

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var operationMapping = map[string]string{
	"mul": "*",
	"add": "+",
	"div": "/",
	"mod": "%",
	"eql": "==",
}

func readData(filename string) []*value {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	program := make([]*value, 0)
	values := map[string]*value{
		"w": {id: 1, operation: "num", number: 0},
		"x": {id: 2, operation: "num", number: 0},
		"y": {id: 3, operation: "num", number: 0},
		"z": {id: 4, operation: "num", number: 0},
	}

	program = append(program, values["w"])
	program = append(program, values["x"])
	program = append(program, values["y"])
	program = append(program, values["z"])
	id := 4
	inputCounter := -1

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		left := line[1]
		if line[0] == "inp" {
			id++
			inputCounter++
			inp := &value{id: id, operation: "inp", number: inputCounter}
			values[left] = inp
			program = append(program, inp)
		} else {
			right := line[2]

			leftValue := values[left]
			rightValue, ok := values[right]
			if !ok {
				number, _ := strconv.Atoi(right)
				id++
				rightValue = &value{id: id, operation: "num", number: number}
				program = append(program, rightValue)
			}

			id++
			val := &value{id: id, operation: operationMapping[line[0]], left: leftValue, right: rightValue}
			values[left] = val
			program = append(program, val)
		}
	}
	return program
}

func MaxModelNumber(filename string) string {
	program := readData(filename)

	optimize(program)
	force(program)
	optimize(program)
	printProgram(program)
	return maxModelNumber(program)
}

func MinModelNumber(filename string) string {
	program := readData(filename)

	optimize(program)
	force(program)
	optimize(program)
	printProgram(program)
	return minModelNumber(program)
}

func maxModelNumber(program []*value) string {
	modelNumber := make([]int, 14)
	for _, val := range program {
		var i, j, a, b, c int
		switch {
		case binaryOperation(binaryOperation(binaryOperation(input(&i), "+", number(&a)), "+", number(&b)), "force", input(&j))(val),
			binaryOperation(binaryOperation(input(&i), "+", number(&a)), "force", input(&j))(val):
			a += b
			if a > 0 {
				modelNumber[i] = 9 - a
				modelNumber[j] = 9
			} else {
				modelNumber[i] = 9
				modelNumber[j] = 9 + a
			}
		case binaryOperation(binaryOperation(binaryOperation(binaryOperation(input(&i), "+", number(&a)), "%", number(&b)), "+", number(&c)), "force", input(&j))(val):
			modelNumber[i] = 9
			modelNumber[j] = ((9 + a) % b) + c
		}
	}
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(modelNumber), " "), ""), "[]")
}

func minModelNumber(program []*value) string {
	modelNumber := make([]int, 14)
	for _, val := range program {
		var i, j, a, b, c int
		switch {
		case binaryOperation(binaryOperation(binaryOperation(input(&i), "+", number(&a)), "+", number(&b)), "force", input(&j))(val),
			binaryOperation(binaryOperation(input(&i), "+", number(&a)), "force", input(&j))(val):
			a += b
			if a > 0 {
				modelNumber[i] = 1
				modelNumber[j] = 1 + a
			} else {
				modelNumber[i] = 1 - a
				modelNumber[j] = 1
			}
		case binaryOperation(binaryOperation(binaryOperation(binaryOperation(input(&i), "+", number(&a)), "%", number(&b)), "+", number(&c)), "force", input(&j))(val):
			modelNumber[i] = int(math.Abs(float64(a+c))) + 1
			modelNumber[j] = ((modelNumber[i] + a) % b) + c
		}
	}
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(modelNumber), " "), ""), "[]")
}

func optimize(program []*value) {
	// id := 1000
	for _, val := range program {
		if val.left != nil && val.left.fwd != nil {
			val.left = val.left.fwd
		}
		if val.right != nil && val.right.fwd != nil {
			val.right = val.right.fwd
		}
		var a, b int
		var x, y *value
		switch {
		case binaryOperation(number(&a), "+", number(&b))(val):
			setNumber(val, a+b)
		case binaryOperation(number(&a), "*", number(&b))(val):
			setNumber(val, a*b)
		case binaryOperation(number(&a), "/", number(&b))(val):
			setNumber(val, a/b)
		case binaryOperation(number(&a), "%", number(&b))(val):
			setNumber(val, a%b)
		case binaryOperation(constant(0), "*", any(&x))(val),
			binaryOperation(any(&x), "*", constant(0))(val),
			binaryOperation(constant(0), "/", any(&x))(val):
			setNumber(val, 0)
		case binaryOperation(any(&x), "+", constant(0))(val),
			binaryOperation(constant(0), "+", any(&x))(val),
			binaryOperation(any(&x), "/", constant(1))(val),
			binaryOperation(any(&x), "*", constant(1))(val), binaryOperation(constant(1), "*", any(&x))(val):
			val.fwd = x
		case binaryOperation(any(&x), "==", any(&y))(val) && (y.min > x.max || y.max < x.min):
			setNumber(val, 0)
		case binaryOperation(number(&a), "==", number(&b))(val) && (a == b):
			setNumber(val, 1)
		case binaryOperation(binaryOperation(binaryOperation(any(&y), "*", number(&b)), "+", any(&x)), "%", number(&a))(val) && a == b && x.max < a:
			val.fwd = x
		case binaryOperation(binaryOperation(binaryOperation(any(&y), "*", number(&b)), "+", any(&x)), "/", number(&a))(val) && x.max < a:
			val.fwd = y
		}

		switch val.operation {
		default:
			panic("min/max " + val.operation)
		case "num":
			val.min = val.number
			val.max = val.number
		case "inp":
			val.min = 1
			val.max = 9
		case "==":
			val.min = 0
			val.max = 1
		case "+":
			val.min = val.left.min + val.right.min
			val.max = val.left.max + val.right.max
		case "*":
			if val.left.min < 0 || val.right.min < 0 {
				panic("min/max neg *")
			}
			val.min = val.left.min * val.right.min
			val.max = val.left.max * val.right.max
		case "%":
			if val.right.operation != "num" {
				panic("division by non const")
			}
			val.min = 0
			val.max = val.right.number - 1
		case "/":
			if val.right.operation != "num" {
				panic("division by non const")
			}
			val.min = val.left.min / val.right.number
			val.max = val.left.max / val.right.number
		case "force":
			val.min = 0
			val.max = 0
		}
	}
}

func setNumber(val *value, n int) {
	*val = value{operation: "num", number: n}
}

func force(program []*value) {
	maxes := make(map[*value]int)
	maxes[program[len(program)-1]] = 0

	updateMax := func(val *value, max int) {
		if old, ok := maxes[val]; ok && old < max {
			return
		}
		maxes[val] = max
	}

	for i := len(program) - 1; i >= 0; i-- {
		val := program[i]
		max, ok := maxes[val]
		if !ok {
			continue
		}
		if val.max <= max {
			continue
		}
		var a int
		var x, y *value
		switch {
		case number(&a)(val):
			if a > max {
				panic("force impossible")
			}
		case binaryOperation(any(&x), "+", any(&y))(val):
			updateMax(x, max-y.min)
			updateMax(y, max-x.min)
		case binaryOperation(any(&x), "*", any(&y))(val):
			if y.min > 0 {
				updateMax(x, max/y.min)
			}
			if x.min > 0 {
				updateMax(y, max/x.min)
			}
		case binaryOperation(any(&x), "/", number(&a))(val):
			updateMax(x, max*a+a-1)
		case binaryOperation(binaryOperation(any(&x), "==", any(&y)), "==", constant(0))(val):
			val.operation = "force"
			val.left = x
			val.right = y
			val.min = 0
			val.max = 0
			updateMax(x, y.max)
			updateMax(y, x.max)
		default:
			//panic("forcing unknown operator " + val.operation)
		}

	}
}

func printProgram(program []*value) {
	count := make(map[*value]int)
	for i := len(program) - 1; i >= 0; i-- {
		val := program[i]
		if count[val] == 0 && i != len(program)-1 {
			continue
		}
		count[val.left]++
		count[val.right]++
	}

	toRender := make(map[*value]string)
	for _, val := range program {
		var x string
		switch val.operation {
		case "inp", "num":
			x = val.Init()
		default:
			x = fmt.Sprintf("(%v %v %v)", toRender[val.left], val.operation, toRender[val.right])
			if count[val] > 1 || val.operation == "force" {
				fmt.Printf("%s = %v // [%d, %d]\n", val.Name(), x, val.min, val.max)
				x = val.Name()
			}
		}
		toRender[val] = x
	}
	fmt.Println(toRender[program[len(program)-1]])
}

//(((((((u * 26) + v) + w) + x) + y) + z) % 26)
