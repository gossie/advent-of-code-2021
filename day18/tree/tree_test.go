package tree

import "testing"

func TestAdd(t *testing.T) {
	left := Node{Value: 1}
	right := Node{Value: 2}
	n := Node{Left: &left, Right: &right}

	other := Node{Value: 5}

	sum := n.Add(&other)

	if sum.Left.Left.Value != 1 {
		t.Fatalf("left.Left is not %v", 1)
	}
	if sum.Left.Right.Value != 2 {
		t.Fatalf("left.Right is not %v", 2)
	}

	if sum.Right.Value != 5 {
		t.Fatalf("right is not %v", 5)
	}
}

// func TestReduce1(t *testing.T) {
// 	parsed := parseNumber("[[[[[9,8],1],2],3],4]")
// 	reduced := parsed.reduceCompletey()
// 	if reduced.asString() != "[[[[0,9],2],3],4]" {
// 		t.Fatalf("parsed = %v", reduced.asString())
// 	}
// }

// func TestReduce2(t *testing.T) {
// 	parsed := parseNumber("[7,[6,[5,[4,[3,2]]]]]")
// 	reduced := parsed.reduceCompletey()
// 	if reduced.asString() != "[7,[6,[5,[7,0]]]]" {
// 		t.Fatalf("parsed = %v", reduced.asString())
// 	}
// }

// func TestReduce3(t *testing.T) {
// 	parsed := parseNumber("[[6,[5,[4,[3,2]]]],1]")
// 	reduced := parsed.reduceCompletey()
// 	if reduced.asString() != "[[6,[5,[7,0]]],3]" {
// 		t.Fatalf("parsed = %v", reduced.asString())
// 	}
// }

// func TestReduce4(t *testing.T) {
// 	parsed := parseNumber("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
// 	reduced := parsed.reduceCompletey()
// 	if reduced.asString() != "[[3,[2,[8,0]]],[9,[5,[7,0]]]]" {
// 		t.Fatalf("parsed = %v", reduced.asString())
// 	}
// }

// func TestReduce5(t *testing.T) {
// 	parsed := parseNumber("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
// 	reduced := parsed.reduceCompletey()
// 	if reduced.asString() != "[[3,[2,[8,0]]],[9,[5,[7,0]]]]" {
// 		t.Fatalf("parsed = %v", reduced.asString())
// 	}
// }

// func TestMagnitude1(t *testing.T) {
// 	parsed := parseNumber("[[1,2],[[3,4],5]]")
// 	magnitude := parsed.magnitude()
// 	if magnitude != 143 {
// 		t.Fatalf("magnitude = %v", magnitude)
// 	}
// }

// func TestMagnitude2(t *testing.T) {
// 	parsed := parseNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
// 	magnitude := parsed.magnitude()
// 	if magnitude != 1384 {
// 		t.Fatalf("magnitude = %v", magnitude)
// 	}
// }

// func TestMagnitude3(t *testing.T) {
// 	parsed := parseNumber("[[[[1,1],[2,2]],[3,3]],[4,4]]")
// 	magnitude := parsed.magnitude()
// 	if magnitude != 445 {
// 		t.Fatalf("magnitude = %v", magnitude)
// 	}
// }

// func TestMagnitude4(t *testing.T) {
// 	parsed := parseNumber("[[[[3,0],[5,3]],[4,4]],[5,5]]")
// 	magnitude := parsed.magnitude()
// 	if magnitude != 791 {
// 		t.Fatalf("magnitude = %v", magnitude)
// 	}
// }

// func TestMagnitude5(t *testing.T) {
// 	parsed := parseNumber("[[[[5,0],[7,4]],[5,5]],[6,6]]")
// 	magnitude := parsed.magnitude()
// 	if magnitude != 1137 {
// 		t.Fatalf("magnitude = %v", magnitude)
// 	}
// }

// func TestMagnitude6(t *testing.T) {
// 	parsed := parseNumber("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
// 	magnitude := parsed.magnitude()
// 	if magnitude != 3488 {
// 		t.Fatalf("magnitude = %v", magnitude)
// 	}
// }

// func TestAddAndReduce(t *testing.T) {
// 	parsed1 := parseNumber("[[[[4,3],4],4],[7,[[8,4],9]]]")
// 	parsed2 := parseNumber("[1,1]")
// 	sum := parsed1.add(parsed2)
// 	if sum.asString() != "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]" {
// 		t.Fatalf("sum = %v", sum.asString())
// 	}

// 	reduced := sum.reduceCompletey()
// 	if reduced.asString() != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
// 		t.Fatalf("reduced = %v", reduced.asString())
// 	}
// }
