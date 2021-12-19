package parser

import "testing"

func TestParsePairOfTwoRegularNumbers(t *testing.T) {
	parsed := parseNumber("[1,2]")
	if parsed.Left.Value != 1 {
		t.Fatalf("left = %v", parsed.Left.Value)
	}
	if parsed.Right.Value != 2 {
		t.Fatalf("right = %v", parsed.Right.Value)
	}
}

func TestParsePairOfPairAndRegularNumber(t *testing.T) {
	parsed := parseNumber("[[1,2],3]")
	if parsed.Left.Left.Value != 1 {
		t.Fatalf("left.Left = %v", parsed.Left.Left.Value)
	}
	if parsed.Left.Right.Value != 2 {
		t.Fatalf("left.Right = %v", parsed.Left.Right.Value)
	}
	if parsed.Right.Value != 3 {
		t.Fatalf("right = %v", parsed.Right.Value)
	}
}

func TestParsePairOfRegularNumberAndPair(t *testing.T) {
	parsed := parseNumber("[1,[2,3]]")
	if parsed.Left.Value != 1 {
		t.Fatalf("left = %v", parsed.Left.Value)
	}
	if parsed.Right.Left.Value != 2 {
		t.Fatalf("right.Left = %v", parsed.Right.Left.Value)
	}
	if parsed.Right.Right.Value != 3 {
		t.Fatalf("right.Right = %v", parsed.Right.Right.Value)
	}
}

func TestParsePairOfTwoPairs(t *testing.T) {
	parsed := parseNumber("[[1,2],[3,4]]")
	if parsed.Left.Left.Value != 1 {
		t.Fatalf("left.Left = %v", parsed.Left.Left.Value)
	}
	if parsed.Left.Right.Value != 2 {
		t.Fatalf("left.Right = %v", parsed.Left.Right.Value)
	}
	if parsed.Right.Left.Value != 3 {
		t.Fatalf("right.Left = %v", parsed.Right.Left.Value)
	}
	if parsed.Right.Right.Value != 4 {
		t.Fatalf("right.Right = %v", parsed.Right.Right.Value)
	}
}
