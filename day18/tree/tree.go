package tree

import (
	"math"
	"strconv"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func (n *Node) leaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) Add(other *Node) *Node {
	newNode := Node{Left: n, Right: other}
	newNode.Left.Parent = &newNode
	newNode.Right.Parent = &newNode
	return &newNode
}

func (n *Node) explode() {
	if (n.Parent.Left) == n {
		// I am the left node
		current := n.Parent.Right
		for current.Left != nil {
			current = current.Left
		}
		current.Value += n.Right.Value

		current = n
		parentNode := n.Parent
		for parentNode != nil && parentNode.Left == current {
			current = parentNode
			parentNode = parentNode.Parent
		}
		if parentNode != nil {
			current = parentNode.Left
			for current.Right != nil {
				current = current.Right
			}
			current.Value += n.Left.Value
		}
	} else {
		// I am the right node
		current := n.Parent.Left
		for current.Right != nil {
			current = current.Right
		}
		current.Value += n.Left.Value

		current = n
		parentNode := n.Parent
		for parentNode != nil && parentNode.Right == current {
			current = parentNode
			parentNode = parentNode.Parent
		}
		if parentNode != nil {
			current = parentNode.Right
			for current.Left != nil {
				current = current.Left
			}
			current.Value += n.Right.Value
		}
	}
}

func (n *Node) ReduceCompletey() *Node {
	changed := true
	current := n
	for changed {
		changed, current = current.reduce(1, true)
		for changed {
			changed, current = current.reduce(1, true)
		}

		changed, current = current.reduce(1, false)
	}

	return current
}

func (n *Node) reduce(depth int, explode bool) (bool, *Node) {
	if explode {
		if !n.leaf() && depth >= 5 {
			n.explode()
			return true, &Node{Value: 0, Parent: n.Parent}
		} else {
			if n.Left != nil {
				leftChanged, newLeft := n.Left.reduce(depth+1, explode)
				if leftChanged {
					n.Left = newLeft
					return true, n
				}
			}

			if n.Right != nil {
				rightChanged, newRight := n.Right.reduce(depth+1, explode)
				if rightChanged {
					n.Right = newRight
					return true, n
				}
			}

			return false, n
		}
	} else {
		if n.leaf() {
			if n.Value >= 10 {
				left := &Node{Value: int(math.Floor(float64(n.Value) / 2.0))}
				right := &Node{Value: int(math.Ceil(float64(n.Value) / 2.0))}
				newNode := Node{Left: left, Right: right, Parent: n.Parent}
				left.Parent = &newNode
				right.Parent = &newNode
				return true, &newNode
			}
			return false, n
		} else {
			if n.Left != nil {
				leftChanged, newLeft := n.Left.reduce(depth+1, explode)
				if leftChanged {
					n.Left = newLeft
					return true, n
				}
			}

			if n.Right != nil {
				rightChanged, newRight := n.Right.reduce(depth+1, explode)
				if rightChanged {
					n.Right = newRight
					return true, n
				}
			}

			return false, n
		}
	}
}

func (n *Node) Magnitude() int {
	if n.leaf() {
		return n.Value
	}
	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func (n *Node) asString() string {
	if n.leaf() {
		return strconv.Itoa(n.Value)
	}
	return "[" + n.Left.asString() + "," + n.Right.asString() + "]"
}
