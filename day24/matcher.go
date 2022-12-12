package day24

type matcher func(*value) bool

func constant(n int) matcher {
	return func(val *value) bool {
		return (val.operation == "num" && val.number == n) || (val.min == n && val.max == n)
	}
}

func number(n *int) matcher {
	return func(val *value) bool {
		if val.operation == "num" {
			*n = val.number
			return true
		}
		if val.min == val.max {
			*n = val.min
			return true
		}
		return false
	}
}

func input(n *int) matcher {
	return func(val *value) bool {
		if val.operation == "inp" {
			*n = val.number
			return true
		}
		return false
	}
}

func any(p **value) matcher {
	return func(val *value) bool {
		*p = val
		return true
	}
}

func binaryOperation(leftOperand matcher, operation string, rightOperand matcher) matcher {
	return func(val *value) bool {
		return val.operation == operation && leftOperand(val.left) && rightOperand(val.right)
	}
}
